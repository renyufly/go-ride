package main

import "go-ride/shared/types"

// 储存 前端通过payload发送过来的各种数据 type
// 参考前端契约：web\src\contracts.ts
type previewTripRequest struct {
	// 大写字母表示public属性
	UserID string	`json:"userID"`
	Pickup types.Coordinate	`json:"pickup"`
	Destination types.Coordinate	`json:"destination"`
} 