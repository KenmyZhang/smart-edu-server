package route

import (
	"github.com/gin-gonic/gin"
)

func RouteWrapper(handle func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	handlers = append([]gin.HandlerFunc{func(context *gin.Context) {
		context.Set("RELATIVE_PATH", relativePath)
	}}, handlers...)

	return handle(relativePath, handlers...)
}
