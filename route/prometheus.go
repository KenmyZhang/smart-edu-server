package route

import (
	"smart-edu-server/common"
)

func (r *Router) InitPrometheus() {
	r.prometheus = r.root.Group("")
	r.prometheus.GET(common.DefaultMetricPath, common.LatestMetrics)
}
