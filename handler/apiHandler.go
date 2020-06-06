package handler

import (
	"fmt"
	"net/http"
)

var count = 0

//ServerStatus -> Function that return the live status
func ServerStatus() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		count++
		fmt.Println("Server request for endpoint /", count)
		w.WriteHeader(http.StatusOK)
	}
}

//CountriesList -> Function returns the list of ISO standard countries
func CountriesList(host string, path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(host + " " + path)
	}
}
