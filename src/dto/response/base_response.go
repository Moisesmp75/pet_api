package response

import "pet_api/src/common"

type BaseResponse[T any] struct {
	Success  bool     `json:"success"`
	Message  []string `json:"message"`
	Resource T        `json:"resource"`
}

type BaseResponsePag[T any] struct {
	Success    bool              `json:"success"`
	Message    []string          `json:"message"`
	Resource   T                 `json:"resource"`
	Pagination common.Pagination `json:"meta,omitempty"`
}

func NewResponse[T interface{}](resource T) BaseResponse[T] {
	return BaseResponse[T]{
		Success:  true,
		Message:  nil,
		Resource: resource,
	}
}

func NewResponsePagination[T interface{}](resource T, meta common.Pagination) BaseResponsePag[T] {
	return BaseResponsePag[T]{
		Success:    true,
		Message:    nil,
		Resource:   resource,
		Pagination: meta,
	}
}

func ErrorResponse(message string) BaseResponse[interface{}] {
	return BaseResponse[interface{}]{
		Success:  false,
		Message:  []string{message},
		Resource: nil,
	}
}

func ErrorsResponse(messages []string) BaseResponse[interface{}] {
	return BaseResponse[interface{}]{
		Success:  false,
		Message:  messages,
		Resource: nil,
	}
}

func MessageResponse[T interface{}](message string, resource T) BaseResponse[T] {
	return BaseResponse[T]{
		Success:  true,
		Message:  []string{message},
		Resource: resource,
	}
}
