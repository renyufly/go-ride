package main

import (
	"encoding/json"
	"go-ride/shared/contracts"
	"net/http"
)

// 所有HTTP Handler 相关的代码

// 处理 Trip Preview 的 handler
func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// 检测前端是否为 post request

	}
	
	var requestBody previewTripRequest
	// 反序列化
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "fail to parse json data", http.StatusBadRequest)
		return
	}

	// Go语言中所有带Close方法都需要手动调用
	defer r.Body.Close()

	// Validation
	if requestBody.UserID == "" {
		http.Error(w, "userID cannot be empty", http.StatusBadRequest)
		return
	}

	// TODO: call trip service

	// 使用contract规定的前端希望的后端response格式
	response := contracts.APIResponse{
		Data: "ok",
	}

	writeJSON(w, http.StatusCreated, response)

}