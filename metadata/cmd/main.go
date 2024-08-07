package main

import (
	"context"
	"flag"
	"fmt"
	"go-microservice/metadata/internal/controller/metadata"
	httpHandler "go-microservice/metadata/internal/handler/http"
	"go-microservice/metadata/internal/repository/memory"
	"go-microservice/pkg/discovery"
	"go-microservice/pkg/discovery/consul"
	"log"
	"net/http"
	"time"
)

const serviceName = "metadata"

func main() {

	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("Starting the metadata service on port%d", port)

	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)

	if err := registry.Register(ctx, instanceID,
		serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.
				ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthystate: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()

	defer registry.Deregister(ctx, instanceID, serviceName)

	repo := memory.NewRepository()
	ctrl := metadata.New(repo)
	handler := httpHandler.NewHandler(ctrl)

	http.Handle("/metadata", http.HandlerFunc(handler.GetMetadata))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
