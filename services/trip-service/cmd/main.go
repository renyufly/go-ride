package main

import (
	"context"
	"go-ride/services/trip-service/internal/domain"
	"go-ride/services/trip-service/internal/infrastructure/repository"
	"go-ride/services/trip-service/internal/service"
)

func main() {

	ctx := context.Background()

	// TODO: 后续替换为MongoDB
	inmemRepo := repository.NewInmemRepository()
	tripservice := service.NewService(inmemRepo)

	fare_mock := domain.RideFareModel{
		UserID: "42",
	}

	tripservice.CreateTrip(ctx, &fare_mock)

}