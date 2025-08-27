package common

// APIResponse is a generic standardized API response structure.
type APIResponse[T any] struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   T      `json:"data,omitempty"`
}

// NewSuccessResponse creates a standardized success response.
func NewSuccessResponse[T any](code int, data T) APIResponse[T] {
	return APIResponse[T]{
		Code:   code,
		Status: "Success",
		Data:   data,
	}
}

// NewErrorResponse creates a standardized error response.
func NewErrorResponse(code int, status string) APIResponse[any] {
	return APIResponse[any]{
		Code:   code,
		Status: status,
	}
}
