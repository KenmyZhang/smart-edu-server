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

var allowedMethods []string = []string{
	"POST",
	"GET",
	"OPTIONS",
	"PUT",
	"PATCH",
	"DELETE",
}

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
	router.Use(middleware.Cors(config.Cfg.AllowCorsFrom, allowedMethods))

	BaseRouter := &Router{root: router}
	BaseRouter.InitPrometheus()
	BaseRouter.InitConfig()
	BaseRouter.InitUser()
	BaseRouter.InitArticle()
	BaseRouter.InitQuestionManger()
	BaseRouter.InitKnowledgePoint()
	return router
}
