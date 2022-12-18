package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"jellyfin-prometheus-exporter/collectors"
	"log"
	"net/http"
	"os"
)

var (
	addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

	jfUrl = os.Getenv("JF_URL")
	token = os.Getenv("TOKEN")
)

func main() {
	var r = prometheus.NewRegistry()
	var collector = collectors.NewStreamCollector(jfUrl, token)
	r.MustRegister(collector)
	flag.Parse()
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
