package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"devstreamline/handlers"

	"github.com/gorilla/mux"
)

const (
	PORT = 8080
)

func main() {
	fmt.Printf("Starting server at port " + strconv.Itoa(PORT) + "\n")
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	}).Methods("GET")

	router.HandleFunc("/convert", handlers.HandleConvertRequest).Methods("GET")

	err := http.ListenAndServe(":"+strconv.Itoa(PORT), router)
	if err != nil {
		log.Fatal(err)
	}
}
