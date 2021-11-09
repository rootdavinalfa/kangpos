/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package common

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Wrap the router for GET method
func Get(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func Post(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func Put(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func Delete(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, f).Methods("DELETE")
}
