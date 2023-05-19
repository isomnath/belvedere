package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/log"
	"github.com/isomnath/belvedere/translator"

	"github.com/stretchr/testify/suite"
)

type RouteNotFoundHandlerTestSuite struct {
	suite.Suite
}

func (suite *RouteNotFoundHandlerTestSuite) SetupSuite() {
	config.LoadBaseConfig()
	config.LoadTranslationsConfig()
	log.Setup()
	suite.prepareTranslationFiles()
	translator.Initialize()
}

func (suite *RouteNotFoundHandlerTestSuite) TearDownSuite() {
	os.RemoveAll(config.GetTranslationConfig().Path())
}

func (suite *RouteNotFoundHandlerTestSuite) TestRouteNotFoundHandlerShouldReturnNotFound() {
	rw := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/invalid_path", nil)
	suite.NoError(err, "failed to create a request")

	RouteNotFoundHandler(rw, r)

	suite.Equal(http.StatusNotFound, rw.Code)
	suite.Equal("{\"success\":false,\"errors\":[{\"message\":\"route /invalid_path not found\"}]}", rw.Body.String())
}

func (suite *RouteNotFoundHandlerTestSuite) prepareTranslationFiles() {
	path := config.GetTranslationConfig().Path()
	os.Mkdir(config.GetTranslationConfig().Path(), os.ModePerm)
	ioutil.WriteFile(fmt.Sprintf("%s/en.json", path), []byte("{\"ERROR_INTERNAL_SERVER\": \"something went wrong\"}"), 0644)
	ioutil.WriteFile(fmt.Sprintf("%s/id.json", path), []byte("{\"ERROR_INTERNAL_SERVER\": \"ada yang salah\"}"), 0644)
}

func TestRouteNotFoundHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(RouteNotFoundHandlerTestSuite))
}
