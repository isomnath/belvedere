package console

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/contracts"
	"github.com/isomnath/belvedere/instrumentation"
	"github.com/isomnath/belvedere/log"
	"github.com/isomnath/belvedere/router/web"

	"github.com/stretchr/testify/suite"
	ddTracerMux "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

type StartServerTestSuite struct {
	suite.Suite
	urlOne   string
	urlTwo   string
	urlThree string
	router   *ddTracerMux.Router
}

func (suite *StartServerTestSuite) SetupTest() {
	config.LoadBaseConfig()
	config.LoadTranslationsConfig()
	instrumentation.StartDDTracer()
	log.Setup()

	suite.urlOne = "/v1/test/route/1"
	suite.urlTwo = "/v1/test/route/2"
	suite.urlThree = "/v1/test/route/3"

	routes := []web.Route{
		{
			Name:            "route_one",
			Path:            suite.urlOne,
			Middlewares:     []web.Middleware{{Function: myTestMiddlewareFuncOne}},
			HandlerFunction: suite.routeOne,
			Method:          http.MethodGet,
		},
		{
			Name:            "route_two",
			Path:            suite.urlTwo,
			Middlewares:     []web.Middleware{{Function: myTestMiddlewareFuncTwo}},
			HandlerFunction: suite.routeTwo,
			Method:          http.MethodGet,
		},
		{
			Name:            "route_three",
			Path:            suite.urlThree,
			Middlewares:     []web.Middleware{{Function: myTestMiddlewareFuncThree}},
			HandlerFunction: suite.routeThree,
			Method:          http.MethodGet,
		},
	}

	suite.router = web.Initialize(routes)
	go StartServer(suite.router)
	time.Sleep(2 * time.Second)
}

func (suite *StartServerTestSuite) TearDownTest() {
	instrumentation.StopDDTracer()
}

func (suite *StartServerTestSuite) TestStartServerSuccess() {
	responsePing, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/%s", config.GetAppWebPort(), "/ping"))
	suite.Equal([]string{"application/json"}, responsePing.Header["Content-Type"])
	suite.NoError(err)

	responseOne, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/%s", config.GetAppWebPort(), suite.urlOne))
	suite.Equal([]string{"application/json"}, responseOne.Header["Content-Type"])
	suite.NoError(err)

	responseTwo, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/%s", config.GetAppWebPort(), suite.urlTwo))
	suite.Equal([]string{"application/json"}, responseTwo.Header["Content-Type"])
	suite.NoError(err)

	bodyPing, err := ioutil.ReadAll(responsePing.Body)
	suite.NoError(err)

	bodyOne, err := ioutil.ReadAll(responseOne.Body)
	suite.NoError(err)

	bodyTwo, err := ioutil.ReadAll(responseTwo.Body)
	suite.NoError(err)

	suite.NoError(err)
	suite.Equal("{\"success\":true,\"data\":{\"message\":\"pong\"}}", string(bodyPing))

	suite.NoError(err)
	suite.Equal("{\"success\":true,\"data\":{\"message\":\"route one successful\"}}", string(bodyOne))

	suite.NoError(err)
	suite.Equal("{\"success\":true,\"data\":{\"message\":\"route two successful\"}}", string(bodyTwo))
}

func (suite *StartServerTestSuite) TestStartServerPanicRecovery() {
	responseThree, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/%s", config.GetAppWebPort(), suite.urlThree))
	suite.Equal([]string{"application/json"}, responseThree.Header["Content-Type"])
	suite.NoError(err)

	_, err = ioutil.ReadAll(responseThree.Body)

	suite.NoError(err)
}

func (suite *StartServerTestSuite) routeOne(rw http.ResponseWriter, _ *http.Request) {
	response := TestStruct{Message: "route one successful"}
	contracts.SuccessResponse(rw, response, contracts.SuccessOK)
}

func (suite *StartServerTestSuite) routeTwo(rw http.ResponseWriter, _ *http.Request) {
	response := TestStruct{Message: "route two successful"}
	contracts.SuccessResponse(rw, response, contracts.SuccessOK)
}

func (suite *StartServerTestSuite) routeThree(_ http.ResponseWriter, _ *http.Request) {
	panic(errors.New("test error"))
}

func myTestMiddlewareFuncOne(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		log.Log.Infof(ctx, "test middleware 1")
		next(rw, r)
	}
}

func myTestMiddlewareFuncTwo(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		log.Log.Infof(ctx, "test middleware 2")
		next(rw, r)
	}
}

func myTestMiddlewareFuncThree(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		log.Log.Infof(ctx, "test middleware 3")
		next(rw, r)
	}
}

type TestStruct struct {
	Message string `json:"message"`
}

func TestStartWebServerTestSuite(t *testing.T) {
	suite.Run(t, new(StartServerTestSuite))
}
