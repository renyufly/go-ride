package contracts

// 让前端可以预期后端会返回什么
// 即前端会接收的response 格式

// APIResponse is the response structure for the API.
type APIResponse struct {
	Data  any       `json:"data,omitempty"`
	Error *APIError `json:"error,omitempty"`
}

// APIError is the error structure for the API.
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
