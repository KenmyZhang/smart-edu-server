package common

import (
	"time"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	historyBuckets = [...]float64{
		10., 20., 30., 50., 80., 100., 200., 300., 500., 1000., 2000., 3000.}
	DefaultMetricPath = "/metrics"

	ResponseCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "smart-edu-server_requests_total",
		Help: "Total request counts"}, []string{"method", "endpoint"})
	ErrorCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "smart-edu-server_error_total",
		Help: "Total Error counts"}, []string{"method", "endpoint"})
	ResponseLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "smart-edu-server_response_latency_millisecond",
		Help:    "Response latency (millisecond)",
		Buckets: historyBuckets[:]}, []string{"method", "endpoint"})

	// CacheLatency : cache latency for promethues
	CacheLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "smart-edu-server_cache_latency_millisecond",
		Help:    "cache latency (millisecond)",
		Buckets: historyBuckets[:]}, []string{"method"})
	// CacheCallCounter : cache counter for promethues
	CacheCallCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "smart-edu-server_cache_total",
		Help: "Total cache call counts"}, []string{"method"})
	// CacheErrorCounter : cache error counter for promethues
	CacheErrorCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "smart-edu-server_cache_error_total",
		Help: "Total cache error counts"}, []string{"method"})

	OfflineSpaceResponseCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "smart-edu-server_offline_space_requests_total",
		Help: "Total offline space request counts"}, []string{"method", "endpoint"})

	OfflineSpaceErrorCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "smart-edu-server_offline_space_error_total",
		Help: "Total offline space error counts"}, []string{"method", "endpoint"})
	OfflineSpaceResponseLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "smart-edu-server_offline_space_response_latency_millisecond",
		Help:    "Offline space Response latency (millisecond)",
		Buckets: historyBuckets[:]}, []string{"method", "endpoint"})
)

func init() {

	fmt.Println("prometheus_init .... ")
	prometheus.MustRegister(ResponseCounter)
	prometheus.MustRegister(ErrorCounter)
	prometheus.MustRegister(ResponseLatency)
	prometheus.MustRegister(CacheCallCounter)
	prometheus.MustRegister(CacheErrorCounter)
	prometheus.MustRegister(CacheLatency)
	prometheus.MustRegister(OfflineSpaceResponseCounter)
	prometheus.MustRegister(OfflineSpaceErrorCounter)
	prometheus.MustRegister(OfflineSpaceResponseLatency)
}

func Prometheus() gin.HandlerFunc {
	return func(c *gin.Context) {
		endPoint := c.Request.URL.Path
		if endPoint == DefaultMetricPath {
			c.Next()
		} else {
			start := time.Now()
			method := c.Request.Method

			c.Next()
			relativePath := c.GetString("RELATIVE_PATH")
			if relativePath != "" {
				endPoint = relativePath
			}

			statusCode := c.Writer.Status()
			if statusCode != http.StatusNotFound {
				elapsed := float64(time.Since(start).Nanoseconds()) / 1000000
				ResponseCounter.WithLabelValues(method, endPoint).Inc()
				ResponseLatency.WithLabelValues(method, endPoint).Observe(elapsed)
			} else {
				elapsed := float64(time.Since(start).Nanoseconds()) / 1000000
				ResponseCounter.WithLabelValues(method, "[!othersPath!]").Inc()
				ResponseLatency.WithLabelValues(method, "[!othersPath!]").Observe(elapsed)
			}
		}
	}
}

func LatestMetrics(c *gin.Context) {
	handler := promhttp.Handler()
	handler.ServeHTTP(c.Writer, c.Request)
}
