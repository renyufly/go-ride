package repository

import (
	"context"
	"go-ride/services/trip-service/internal/domain"
)

// inmem (in memory): 内存存储的具体实现

type inmemRepository struct{
	// key-string , value-*domain.TripModel
	// 直接从map中通过ID获取Trip
	trips map[string]*domain.TripModel
	rideFares map[string]*domain.RideFareModel
}

// constructor-构造函数
func NewInmemRepository () *inmemRepository{
	return &inmemRepository{
		// 创建 empty map
		trips: make(map[string]*domain.TripModel),
		rideFares: make(map[string]*domain.RideFareModel),
	}
}

// implement the interface
// r(reveiver): 给 *inmemRepository 类型定义一个 CreateTrip 方法
// 在方法内部通过 r 访问这个 repository
func (r *inmemRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	// .Hex() 把 primitive.ObjectID 转成 string
	r.trips[trip.ID.Hex()] = trip
	return trip, nil
}	
