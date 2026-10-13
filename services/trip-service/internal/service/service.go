package service

import (
	"context"
	"go-ride/services/trip-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type TripService struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *TripService{
	return &TripService{
		repo: repo,
	}
}

func (s *TripService) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error){
	t := &domain.TripModel{
		ID: primitive.NewObjectID(),
		UserID: fare.UserID,
		Status: "pending",
		RideFare: fare,
	}
	
	return s.repo.CreateTrip(ctx, t)
}