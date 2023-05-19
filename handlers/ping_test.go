package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PingHandlerTestSuite struct {
	suite.Suite
}

func (suite *PingHandlerTestSuite) TestPing() {
	rw := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/ping", nil)

	Ping(rw, r)

	suite.Nil(err)
	suite.Equal("{\"success\":true,\"data\":{\"message\":\"pong\"}}", rw.Body.String())
	suite.Equal(http.StatusOK, rw.Code)
}

func TestPingHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(PingHandlerTestSuite))
}
