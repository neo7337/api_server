package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"api_server/config"
	"api_server/handler"
)

func main() {

	//loading global configs
	var configs config.Conf
	configs.ReadConf()

	//server routes
	http.HandleFunc("/", handler.ServerStatus())
	http.HandleFunc("/api/v1/listCountries", handler.CountriesList(configs.CountriesHost, configs.CountriesPath))

	server := &http.Server{
		Addr:         ":8989",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server is up at ", server.Addr, time.Now().Format(time.Stamp))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
