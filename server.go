package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	PORT = 8080
)

func main() {
	fmt.Printf("Starting server at port " + strconv.Itoa(PORT) + "\n")
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	}).Methods("GET")

	router.HandleFunc("/convert", handleConvertRequest).Methods("GET")

	err := http.ListenAndServe(":"+strconv.Itoa(PORT), router)
	if err != nil {
		log.Fatal(err)
	}
}
