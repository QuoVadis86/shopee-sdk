package shopee

import "fmt"

// APIError represents a Shopee API error response.
type APIError struct {
	ErrorCode string
	Message   string
	RequestID string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("shopee api error [%s]: %s (request_id: %s)", e.ErrorCode, e.Message, e.RequestID)
}
