package app

import (
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/KenmyZhang/smart-edu-server/common/config"
	"github.com/KenmyZhang/smart-edu-server/log"
	"github.com/KenmyZhang/smart-edu-server/model"
	"github.com/KenmyZhang/smart-edu-server/route"
	"net/http"
	"os"
	"time"
	"gopkg.in/mgo.v2"
)

type GracefulServer struct {
	Server      *http.Server
	Router      *gin.Engine
	SqlSupplier *model.SqlSupplier
	MgoSupplier *mgo.Session
	Log         *log.Logger
	Cfg         *config.Config
}

var Srv *GracefulServer

func NewServer() *GracefulServer {
	log.Info("new server")
	Srv = &GracefulServer{Cfg: config.Cfg}
	Srv.Log = log.NewLogger(Srv.Cfg.LoggerConfigFromLoggerConfig())
	Srv.Router = route.NewRoute()
	Srv.SqlSupplier = model.NewSqlSupplier(Srv.Cfg.SqlSettings)
	Srv.MgoSupplier = model.NewMgoSupplier(Srv.Cfg.MgoEndpoint)
	// 将golang中默认的 logger重定向到这个指定的server logger
	log.RedirectStdLog(Srv.Log)
	// 使用server logger 作为全局的logger
	log.InitGlobalLogger(Srv.Log)
	return Srv
}

func (gracefulServer *GracefulServer) Start() {
	log.Info("server start")
	gracefulServer.Server = &http.Server{
		Addr:         gracefulServer.Cfg.ServiceSettings.ListenAddress,
		Handler:      gracefulServer.Router,
		ReadTimeout:  time.Duration(gracefulServer.Cfg.ServiceSettings.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(gracefulServer.Cfg.ServiceSettings.WriteTimeout) * time.Second,
	}

	log.Info("Start Listening and serving HTTP on " + gracefulServer.Cfg.ServiceSettings.ListenAddress)
	if err := gracehttp.Serve(gracefulServer.Server); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	return
}

func (gracefulServer *GracefulServer) ShutDown() {
	if gracefulServer.Server != nil {
		if err := gracefulServer.Server.Close(); err != nil {
			log.Warn(err.Error())
			os.Exit(1)
		}
	}
	if gracefulServer.SqlSupplier != nil {
		if err := gracefulServer.SqlSupplier.Close(); err != nil {
			log.Warn(err.Error())
			os.Exit(1)
		}
	}
	log.Info("shutdown server successfully")
}
