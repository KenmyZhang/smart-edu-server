package route

import "github.com/KenmyZhang/smart-edu-server/api"

func (r *Router) InitUser() {
	r.utils = r.root.Group("/smart-edu-server")
	r.utils.GET("/user/:user_id", api.GetUser)
	r.utils.POST("/user", api.CreateUser)
}
