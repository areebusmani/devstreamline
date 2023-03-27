package main

import (
	"net/http"
)

func handleConvertRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	return
}
