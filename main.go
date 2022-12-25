package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"jellyfin-prometheus-exporter/api"
	"jellyfin-prometheus-exporter/collectors"
	"jellyfin-prometheus-exporter/service"
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
	jfa := api.NewJellyfinApi(jfUrl, token)
	serv := service.RelativeFileSizeService{Api: jfa}

	var r = prometheus.NewRegistry()
	var collector = collectors.NewStreamCollector(jfa)
	r.MustRegister(collector)
	flag.Parse()
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	http.HandleFunc("/rel-sizes", serv.GetRelativeFileSizes)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
