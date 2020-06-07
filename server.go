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
	http.HandleFunc("/api/v1/listCountries", handler.RestHandler(configs.CountriesHost, configs.CountriesPath))
	http.HandleFunc("/api/v1/indStats", handler.RestHandler(configs.LocaleHost, configs.LocalePath))
	http.HandleFunc("/api/v1/indDistrict", handler.RestHandler(configs.IndHost, configs.DistrictIndPath))
	http.HandleFunc("/api/v1/dataJson", handler.RestHandler(configs.IndHost, configs.HistoricalINDPath))

	server := &http.Server{
		Addr:         ":8989",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server is Up and Running ", time.Now().Format(time.Stamp))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
