package handlers

import (
	"net/http"

	"github.com/isomnath/belvedere/contracts"
)

// Ping - Application Health Check API
// @Summary PING
// @Description Check Application Liveliness
// @Tags Health Check
// @Produce json
// @Success 200 {object} base.PingResponse
// @Router /ping [get]
func Ping(rw http.ResponseWriter, r *http.Request) {
	res := contracts.PingResponse{Message: "pong"}
	contracts.SuccessResponse(rw, res, contracts.SuccessOK)
}
