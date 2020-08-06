package restful

import (
	"github.com/deliangyang/redis-shake/internal/common"
	"github.com/deliangyang/redis-shake/internal/metric"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// register all rest api
func RestAPI() {
	registerMetric()           // register metric
	registerPrometheusMetric() // register prometheus metrics
	// add below if has more
}

func registerMetric() {
	utils.HttpApi.RegisterAPI("/metric", nimo.HttpGet, func([]byte) interface{} {
		return metric.NewMetricRest()
	})
}

func registerPrometheusMetric() {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		metric.CalcPrometheusMetrics()
		promhttp.Handler().ServeHTTP(w, req)
	})
}
