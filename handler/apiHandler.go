package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//ServerStatus -> Function that return the live status
func ServerStatus() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ping Server Test")
		w.WriteHeader(http.StatusOK)
	}
}

//RestHandler -> Generic handler for the rest calls in the server
func RestHandler(host string, path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Fetching Data")
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
