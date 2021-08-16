package data

import (
	"encoding/json"
	"fmt"
	"io"
)

// NumberOfBuckets defines the number of buckets used
const NumberOfBuckets = 10

// Object holds the data
type Object struct {
	ID   int    `json:"id"`
	Data string `json:"data"`
}

// Bucket holds multiple objects
var Bucket []*Object

var ObjectMap = make(map[Object]int)

// Buckets is a slice of multiple buckets
var Buckets [][]*Object

// ErrNotFound is the error message returned when the object is not found
var ErrNotFound = fmt.Errorf("Object not found")

// ToJSON converts the object into json
func (o *Object) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(o)
}

// FromJSON converts the json input into the data object format
func (o *Object) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(o)
}

// findobject is the helper function that searches for the given object
func findObject(objectID int, bucketID int) (*Object, error) {
	if len(Buckets) > 0 {
		if len(Buckets[bucketID]) > 0 {
			for _, object := range Buckets[bucketID] {
				if object.ID == objectID {
					return object, nil
				}
			}
		}
	}
	return nil, ErrNotFound
}

// GetObject looks for the given object in the request and returns the data
func GetObject(objectID int, bucketID int) (*Object, error) {

	object, err := findObject(objectID, bucketID)
	if err != nil {
		return nil, err
	}
	if object != nil {
		return object, nil
	}
	return nil, nil
}

// UpdateObject looks for the given object, if found updates its data and if not found creates the new object
func UpdateObject(objectID int, bucketID int, obj *Object) error {

	object, _ := findObject(objectID, bucketID)

	if object != nil {
		object.Data = obj.Data
	} else if len(Buckets) == 0 {
		Buckets = make([][]*Object, NumberOfBuckets)
		Buckets[bucketID] = append(Buckets[bucketID], obj)
	} else {
		Buckets[bucketID] = append(Buckets[bucketID], obj)
	}
	return nil
}

// DeleteObject deleted the object in the request
func DeleteObject(objectID int, bucketID int) error {
	if len(Buckets) > 0 {
		if len(Buckets[bucketID]) > 0 {
			for index, object := range Buckets[bucketID] {
				if object.ID == objectID {
					Buckets[bucketID] = append(Buckets[bucketID][:index], Buckets[bucketID][index+1:]...)
					return nil
				}
			}
		}
	}
	return ErrNotFound
}
