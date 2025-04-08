package render

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data"`
}

func RenderJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := Response{
		Data: data,
	}

	json.NewEncoder(w).Encode(response)
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  *int   `json:"code,omitempty"`
}

type ErrorOptions struct {
	Code *int `json:"code,omitempty"`
}

func RenderError(w http.ResponseWriter, status int, err error, opts *ErrorOptions) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response := ErrorResponse{
		Error: err.Error(),
		Code:  opts.Code,
	}

	json.NewEncoder(w).Encode(response)
}
