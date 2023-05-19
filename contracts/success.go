package contracts

import "net/http"

type successDetails struct {
	status int
}

var successObjects = map[string]successDetails{
	SuccessOK: {
		status: http.StatusOK,
	},
	SuccessCreated: {
		status: http.StatusCreated,
	},
	SuccessNoContent: {
		status: http.StatusNoContent,
	},
	SuccessPartialContent: {
		status: http.StatusPartialContent,
	},
}
