package route

import "github.com/KenmyZhang/smart-edu-server/api"

func (r *Router) InitKnowledgePoint() {
	r.utils = r.root.Group("/smart-edu-server/knowledge/point")
	r.utils.POST("", api.CreateKnowlegePoint)
	r.utils.GET("", api.GetKnowlegePointById)
	r.utils.GET("/list", api.GetChildKnowledgePoints)
	r.utils.GET("/multi/list", api.GetChildKnowledgePointAndChild)
}
