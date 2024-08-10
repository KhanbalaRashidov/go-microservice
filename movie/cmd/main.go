package main

import (
	"context"
	"flag"
	"fmt"
	"go-microservice/movie/internal/controller/movie"
	metadatagateway "go-microservice/movie/internal/gateway/metadata/http"
	ratinggateway "go-microservice/movie/internal/gateway/rating/http"
	httphandler "go-microservice/movie/internal/handler/http"
	"go-microservice/pkg/discovery"
	"go-microservice/pkg/discovery/consul"
	"log"
	"net/http"
	"time"
)

const serviceName = "movie"

func main() {
	var port int

	flag.IntVar(&port, "port", 8083, "API handler port")
	flag.Parse()
	log.Printf("Starting the movie service on port %d",
		port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.
				ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			if err := registry.
				ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()

	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	svc := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(svc)
	http.Handle("/movie", http.HandlerFunc(h.
		GetMovieDetails))
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d",
		port), nil); err != nil {
		panic(err)
	}
}
