package metrics

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	registerOnce sync.Once

	apiRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "currencyapi",
			Name:      "api_requests_total",
			Help:      "Total number of API requests.",
		},
		[]string{"result"},
	)
)

func Register() {
	registerOnce.Do(func() {
		prometheus.MustRegister(
			apiRequestsTotal,
		)
	})
}

func Handler() http.Handler {
	Register()
	return promhttp.Handler()
}

func GinMiddleware() gin.HandlerFunc {
	Register()

	return func(c *gin.Context) {
		if c.FullPath() == "/metrics" {
			c.Next()
			return
		}

		c.Next()

		result := "success"
		if c.Writer.Status() >= 400 {
			result = "error"
		}
		apiRequestsTotal.WithLabelValues(result).Inc()
	}
}
