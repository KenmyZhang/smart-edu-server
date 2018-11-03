package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"smart-edu-server/log"
	"strings"
)

const (
	LOG_FILENAME = "smart-edu-server.log"
)

var (
	commonBaseSearchPaths = []string{
		".",
		"..",
		"../..",
		"../../..",
	}
)

type Config struct {
	FilePath             string
	ReleaseMode          bool
	ServiceSettings      ServiceSettings
	LogSettings          LogSettings
	SentrySettings       SentrySettings
	SqlSettings          *SqlSettings
	RedisSettings        RedisSettings
	WexinSettings        WexinSettings
	OfflineSpaceSettings OfflineSpaceSettings
}

type ServiceSettings struct {
	ListenAddress string
	ReadTimeout   int
	WriteTimeout  int
}

type SqlSettings struct {
	DriverName               *string
	DataSource               *string
	DataSourceReplicas       []string
	DataSourceSearchReplicas []string
	MaxIdleConns             *int
	MaxOpenConns             *int
	Trace                    *bool
	QueryTimeout             *int
}

type OfflineSpaceSettings struct {
	Url string
}

type RedisSettings struct {
	Addr         *string
	Password     *string
	DialTimeout  *int64
	ReadTimeout  *int64
	WriteTimeout *int64
	PoolSize     *int
	PoolTimeOut  *int64
	DB           *int
}

type LogSettings struct {
	EnableConsole bool
	ConsoleLevel  string
	ConsoleJson   *bool
	EnableFile    bool
	FileLevel     string
	FileJson      *bool
	FileLocation  string
	MaxSize       *int
	MaxAge        *int
	MaxBackups    *int
	LocalTime     *bool
	Compress      *bool
}

type WexinSettings struct {
	AppId     string
	AppSecret string
}

type SentrySettings struct {
	SentryDsn string
	SentryLog string
}

const (
	SERVICE_SETTINGS_DEFAULT_LISTEN_AND_ADDRESS = ":8065"
	DATABASE_DRIVER_MYSQL                       = "mysql"
	SQL_SETTINGS_DEFAULT_DATA_SOURCE            = "mmuser:mostest@tcp(dockerhost:3306)/matterst_test?charset=utf8mb4,utf8&readTimeout=30s&writeTimeout=30s"
)

var Cfg = &Config{}

var configPath = flag.String("config", "config.json", "path to the config file")

func init() {
	flag.Parse()
	LoadConfig(*configPath)
}

func (o *Config) SetDefaults() {
	o.LogSettings.SetDefaults()
	o.ServiceSettings.SetDefaults()
	o.SqlSettings.SetDefaults()
	o.RedisSettings.SetDefaults()
}

func (s *SqlSettings) SetDefaults() {
	if s.DriverName == nil {
		s.DriverName = NewString(DATABASE_DRIVER_MYSQL)
	}

	if s.DataSource == nil {
		s.DataSource = NewString(SQL_SETTINGS_DEFAULT_DATA_SOURCE)
	}

	if s.MaxIdleConns == nil {
		s.MaxIdleConns = NewInt(20)
	}

	if s.MaxOpenConns == nil {
		s.MaxOpenConns = NewInt(300)
	}

	if s.QueryTimeout == nil {
		s.QueryTimeout = NewInt(30)
	}
}

func (s *RedisSettings) SetDefaults() {
	if s.Addr == nil {
		s.Addr = NewString("127.0.0.1:6379")
	}

	if s.Password == nil {
		s.Password = NewString("")
	}

	if s.DialTimeout == nil {
		s.DialTimeout = NewInt64(5)
	}

	if s.ReadTimeout == nil {
		s.ReadTimeout = NewInt64(3)
	}

	if s.WriteTimeout == nil {
		s.WriteTimeout = NewInt64(3)
	}

	if s.PoolSize == nil {
		s.PoolSize = NewInt(30)
	}

	if s.PoolTimeOut == nil {
		s.PoolTimeOut = NewInt64(3)
	}
}

func (s *ServiceSettings) SetDefaults() {
	if s.ListenAddress == "" {
		s.ListenAddress = SERVICE_SETTINGS_DEFAULT_LISTEN_AND_ADDRESS
	}
}

func (s *LogSettings) SetDefaults() {
	if s.ConsoleJson == nil {
		s.ConsoleJson = NewBool(true)
	}

	if s.FileJson == nil {
		s.FileJson = NewBool(true)
	}

	if s.MaxSize == nil {
		s.MaxSize = NewInt(100)
	}

	if s.MaxAge == nil {
		s.MaxAge = NewInt(365)
	}

	if s.MaxBackups == nil {
		s.MaxBackups = NewInt(30)
	}

	if s.LocalTime == nil {
		s.LocalTime = NewBool(true)
	}

	if s.Compress == nil {
		s.Compress = NewBool(true)
	}
}

func NewBool(b bool) *bool       { return &b }
func NewInt(n int) *int          { return &n }
func NewInt64(n int64) *int64    { return &n }
func NewString(s string) *string { return &s }

func (o *Config) LoggerConfigFromLoggerConfig() *log.LoggerConfiguration {
	return &log.LoggerConfiguration{
		EnableConsole: o.LogSettings.EnableConsole,
		ConsoleJson:   *o.LogSettings.ConsoleJson,
		ConsoleLevel:  strings.ToLower(o.LogSettings.ConsoleLevel),
		EnableFile:    o.LogSettings.EnableFile,
		FileJson:      *o.LogSettings.FileJson,
		FileLevel:     strings.ToLower(o.LogSettings.FileLevel),
		FileLocation:  GetLogFileLocation(o.LogSettings.FileLocation),
		MaxSize:       *o.LogSettings.MaxSize,
		MaxAge:        *o.LogSettings.MaxAge,
		MaxBackups:    *o.LogSettings.MaxBackups,
		LocalTime:     *o.LogSettings.LocalTime,
		Compress:      *o.LogSettings.Compress,
	}
}

func GetConfig() *Config {
	return Cfg
}

func LoadConfig(filename string) *Config {
	if filename == "" {
		filename = "config.json"
	}

	filename = FindConfigFile(filename)
	log.Info("config file:" + filename)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(data, Cfg)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	Cfg.SetDefaults()
	j, _ := json.MarshalIndent(Cfg, "", "\t")
	fmt.Println(string(j))
	return Cfg
}

func FindConfigFile(fileName string) (path string) {
	found := FindFile(filepath.Join("conf", fileName))
	if found == "" {
		if found = FindFile(filepath.Join("config", fileName)); found == "" {
			found = FindPath(fileName, []string{"."}, nil)
		}
	}

	return found
}

func FindPath(path string, baseSearchPaths []string, filter func(os.FileInfo) bool) string {
	//判斷是否是絕對路徑
	if filepath.IsAbs(path) {
		if _, err := os.Stat(path); err == nil {
			return path
		}

		return ""
	}

	searchPaths := []string{}
	for _, baseSearchPath := range baseSearchPaths {
		searchPaths = append(searchPaths, baseSearchPath)
	}

	var binaryDir string
	//返回启动当前进程的可执行文件的路径名称。
	if exe, err := os.Executable(); err == nil {
		if exe, err = filepath.EvalSymlinks(exe); err == nil {
			if exe, err = filepath.Abs(exe); err == nil {
				binaryDir = filepath.Dir(exe)
			}
		}
	}
	if binaryDir != "" {
		for _, baseSearchPath := range baseSearchPaths {
			searchPaths = append(
				searchPaths,
				filepath.Join(binaryDir, baseSearchPath),
			)
		}
	}

	for _, parent := range searchPaths {
		found, err := filepath.Abs(filepath.Join(parent, path))
		if err != nil {
			continue
		} else if fileInfo, err := os.Stat(found); err == nil {
			if filter != nil {
				if filter(fileInfo) {
					return found
				}
			} else {
				return found
			}
		}
	}

	return ""
}

func GetLogFileLocation(fileLocation string) string {
	if fileLocation == "" {
		fileLocation, _ = FindDir("logs")
	}

	return filepath.Join(fileLocation, LOG_FILENAME)
}

func FindDir(dir string) (string, bool) {
	found := FindPath(dir, commonBaseSearchPaths, func(fileInfo os.FileInfo) bool {
		return fileInfo.IsDir()
	})
	if found == "" {
		return "./", false
	}

	return found, true
}

func FindFile(path string) string {
	return FindPath(path, commonBaseSearchPaths, func(fileInfo os.FileInfo) bool {
		return !fileInfo.IsDir()
	})
}
