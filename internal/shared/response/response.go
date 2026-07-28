package response

type Response struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	Error     any    `json:"error,omitempty"`
	Meta      any    `json:"meta,omitempty"`
	RequestID string `json:"requestId,omitempty"`
}
