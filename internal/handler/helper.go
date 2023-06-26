package handler

import (
	"log"
	"net/http"
)

func ErrorHandlerReturn(err error, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)

	return
}
