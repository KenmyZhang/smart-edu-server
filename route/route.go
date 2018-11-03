package route

import (
	"smart-edu-server/common"
	"smart-edu-server/common/config"

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
	router.Use(common.Logger(common.DefaultMetricPath))
	router.Use(common.Prometheus())
	router.Use(common.Recovery())

	BaseRouter := &Router{root: router}
	BaseRouter.InitPrometheus()
	BaseRouter.InitConfig()
	BaseRouter.InitUser()
	BaseRouter.InitArticle()
	return router
}
