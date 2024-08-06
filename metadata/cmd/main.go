package main

import (
	"go-microservice/metadata/internal/controller/metadata"
	httpHandler "go-microservice/metadata/internal/handler/http"
	"go-microservice/metadata/internal/repository/memory"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the movie metadata service")
	repo := memory.NewRepository()
	ctrl := metadata.New(repo)
	handler := httpHandler.NewHandler(ctrl)

	http.Handle("/metadata", http.HandlerFunc(handler.GetMetadata))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
