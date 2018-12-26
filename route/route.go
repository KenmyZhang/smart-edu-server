package route

import (
	"github.com/KenmyZhang/golang-lib/middleware"
	"github.com/KenmyZhang/smart-edu-server/common/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	root            *gin.Engine
	smartEduServer  *gin.RouterGroup
	customerService *gin.RouterGroup
	utils           *gin.RouterGroup
	prometheus      *gin.RouterGroup
}

var BaseRouter *Router

func NewRoute() *gin.Engine {
	router := gin.New()
	if config.Cfg.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	router.Use(middleware.Logger(middleware.DefaultMetricPath))
	router.Use(middleware.Prometheus())
	router.Use(middleware.Recovery())

	BaseRouter := &Router{root: router}
	BaseRouter.InitPrometheus()
	BaseRouter.InitConfig()
	BaseRouter.InitUser()
	BaseRouter.InitArticle()
	BaseRouter.InitQuestionManger()
	return router
}
