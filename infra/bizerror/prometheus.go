package bizerror

//import (
//	"github.com/prometheus/client_golang/prometheus"
//)
//
//var (
//	GrpcServerBizerrorTotal = prometheus.NewCounterVec(
//		prometheus.CounterOpts{
//			Name: "grpc_server_bizerror_total",
//			Help: "grpc server bizerror, partitioned by code(grpc_service|grpc_method|bizerror_code)",
//		},
//		[]string{"grpc_service", "grpc_method", "bizerror_code"},
//	)
//)
//
//func RegisterMetrics() {
//	prometheus.MustRegister(GrpcServerBizerrorTotal)
//}
