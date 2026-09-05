package domain

// 有不同类型的费用标准
// 根据距离和价格,我们需要为单次行程计算出所有费用

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 要存入数据库的模型结构 - Model

type TripModel struct{
	ID primitive.ObjectID  // 方便存储于MongoDB
	UserID string    // 一个trip行程绑定一个User(乘客)
	Status string
	RideFare RideFareModel
}

// 为 仓库层(DAO) 和 服务层(业务代码) 定义接口规范
type TripRepository interface{
	// 若与任何数据库驱动通信,将context从传输层一直传递到仓库层是合理做法
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)

}

type TripService interface{
	// 依据fare创建trip，注意传入的是pointer
	// pointer: 修改原来的对象，而不是修改副本
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
}