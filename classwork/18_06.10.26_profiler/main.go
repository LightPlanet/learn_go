// go tool pprof -http=":9090" -seconds=30 http://localhost:8080/debug/pprof/profile
// go tool pprof -seconds=30 http://localhost:8080/debug/pprof/profile

// go get github.com/prometheus/client_golang/prometheus
// go get github.com/prometheus/client_golang/prometheus/promhttp

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Profiler -----------------------------------------------------------------------------

const maxSize = 10_000_000

func appendSlice() {
	for {
		var s []int
		for i := 0; i < maxSize; i++ {
			s = append(s, i)
		}
	}
}

func ProfilerExample() {
	go appendSlice()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
		return
	}
}

// Prometheus ---------------------------------------------------------------------------

var httpTotalRequest = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "total_requests",
		Help: "Total HTTP requests",
	},
	[]string{"path"},
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	httpTotalRequest.WithLabelValues(r.URL.Path).Inc()
	fmt.Fprintf(w, "INDEX PAGE")
}

func PrometheusExample() {
	prometheus.MustRegister(httpTotalRequest)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", indexHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
		return
	}
}

func main() {
	// ProfilerExample()
	PrometheusExample()
}
