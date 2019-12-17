package prometheus

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	prometheus2 "k8s-lx1036/app/k8s/prometheus/client-go/prometheus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Options struct {
	AppName          string
	Idc              string
	WatchPath        map[string]struct{}
	HistogramBuckets []float64
}

type Prometheus struct {
	AppName   string
	Idc       string
	WatchPath map[string]struct{}
	Counter   *prometheus2.CounterVec
	//Histogram *prometheus2.HistogramVec
}

var Metrics *Prometheus

func Init(options Options) {
	if strings.TrimSpace(options.AppName) == "" || strings.TrimSpace(options.Idc) == "" || len(options.HistogramBuckets) == 0 {
		panic(options.AppName + " or " + options.Idc + " or HistogramBuckets is empty.")
	}

	Metrics = &Prometheus{
		AppName:   options.AppName,
		Idc:       options.Idc,
		WatchPath: options.WatchPath,
		Counter: prometheus2.NewCounterVec(
			prometheus2.CounterOpts{
				Name: "module_responses",
				Help: "calculate qps",
			},
			[]string{"app", "module", "api", "method", "code", "idc"},
		),
		/*Histogram: prometheus2.NewHistogramVec(
			prometheus2.HistogramOpts{
				Namespace:   "",
				Subsystem:   "",
				Name:        "response_duration_milliseconds",
				Help:        "HTTP latency distributions",
				ConstLabels: nil,
				//Buckets:     options.HistogramBuckets,
			},
			[]string{"app", "module", "api", "method", "idc"},
		),*/
	}

	prometheus2.MustRegister(Metrics.Counter)
	//prometheus2.MustRegister(Metrics.Histogram)
}

type LatencyRecord struct {
	Time   float64
	Api    string
	Module string
	Method string
	Code   int
}

type QpsRecord struct {
	Api    string
	Module string
	Method string
	Code   int
}

func (metrics *Prometheus) LatencyLog(record LatencyRecord) {
	if strings.TrimSpace(record.Module) == "" {
		record.Module = "self"
	}

	/*metrics.Histogram.WithLabelValues(
		metrics.AppName,
		record.Module,
		record.Api,
		record.Method,
		metrics.Idc,
	).Observe(record.Time)*/
}

func (metrics *Prometheus) QpsCounterLog(record QpsRecord) {
	if strings.TrimSpace(record.Module) == "" {
		record.Module = "self"
	}

	metrics.Counter.WithLabelValues(
		metrics.AppName,
		record.Module,
		record.Api,
		record.Method,
		strconv.Itoa(record.Code),
		metrics.Idc,
	).Inc()
}

func MetricsServerStart(path string, port int) {
	go func() {
		http.Handle(path, promhttp.Handler())
		fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	}()
}

func MiddlewarePrometheusAccessLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		times := time.Now()

		context.Next()

		if _, ok := Metrics.WatchPath[context.Request.URL.Path]; ok {
			latency := float64(time.Since(times).Milliseconds())
			/*Metrics.LatencyLog(LatencyRecord{
				Time:   latency,
				Api:    context.Request.URL.Path,
				Method: context.Request.Method,
				Code:   context.Writer.Status(),
			})*/

			Metrics.QpsCounterLog(QpsRecord{
				Api:    context.Request.URL.Path,
				Method: context.Request.Method,
				Code:   context.Writer.Status(),
			})

			fmt.Println(context.Request.URL.Path, context.Request.Method, context.Writer.Status(), latency)
		}
	}
}
