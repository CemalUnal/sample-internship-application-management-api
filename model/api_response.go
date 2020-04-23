package model

import (
	"net/http"
)

// ApiResponse is used to return data from the api
// it provides a unified model for all api responses
type ApiResponse struct {
	Code    	int 		`json:"code"`
	Message		string 		`json:"message"`
	Data 		interface{}	`json:"data"`
}

func GetOkResponse(w http.ResponseWriter, message string, data interface{}) (ApiResponse, http.ResponseWriter) {
	var response ApiResponse
	response.Code = 200
	response.Message = message
	response.Data = data

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return response, w
}

func GetCreatedResponse(w http.ResponseWriter, message string, data interface{}) (ApiResponse, http.ResponseWriter) {
	var response ApiResponse
	response.Code = 201
	response.Message = message
	response.Data = data

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	return response, w
}

func GetBadRequestResponse(w http.ResponseWriter, message string) (ApiResponse, http.ResponseWriter) {
	var response ApiResponse
	response.Code = 400
	response.Message = message

	w.WriteHeader(http.StatusBadRequest)

	return response, w
}

func GetInternalServerErrorResponse(w http.ResponseWriter, message string) (ApiResponse, http.ResponseWriter) {
	var response ApiResponse
	response.Code = 500
	response.Message = message

	w.WriteHeader(http.StatusInternalServerError)

	return response, w
}
