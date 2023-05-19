package contracts

import (
	"net/http"
)

type errorDetails struct {
	message string
	status  int
}

var errorObjects = map[string]errorDetails{
	ErrorBadRequest: {
		status: http.StatusBadRequest,
	},
	ErrorUnauthorized: {
		status: http.StatusUnauthorized,
	},
	ErrorForbidden: {
		status: http.StatusForbidden,
	},
	ErrorNotFound: {
		status: http.StatusNotFound,
	},
	ErrorMethodNotAllowed: {
		status: http.StatusMethodNotAllowed,
	},
	ErrorNotAcceptable: {
		status: http.StatusNotAcceptable,
	},
	ErrorRequestTimeout: {
		status: http.StatusRequestTimeout,
	},
	ErrorConflict: {
		status: http.StatusConflict,
	},
	ErrorRequestEntityTooLarge: {
		status: http.StatusRequestEntityTooLarge,
	},
	ErrorUnprocessableEntity: {
		status: http.StatusUnprocessableEntity,
	},
	ErrorUpgradeRequired: {
		status: http.StatusUpgradeRequired,
	},
	ErrorTooManyRequests: {
		status: http.StatusTooManyRequests,
	},
	ErrorInternalServer: {
		status: http.StatusInternalServerError,
	},
	ErrorServiceUnavailable: {
		status: http.StatusServiceUnavailable,
	},
}
