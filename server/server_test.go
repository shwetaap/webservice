package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/shwetaap/webservice/server/data"
	"github.com/shwetaap/webservice/server/handlers"
)

var testobj1 = data.Object{
	ID:   1,
	Data: "first",
}

var testbucketindex = 0

func executeRequest(req *http.Request, router *mux.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetObjectEmpty(t *testing.T) {

	l := log.New(os.Stdout, "server-test-get-empty-object", log.LstdFlags)
	serverhandler := handlers.NewObjects(l)
	req, _ := http.NewRequest("GET", "/objects/0/1", nil)
	router := mux.NewRouter()
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.GetObject).Methods("GET")
	response := executeRequest(req, router)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestGetObject(t *testing.T) {

	_ = data.UpdateObject(testobj1.ID, testbucketindex, &testobj1)

	l := log.New(os.Stdout, "server-test-get-object", log.LstdFlags)
	serverhandler := handlers.NewObjects(l)
	req, _ := http.NewRequest("GET", "/objects/0/1", nil)
	router := mux.NewRouter()
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.GetObject).Methods("GET")
	response := executeRequest(req, router)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateObject(t *testing.T) {

	l := log.New(os.Stdout, "server-test-update-object", log.LstdFlags)

	serverhandler := handlers.NewObjects(l)

	j, _ := json.Marshal(testobj1)
	req, _ := http.NewRequest("PUT", "/objects/0/1", bytes.NewBuffer(j))
	router := mux.NewRouter()
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.UpdateObject).Methods("PUT")
	response := executeRequest(req, router)
	checkResponseCode(t, http.StatusOK, response.Code)

	// Get the newly added data
	req, _ = http.NewRequest("GET", "/objects/0/1", nil)
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.GetObject).Methods("GET")
	response = executeRequest(req, router)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestDeleteExistingObject(t *testing.T) {

	l := log.New(os.Stdout, "server-test-delete-existing-object", log.LstdFlags)

	serverhandler := handlers.NewObjects(l)

	j, _ := json.Marshal(testobj1)
	req, _ := http.NewRequest("PUT", "/objects/0/1", bytes.NewBuffer(j))
	router := mux.NewRouter()
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.UpdateObject).Methods("PUT")
	response := executeRequest(req, router)
	checkResponseCode(t, http.StatusOK, response.Code)

	// Get the newly added data
	req, _ = http.NewRequest("GET", "/objects/0/1", nil)
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.GetObject).Methods("GET")
	response = executeRequest(req, router)
	checkResponseCode(t, http.StatusOK, response.Code)

	// Delete the data
	req, _ = http.NewRequest("DELETE", "/objects/0/1", nil)
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.DeleteObject).Methods("DELETE")
	response = executeRequest(req, router)
	checkResponseCode(t, http.StatusOK, response.Code)

	// Check if the data is deleted
	req, _ = http.NewRequest("GET", "/objects/0/1", nil)
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.GetObject).Methods("GET")
	response = executeRequest(req, router)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestDeleteNonExistingObject(t *testing.T) {

	l := log.New(os.Stdout, "server-test-delete-nonexisting-object", log.LstdFlags)

	serverhandler := handlers.NewObjects(l)

	// Delete the data
	router := mux.NewRouter()
	req, _ := http.NewRequest("DELETE", "/objects/0/1", nil)
	router.HandleFunc("/objects/{bucket:[0-9]}/{objectID:[0-9]+}", serverhandler.DeleteObject).Methods("DELETE")
	response := executeRequest(req, router)
	checkResponseCode(t, http.StatusNotFound, response.Code)

}
