package route

import "github.com/KenmyZhang/smart-edu-server/api"

func (r *Router) InitQuestionManger() {
	r.utils = r.root.Group("/smart-edu-server")
	r.utils.POST("/question", api.CreateQuestion)
	r.utils.GET("/question/list", api.GetQuestions)
}
