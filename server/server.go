package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/shwetaap/webservice/server/handlers"
)

// initialize function starts the webserver
func initialize(port string) {
	l := log.New(os.Stdout, "server-log", log.LstdFlags)
	//serverhandler = handlers.
	serverhandler := handlers.NewObjects(l)
	router := mux.NewRouter()
	// Read
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.GetObject).Methods("GET")
	// Update
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.UpdateObject).Methods("PUT")
	// Delete
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.DeleteObject).Methods("DELETE")

	log.Printf("Serving on %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {

	port := os.Getenv("BIND_PORT")
	if len(port) == 0 {

		log.Fatal("Program exits as environment variable BIND_PORT is not set. Please set the environment variable BIND_PORT to start the server")
	}

	initialize(port)
}
