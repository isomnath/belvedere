package contracts

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/isomnath/belvedere/log"
	"github.com/isomnath/belvedere/translator"
)

type errorList struct {
	Message string `json:"message,omitempty"`
}

type BaseResponse struct {
	Success bool        `json:"success"`
	Errors  []errorList `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ReadRequestBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Log.Errorf(r.Context(), "failed to read request body with error: %v", err)
		return nil, errors.New("failed to read request body")
	}
	return body, nil
}

func UnmarshalRequest(r *http.Request, destination interface{}) error {
	body, err := ReadRequestBody(r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, destination)
	if err != nil {
		log.Log.Errorf(r.Context(), "failed to deserialize json request body to destination interface with error: %v", err)
		return errors.New("failed to deserialize json request body to destination interface")
	}
	return nil
}

func CustomResponse(rw http.ResponseWriter, responseData map[string]interface{}, statusCode int) {
	rw.WriteHeader(statusCode)

	responseJSON, _ := json.Marshal(responseData)
	rw.Write(responseJSON)

	return
}

func SuccessResponse(rw http.ResponseWriter, responseData interface{}, status string) {
	successDetails := successObjects[status]

	rw.WriteHeader(successDetails.status)
	response := BaseResponse{
		Success: true,
		Data:    responseData,
	}

	responseJSON, _ := json.Marshal(response)
	rw.Write(responseJSON)

	return
}

func ErrorResponse(rw http.ResponseWriter, errors []string, locale string, status string) {
	httpStatus := errorObjects[status].status
	rw.WriteHeader(httpStatus)

	var errorSlice []errorList

	for _, error := range errors {
		translatedMessage := translator.Translate(error, locale)
		errorSlice = append(errorSlice, errorList{Message: translatedMessage})
	}

	response := BaseResponse{
		Success: false,
		Errors:  errorSlice,
	}

	responseJSON, _ := json.Marshal(response)
	rw.Write(responseJSON)
	return
}

func ErrorResponseV2(rw http.ResponseWriter, errors []error, locale string, status string) {
	httpStatus := errorObjects[status].status
	rw.WriteHeader(httpStatus)

	var errorSlice []errorList

	for _, error := range errors {
		translatedMessage := translator.Translate(error.Error(), locale)
		errorSlice = append(errorSlice, errorList{Message: translatedMessage})
	}

	response := BaseResponse{
		Success: false,
		Errors:  errorSlice,
	}

	responseJSON, _ := json.Marshal(response)
	rw.Write(responseJSON)
	return
}
