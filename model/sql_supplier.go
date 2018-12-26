package model

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/KenmyZhang/smart-edu-server/common/config"
	"github.com/KenmyZhang/smart-edu-server/common/util"
	"github.com/KenmyZhang/smart-edu-server/log"
)

const (
	EXIT_DB_OPEN = 101
	EXIT_PING    = 102
)

const (
	MAX_DB_CONN_LIFETIME = 60
	DB_PING_ATTEMPTS     = 18
	DB_PING_TIMEOUT_SECS = 10
	DB_OPEN_ATTEMPTS     = 18
	DB_OPEN_TIMEOUT_SECS = 10
)

type SqlSupplier struct {
	master         *gorm.DB
	replicas       []*gorm.DB
	searchReplicas []*gorm.DB
	cfg            *config.SqlSettings
	rrCounter      int64
	srCounter      int64
}

var sqlSupplier *SqlSupplier

func NewSqlSupplier(cfg *config.SqlSettings) *SqlSupplier {
	sqlSupplier = &SqlSupplier{
		rrCounter: 0,
		srCounter: 0,
	}

	sqlSupplier.cfg = cfg
	sqlSupplier.InitConnection()
	InitCollectionTable(sqlSupplier)
	return sqlSupplier
}

func InitCollectionTable(ss *SqlSupplier) {
	if !ss.GetMaster().HasTable(&User{}) {
		if err := ss.GetMaster().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	if !ss.GetMaster().HasTable(&Follower{}) {
		if err := ss.GetMaster().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Follower{}).Error; err != nil {
			panic(err)
		}
	}
}

func (ss *SqlSupplier) InitConnection() {
	ss.master = ss.setupConnection("master", *ss.cfg.DriverName, *ss.cfg.DataSource,
		*ss.cfg.MaxIdleConns, *ss.cfg.MaxOpenConns, *ss.cfg.Trace)
	if len(ss.cfg.DataSourceReplicas) == 0 {
		ss.replicas = make([]*gorm.DB, 1)
		ss.replicas[0] = ss.master
	} else {
		ss.replicas = make([]*gorm.DB, len(ss.cfg.DataSourceReplicas))
		for i, replica := range ss.cfg.DataSourceReplicas {
			ss.replicas[i] = ss.setupConnection(fmt.Sprintf("replica-%v", i),
				*ss.cfg.DriverName, replica, *ss.cfg.MaxIdleConns, *ss.cfg.MaxOpenConns, *ss.cfg.Trace)
		}
	}

	if len(ss.cfg.DataSourceSearchReplicas) == 0 {
		ss.searchReplicas = ss.replicas
	} else {
		ss.searchReplicas = make([]*gorm.DB, len(ss.cfg.DataSourceSearchReplicas))
		for i, replica := range ss.cfg.DataSourceSearchReplicas {
			ss.searchReplicas[i] = ss.setupConnection(fmt.Sprintf("search-replica-%v", i),
				*ss.cfg.DriverName, replica, *ss.cfg.MaxIdleConns, *ss.cfg.MaxOpenConns, *ss.cfg.Trace)
		}
	}
	return
}

func (ss *SqlSupplier) setupConnection(con_type string, driver string, dataSource string, maxIdle int, maxOpen int, trace bool) *gorm.DB {
	log.Info("driver:" + con_type)
	log.Info("open database:" + dataSource)
	var db *gorm.DB
	var err error
	for i := 0; i < DB_OPEN_ATTEMPTS; i++ {
		db, err = gorm.Open(driver, dataSource)
		if err == nil {
			log.Info("open database successfully")
			break
		} else {
			if i == DB_OPEN_ATTEMPTS-1 {
				log.Critical(err.Error())
				time.Sleep(time.Second)
				os.Exit(EXIT_DB_OPEN)
				return nil
			} else {
				log.Error(fmt.Sprintf("Failed to Open DB retrying  in %v seconds err = %v", DB_OPEN_TIMEOUT_SECS, err))
				time.Sleep(DB_OPEN_TIMEOUT_SECS * time.Second)
			}
		}
	}

	for i := 0; i < DB_PING_ATTEMPTS; i++ {
		log.Info("Pinging SQL " + con_type + " database")
		if err := db.DB().Ping(); err == nil {
			break
		} else {
			if i == DB_PING_ATTEMPTS-1 {
				log.Critical("Failed to ping DB, server will exit err =" + err.Error())
				time.Sleep(time.Second)
				os.Exit(EXIT_PING)
				return nil
			} else {
				log.Error(fmt.Sprintf("Failed to ping DB retrying  in %v seconds err = %v", DB_PING_TIMEOUT_SECS, err))
				time.Sleep(DB_PING_TIMEOUT_SECS * time.Second)
			}
		}
	}
	//设置数据库的空闲连接和最大打开连接
	db.DB().SetMaxIdleConns(*ss.cfg.MaxIdleConns)
	db.DB().SetMaxOpenConns(*ss.cfg.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(MAX_DB_CONN_LIFETIME) * time.Minute)
	db.LogMode(*ss.cfg.Trace)
	return db
}
func (ss *SqlSupplier) GetMaster() *gorm.DB {
	return ss.master
}

func (ss *SqlSupplier) GetSearchReplica() *gorm.DB {
	rrNum := atomic.AddInt64(&ss.srCounter, 1) % int64(len(ss.searchReplicas))
	return ss.searchReplicas[rrNum]
}

func (ss *SqlSupplier) GetReplica() *gorm.DB {
	rrNum := atomic.AddInt64(&ss.rrCounter, 1) % int64(len(ss.replicas))
	return ss.replicas[rrNum]
}

func (ss *SqlSupplier) TotalMasterDbConnections() int {
	return ss.GetMaster().DB().Stats().OpenConnections
}

func (ss *SqlSupplier) TotalReadDbConnections() int {

	if len(ss.cfg.DataSourceReplicas) == 0 {
		return 0
	}

	count := 0
	for _, db := range ss.replicas {
		count = count + db.DB().Stats().OpenConnections
	}

	return count
}

func (ss *SqlSupplier) TotalSearchDbConnections() int {
	if len(ss.cfg.DataSourceSearchReplicas) == 0 {
		return 0
	}

	count := 0
	for _, db := range ss.searchReplicas {
		count = count + db.DB().Stats().OpenConnections
	}

	return count
}

func (ss *SqlSupplier) Close() *util.Err {
	log.Info("Closing SqlStore")
	if err := ss.master.DB().Close(); err != nil {
		return util.NewInternalServerError("SqlSupplier.Close", err.Error())
	}
	for _, replica := range ss.replicas {
		if err := replica.DB().Close(); err != nil {
			return util.NewInternalServerError("SqlSupplier.replicas.Close", err.Error())
		}
	}
	return nil
}
