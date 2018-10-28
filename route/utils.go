package route

import "smart-edu-server/api"

func (r *Router) InitConfig() {
	r.utils = r.root.Group("/utils")
	r.utils.GET("/config", api.GetConfig)
	r.utils.GET("/version", api.GetVersionDetails)
}
