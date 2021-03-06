package route

import "github.com/KenmyZhang/smart-edu-server/api"

func (r *Router) InitArticle() {
	r.utils = r.root.Group("/smart-edu-server")
	r.utils.GET("/article/:article_id", api.GetArticle)
}
