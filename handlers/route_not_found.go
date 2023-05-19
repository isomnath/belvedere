package handlers

import (
	"fmt"
	"net/http"

	"github.com/isomnath/belvedere/contracts"
)

// RouteNotFoundHandler - All requests are redirected to this router when a request is received for an unsupported path
func RouteNotFoundHandler(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	errors := []string{fmt.Sprintf("route %s not found", path)}
	contracts.ErrorResponse(rw, errors, "en", contracts.ErrorNotFound)
}
