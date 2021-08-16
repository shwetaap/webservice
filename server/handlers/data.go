package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/shwetaap/webservice/server/data"

	"github.com/gorilla/mux"
)

// Objects is the Handler
type Objects struct {
	l *log.Logger
}

// NewObjects functions creates a Object handler with a logger
func NewObjects(l *log.Logger) *Objects {
	return &Objects{l}
}

// GetObject returns the data for the HTTP GET Request
func (o *Objects) GetObject(w http.ResponseWriter, r *http.Request) {
	o.l.Println("Handle Get Request")

	vars := mux.Vars(r)

	bucket, err := strconv.Atoi(vars["bucket"])
	if err != nil {
		http.Error(w, "Unable to convert bucket", http.StatusBadRequest)
		return
	}

	objID, err := strconv.Atoi(vars["objectID"])
	if err != nil {
		http.Error(w, "Unable to convert objectID", http.StatusBadRequest)
		return
	}
	obj, err := data.GetObject(objID, bucket)
	if err == data.ErrNotFound {
		http.Error(w, "Object not found", http.StatusNotFound)
		return
	}

	err = obj.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

// UpdateObject creates or updates the object referenced in the PUT Request
func (o *Objects) UpdateObject(w http.ResponseWriter, r *http.Request) {
	o.l.Println("Handle Put Request")
	vars := mux.Vars(r)

	bucket, err := strconv.Atoi(vars["bucket"])
	if err != nil {
		http.Error(w, "Unable to convert bucket", http.StatusBadRequest)
		return
	}

	objID, err := strconv.Atoi(vars["objectID"])
	if err != nil {
		http.Error(w, "Unable to convert objectID", http.StatusBadRequest)
		return
	}

	obj := &data.Object{}

	err = obj.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateObject(objID, bucket, obj)
	if err == data.ErrNotFound {
		http.Error(w, "Object not found", http.StatusNotFound)
		return
	}

}

// DeleteObject deletes the object listed in the HTTP DELETE Request
func (o *Objects) DeleteObject(w http.ResponseWriter, r *http.Request) {
	o.l.Println("Handle Delete Request")

	vars := mux.Vars(r)

	bucket, err := strconv.Atoi(vars["bucket"])
	if err != nil {
		http.Error(w, "Unable to convert bucket", http.StatusBadRequest)
		return
	}

	objID, err := strconv.Atoi(vars["objectID"])
	if err != nil {
		http.Error(w, "Unable to convert objectID", http.StatusBadRequest)
		return
	}

	err = data.DeleteObject(objID, bucket)
	if err == data.ErrNotFound {
		http.Error(w, "Object not found", http.StatusNotFound)
		return
	}

}
