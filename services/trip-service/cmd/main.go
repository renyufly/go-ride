package main

import (
	"context"
	"go-ride/services/trip-service/internal/domain"
	h "go-ride/services/trip-service/internal/infrastructure/http"
	"go-ride/services/trip-service/internal/infrastructure/repository"
	"go-ride/services/trip-service/internal/service"
	"log"
	"net/http"
)

func main() {

	ctx := context.Background()

	// TODO: 后续替换为MongoDB
	inmemRepo := repository.NewInmemRepository()
	tripservice := service.NewService(inmemRepo)

	mux := http.NewServeMux()

	httphandler := h.HttpHandler{
		Service: tripservice,
	}

	mux.HandleFunc("POST /preview", httphandler.HandleTripPreview)

	server := &http.Server{
		Addr: ":8083",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP Server error: %v", err)
	}

}
