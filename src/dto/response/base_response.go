package response

import "pet_api/src/common"

type BaseResponse[T any] struct {
	Success  bool     `json:"success"`
	Message  []string `json:"message"`
	Resource T        `json:"resource"`
	// Pagination *common.Pagination `json:"meta"`
}

type BaseResponsePag[T any] struct {
	Success    bool               `json:"success"`
	Message    []string           `json:"message"`
	Resource   T                  `json:"resource"`
	Pagination *common.Pagination `json:"meta"`
}

func NewResponse[T any](resource T) BaseResponse[T] {
	return BaseResponse[T]{
		Success:  true,
		Message:  nil,
		Resource: resource,
		// Pagination: nil,
	}
}

func NewResponsePagination[T any](resource T, meta common.Pagination) BaseResponsePag[T] {
	return BaseResponsePag[T]{
		Success:    true,
		Message:    nil,
		Resource:   resource,
		Pagination: &meta,
	}
}

func ErrorResponse(message string) BaseResponse[*any] {
	return BaseResponse[*any]{
		Success:  false,
		Message:  []string{message},
		Resource: nil,
	}
}

func ErrorsResponse(messages []string) BaseResponse[*any] {
	return BaseResponse[*any]{
		Success:  false,
		Message:  messages,
		Resource: nil,
	}
}
