package contracts

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/isomnath/belvedere/instrumentation"

	"github.com/isomnath/belvedere/config"
	"github.com/isomnath/belvedere/log"
	"github.com/isomnath/belvedere/translator"

	"github.com/stretchr/testify/suite"
)

type BaseContractTestSuite struct {
	suite.Suite
}
type errReader int

func (suite *BaseContractTestSuite) SetupTest() {
	config.LoadBaseConfig()
	config.LoadTranslationsConfig()
	instrumentation.StartDDTracer()
	defer instrumentation.StopDDTracer()
	log.Setup()

	path := config.GetTranslationConfig().Path()
	_ = os.Mkdir(config.GetTranslationConfig().Path(), os.ModePerm)
	_ = ioutil.WriteFile(fmt.Sprintf("%s/en.json", path), []byte(`{"error_1": "some error 1", "error_2": "some error 2"}`), 0644)
}

func (suite *BaseContractTestSuite) TearDownTest() {
	_ = os.RemoveAll(config.GetTranslationConfig().Path())
}

func (suite *BaseContractTestSuite) TestCustomResponse() {
	type TestData map[string]interface{}
	rw := httptest.NewRecorder()

	CustomResponse(rw, TestData{"message": "failure occurred"}, http.StatusBadRequest)
	suite.Equal("{\"message\":\"failure occurred\"}", rw.Body.String())
	suite.Equal(http.StatusBadRequest, rw.Code)
}

func (suite *BaseContractTestSuite) TestSuccessResponse() {
	type TestData struct {
		Message string `json:"message"`
	}
	rw := httptest.NewRecorder()

	SuccessResponse(rw, TestData{Message: "successful"}, SuccessOK)
	suite.Equal("{\"success\":true,\"data\":{\"message\":\"successful\"}}", rw.Body.String())
	suite.Equal(http.StatusOK, rw.Code)
}

func (suite *BaseContractTestSuite) TestErrorResponseWithoutTranslations() {
	translator.Kill()
	errors := []string{"error_1", "error_2"}
	rw := httptest.NewRecorder()

	ErrorResponse(rw, errors, "en", ErrorBadRequest)
	suite.Equal("{\"success\":false,\"errors\":[{\"message\":\"error_1\"},{\"message\":\"error_2\"}]}", rw.Body.String())
	suite.Equal(http.StatusBadRequest, rw.Code)
}

func (suite *BaseContractTestSuite) TestErrorResponseWithTranslations() {
	translator.Initialize()
	errors := []string{"error_1", "error_2"}
	rw := httptest.NewRecorder()

	ErrorResponse(rw, errors, "en", ErrorBadRequest)
	suite.Equal("{\"success\":false,\"errors\":[{\"message\":\"some error 1\"},{\"message\":\"some error 2\"}]}", rw.Body.String())
	suite.Equal(http.StatusBadRequest, rw.Code)
}

func (suite *BaseContractTestSuite) TestErrorResponseV2WithoutTranslations() {
	translator.Kill()
	errors := []error{errors.New("error_1"), errors.New("error_2")}
	rw := httptest.NewRecorder()

	ErrorResponseV2(rw, errors, "en", ErrorBadRequest)
	suite.Equal("{\"success\":false,\"errors\":[{\"message\":\"error_1\"},{\"message\":\"error_2\"}]}", rw.Body.String())
	suite.Equal(http.StatusBadRequest, rw.Code)
}

func (suite *BaseContractTestSuite) TestErrorResponseV2WithTranslations() {
	translator.Initialize()
	errors := []error{errors.New("error_1"), errors.New("error_2")}
	rw := httptest.NewRecorder()

	ErrorResponseV2(rw, errors, "en", ErrorBadRequest)
	suite.Equal("{\"success\":false,\"errors\":[{\"message\":\"some error 1\"},{\"message\":\"some error 2\"}]}", rw.Body.String())
	suite.Equal(http.StatusBadRequest, rw.Code)
}

func (suite *BaseContractTestSuite) TestErrorResponseReturnsUntranslatedMessageWhenTranslationErrorOccurs() {
	config.LoadTranslationsConfig()
	suite.prepareTranslationFiles()
	translator.Initialize()

	errors := []string{"error_1", "error_2"}
	rw := httptest.NewRecorder()

	ErrorResponse(rw, errors, "en", ErrorBadRequest)
	suite.Equal("{\"success\":false,\"errors\":[{\"message\":\"error_1\"},{\"message\":\"error_2\"}]}", rw.Body.String())
	suite.Equal(http.StatusBadRequest, rw.Code)
}

func (suite *BaseContractTestSuite) TestUnmarshalRequestSuccessfully() {
	type TestData struct {
		ID   int64  `json:"id"`
		Data string `json:"data"`
	}
	expectedTestData := TestData{
		ID:   123,
		Data: "test",
	}

	var dest TestData
	r, _ := http.NewRequest(http.MethodPost, "/v1/test/route/1", bytes.NewBuffer([]byte(`{"id": 123, "data": "test"}`)))

	err := UnmarshalRequest(r, &dest)
	suite.NoError(err)
	suite.Equal(expectedTestData, dest)
}

func (suite *BaseContractTestSuite) TestUnmarshalRequestJSONUnmarshalError() {
	type TestData struct {
		ID   int64  `json:"id"`
		Data string `json:"data"`
	}

	var dest TestData
	r, _ := http.NewRequest(http.MethodPost, "/v1/test/route/1", bytes.NewBuffer([]byte(`{"id": "123", "data": "test"}`)))

	err := UnmarshalRequest(r, &dest)
	suite.Equal("failed to deserialize json request body to destination interface", err.Error())
}

func (suite *BaseContractTestSuite) TestUnmarshalRequestIOReaderError() {
	type TestData struct {
		ID   int64  `json:"id"`
		Data string `json:"data"`
	}

	var dest TestData
	r, _ := http.NewRequest(http.MethodPost, "/v1/test/route/1", errReader(0))

	err := UnmarshalRequest(r, &dest)
	suite.Equal("failed to read request body", err.Error())
}

func (suite *BaseContractTestSuite) prepareTranslationFiles() {
	path := config.GetTranslationConfig().Path()
	_ = os.Mkdir(config.GetTranslationConfig().Path(), os.ModePerm)
	_ = ioutil.WriteFile(fmt.Sprintf("%s/en.json", path), []byte("{\"ERROR_INTERNAL_SERVER\": \"something went wrong\"}"), 0644)
	_ = ioutil.WriteFile(fmt.Sprintf("%s/id.json", path), []byte("{\"ERROR_INTERNAL_SERVER\": \"ada yang salah\"}"), 0644)
}

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestBaseContractTestSuite(t *testing.T) {
	suite.Run(t, new(BaseContractTestSuite))
}
