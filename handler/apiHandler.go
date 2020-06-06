package handler

import (
	"fmt"
	"io/ioutil"
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
		fmt.Println("Fetching List of Countries")
		uri := "https://" + host + path
		response, err := http.Get(uri)
		if err != nil {
			fmt.Printf("Failed HTTP Request %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			//fmt.Println(string(data))
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	}
}
