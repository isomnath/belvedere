package contracts

/*
	Error Constants
*/
// Client Errors - 4xx Errors
const (
	ErrorBadRequest            = "ERROR_BAD_REQUEST"
	ErrorUnauthorized          = "ERROR_UNAUTHORIZED"
	ErrorForbidden             = "ERROR_FORBIDDEN"
	ErrorNotFound              = "ERROR_NOT_FOUND"
	ErrorMethodNotAllowed      = "ERROR_METHOD_NOT_ALLOWED"
	ErrorNotAcceptable         = "ERROR_NOT_ACCEPTABLE"
	ErrorRequestTimeout        = "ERROR_REQUEST_TIMEOUT"
	ErrorConflict              = "ERROR_CONFLICT"
	ErrorRequestEntityTooLarge = "ERROR_REQUEST_ENTITY_TOO_LARGE"
	ErrorUnprocessableEntity   = "ERROR_UNPROCESSABLE_ENTITY"
	ErrorUpgradeRequired       = "ERROR_UPGRADE_REQUIRED"
	ErrorTooManyRequests       = "ERROR_TOO_MANY_REQUESTS"
)

// Server Errors - 5xx
const (
	ErrorInternalServer     = "ERROR_INTERNAL_SERVER"
	ErrorServiceUnavailable = "ERROR_SERVICE_UNAVAILABLE"
)

/*
Success Constants
*/
const (
	SuccessOK             = "SUCCESS_OK"
	SuccessCreated        = "SUCCESS_CREATED"
	SuccessNoContent      = "SUCCESS_NO_CONTENT"
	SuccessPartialContent = "SUCCESS_PARTIAL_CONTENT"
)
