package handler

import (
	"log"
	"net/http"
)

func (app *ApplicationConfig) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("Internal Server Error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusInternalServerError, err)
}

func (app *ApplicationConfig) BadRequestError(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("Bad Request Error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusBadRequest, err)
}


func (app *ApplicationConfig) NotfoundError(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("Not Found Error: %s path: %s error: %s", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusNotFound, err)
}
