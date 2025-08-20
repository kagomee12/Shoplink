package dto

type ApiResponse[T any] struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
	Data     T      `json:"data,omitempty"`
}
