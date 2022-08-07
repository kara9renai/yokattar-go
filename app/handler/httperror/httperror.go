package httperror

import (
	"log"
	"net/http"
)

func Error(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func BadRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func InternalServerError(w http.ResponseWriter, err error) {
	log.Printf("[Internal Server Error] %+v", err)

	Error(w, http.StatusInternalServerError)
}
