package shopee

import "fmt"

// Sentinel errors for common Shopee API error codes.
// Use errors.Is to check for specific error types.
var (
	ErrAuth          = &APIError{ErrorCode: "error_auth"}
	ErrSign          = &APIError{ErrorCode: "error_sign"}
	ErrParam         = &APIError{ErrorCode: "error_param"}
	ErrRateLimit     = &APIError{ErrorCode: "error_rate_limit"}
	ErrLimit         = &APIError{ErrorCode: "error_limit"}
	ErrServer        = &APIError{ErrorCode: "error_server"}
	ErrPermission    = &APIError{ErrorCode: "error_permission"}
	ErrSuspended     = &APIError{ErrorCode: "api_suspended"}
	ErrShopBanned    = &APIError{ErrorCode: "shop_banned"}
	ErrShopNotLinked = &APIError{ErrorCode: "shop_no_linked"}
)

// APIError represents a Shopee API error response.
type APIError struct {
	ErrorCode string
	Message   string
	RequestID string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("shopee api error [%s]: %s (request_id: %s)", e.ErrorCode, e.Message, e.RequestID)
}

// Is reports whether target matches this APIError by error code.
// This allows using errors.Is(err, ErrAuth) to check for specific error types.
func (e *APIError) Is(target error) bool {
	if t, ok := target.(*APIError); ok {
		// Match by error code only (Message and RequestID are dynamic)
		if t.ErrorCode != "" {
			return e.ErrorCode == t.ErrorCode
		}
	}
	return false
}

// IsAuthError returns true if the error is an authentication/authorization error.
func (e *APIError) IsAuthError() bool {
	return e.ErrorCode == "error_auth" || e.ErrorCode == "invalid_acceess_token"
}

// IsRateLimitError returns true if the error is a rate limiting error.
func (e *APIError) IsRateLimitError() bool {
	return e.ErrorCode == "error_rate_limit" || e.ErrorCode == "error_limit"
}

// IsServerError returns true if the error is a server-side error.
func (e *APIError) IsServerError() bool {
	return e.ErrorCode == "error_server" || e.ErrorCode == "error_network"
}
