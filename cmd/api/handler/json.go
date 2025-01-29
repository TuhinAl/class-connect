package handler

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func ReadJSONRequest(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024 * 3 // 3 MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

func WriteJSONError(w http.ResponseWriter, statusCode int, err error) error {
return WriteJSONResponse(w, statusCode, envelop{
	Success: false,
	Message: err.Error(),
})
}

type envelop struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ApiResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`	
}

type Pagination struct {
    Page  int    `json:"page"`
    Limit int    `json:"limit"`
    Sort  string `json:"sort"`
}