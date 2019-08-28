package kail

import "github.com/prometheus/client_golang/prometheus"

var (
	CurrentOnlines = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "k8s_current_logs",
			Help: "Current logs websocket from k8s",
		},
		[]string{"namespace","pod","container"},
	)

	LogBufferTraffic = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "k8s_log_traffic",
			Help:"k8s container log traffic",
		},[]string{"namespace","pod","container"})
)
func init() {
	prometheus.DefaultRegisterer.MustRegister(CurrentOnlines, LogBufferTraffic)
}
