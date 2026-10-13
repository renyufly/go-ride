package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 所有与票价相关的内容

type RideFareModel struct{
	ID primitive.ObjectID
	UserID string
	PackageSlug string  // 选择乘车级别：商务、豪华、
	TotalPriceInCents float64  // 根据距离等因素计算出的车费价格 (按美分定价)
	ExpiresAt time.Time   // 过期时间
}