package web

import (
	"net/http"
	"testing"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/instrumentation"

	"github.com/stretchr/testify/suite"
)

type WebRouterTestSuite struct {
	suite.Suite
}

func (suite *WebRouterTestSuite) SetupSuite() {
	config.LoadBaseConfig()
	instrumentation.StartDDTracer()
}

func (suite *WebRouterTestSuite) TearDownSuite() {
	instrumentation.StopDDTracer()
}

func MyTestHandlerFunc(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(`{"success": true, "data": {"message": "successful"}}`))
}

func (suite *WebRouterTestSuite) TestInitialize() {
	routesWithoutMiddleware := []Route{{
		Name:            "Test Route",
		Path:            "/test_route_1",
		Middlewares:     nil,
		HandlerFunction: MyTestHandlerFunc,
		Method:          http.MethodGet,
	}}

	routesWithMiddleware := []Route{{
		Name:            "Test Route",
		Path:            "/test_route_2",
		Middlewares:     []Middleware{{Function: myTestMiddlewareFuncOne}, {Function: myTestMiddlewareFuncTwo}},
		HandlerFunction: MyTestHandlerFunc,
		Method:          http.MethodPut,
	}}

	suite.NotNil(Initialize(routesWithoutMiddleware))
	suite.NotNil(Initialize(routesWithMiddleware))

}

func myTestMiddlewareFuncOne(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		next(rw, r)
	}
}

func myTestMiddlewareFuncTwo(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		next(rw, r)
	}
}

func TestWebRouterTestSuite(t *testing.T) {
	suite.Run(t, new(WebRouterTestSuite))
}
