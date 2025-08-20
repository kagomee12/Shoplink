package pkg

import (
	"shoplink/app/constant"
	"shoplink/app/domain/dto"
)

func BuildResponse[T any](status constant.ResponseStatus, data T) dto.ApiResponse[T] {
	return BuildResponse_(status.GetResponseStatus(), status.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, messages string, data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		Status:   status,
		Messages: messages,
		Data:     data,
	}
}