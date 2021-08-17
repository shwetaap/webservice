package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/shwetaap/webservice/server/handlers"
)

// initialize function starts the webserver
func initializeRouter() *mux.Router {
	l := log.New(os.Stdout, "server-log ", log.LstdFlags)
	//serverhandler = handlers.
	serverhandler := handlers.NewObjects(l)
	router := mux.NewRouter()
	// Read
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.GetObject).Methods("GET")
	// Update
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.UpdateObject).Methods("PUT")
	// Delete
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.DeleteObject).Methods("DELETE")

	return router
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

func main() {

	port := os.Getenv("BIND_PORT")
	if len(port) == 0 {

		log.Fatal("Program exits as environment variable BIND_PORT is not set. Please set the environment variable BIND_PORT to start the server")
	}

	router := initializeRouter()
	addr := ":" + port
	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		log.Printf("Serving on %s\n", port)
		log.Fatal(srv.ListenAndServe())
	}()
	// Graceful Shutdown
	waitForShutdown(srv)
}
