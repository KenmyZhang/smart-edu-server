package route

import (
	"github.com/KenmyZhang/golang-lib/middleware"
)

func (r *Router) InitPrometheus() {
	r.prometheus = r.root.Group("")
	r.prometheus.GET(middleware.DefaultMetricPath, middleware.GetMetrics)
}
