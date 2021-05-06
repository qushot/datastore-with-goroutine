package main

import (
	"log"
	"net/http"
	"os"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"

	"github.com/qushot/datastore-with-goroutine/handler"
)

func main() {
	sd, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT"),
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(sd)
	// AlwaysSampleを使うと全てのリクエストをトレースする。検証時以外は使わないほうがいいかも。
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	s := http.NewServeMux()
	s.HandleFunc("/", handler.Index)
	s.HandleFunc("/async", handler.Async)
	s.HandleFunc("/sync", handler.Sync)
	s.HandleFunc("/setup", handler.SetUp)
	s.HandleFunc("/teardown", handler.TearDown)

	h := &ochttp.Handler{
		Propagation: &propagation.HTTPFormat{},
		Handler:     s,
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, h); err != nil {
		log.Fatal(err)
	}
}
