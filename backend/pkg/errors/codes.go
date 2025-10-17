package errors

// ErrorCode represents standardized error codes
type ErrorCode string

const (
	// Authentication & Authorization
	ErrCodeUnauthorized     ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden        ErrorCode = "FORBIDDEN"
	ErrCodeInvalidToken     ErrorCode = "INVALID_TOKEN"
	ErrCodeExpiredToken     ErrorCode = "EXPIRED_TOKEN"
	ErrCodeInvalidPassword  ErrorCode = "INVALID_PASSWORD"

	// Validation
	ErrCodeValidation       ErrorCode = "VALIDATION_ERROR"
	ErrCodeInvalidEmail     ErrorCode = "INVALID_EMAIL"
	ErrCodeInvalidURL       ErrorCode = "INVALID_URL"
	ErrCodePasswordTooWeak  ErrorCode = "PASSWORD_TOO_WEAK"
	ErrCodeInvalidInput     ErrorCode = "INVALID_INPUT"

	// Resource
	ErrCodeNotFound         ErrorCode = "NOT_FOUND"
	ErrCodeAlreadyExists    ErrorCode = "ALREADY_EXISTS"
	ErrCodeConflict         ErrorCode = "CONFLICT"

	// Rate Limiting
	ErrCodeRateLimitExceeded ErrorCode = "RATE_LIMIT_EXCEEDED"
	ErrCodeTooManyRequests   ErrorCode = "TOO_MANY_REQUESTS"

	// Internal
	ErrCodeInternal         ErrorCode = "INTERNAL_ERROR"
	ErrCodeDatabaseError    ErrorCode = "DATABASE_ERROR"
	ErrCodeCacheError       ErrorCode = "CACHE_ERROR"

	// Business Logic
	ErrCodeInvalidOperation ErrorCode = "INVALID_OPERATION"
	ErrCodeInsufficientPermission ErrorCode = "INSUFFICIENT_PERMISSION"
)
