package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go-ride/services/trip-service/internal/domain"
	"go-ride/shared/types"
	"io"
	"net/http"

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

// 给出pickup-起点，destination-终点
// 调用外部API-osrm，找出最优路径并显示
func (s *TripService) GetRoute(ctx context.Context, pickup, destination *types.Coordinate) (*types.OsrmAPIResponse, error) {
	// GET /{service}/{version}/{profile}/{coordinates}?option=value
	// 'http://router.project-osrm.org/route/v1/driving/13.38,52.51;13.39,52.52;13.42,52.52?overview=full&geometries=geojson'
	osrmurl := fmt.Sprintf(
		"http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=full&geometries=geojson",
		pickup.Longitude, pickup.Latitude,
		destination.Longitude, destination.Latitude,
	)
	// call external API
	response, err := http.Get(osrmurl)
	if err != nil{
		return nil, fmt.Errorf("fail to fetch route from OSRM API, %v", err)
	}

	//
	defer response.Body.Close() 

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("fail to read the response, %v", err)
	}

	// 将[]byte 的 body转成 struct
	var routeResponse types.OsrmAPIResponse

	// Unmarshal parses the JSON-encoded data and stores the result
	// 注意传指针
	if err := json.Unmarshal(body, &routeResponse); err != nil {
		return nil , fmt.Errorf("fail to parse the response, %v", err)
	}

	return &routeResponse, nil
}