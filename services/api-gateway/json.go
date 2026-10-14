package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	// 设置 response header
	w.Header().Set("Content-Type", "application/json")
	// 写入 状态码 status
	w.WriteHeader(status)
	
	return json.NewEncoder(w).Encode(data)
}