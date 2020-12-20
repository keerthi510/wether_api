package main

import (
	"WeatherAPI/internal/config"
	"WeatherAPI/internal/handler"
	"WeatherAPI/internal/sqllite"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//Loading the env configs
	config, _ := config.Load()

	// Initializing the db to store data to avoid constant  API usage
	err, dbfunc := sqllite.Initdatabase()
	if err != nil {
		fmt.Println(err)
	}
	err = dbfunc.CreateSchema()
	if err != nil {
		fmt.Println(err)
	}
	// Go routine to constantly delete data from  sqllite in casee it has existed in db for more than 2 mins
	go func() {
		for {
			err := dbfunc.DeleteData()
			if err != nil {
				fmt.Println("Errror occurred in deleting")
			}
			time.Sleep(60 * (time.Second))
		}
	}()
	http.HandleFunc("/weather/", handler.WeatherInfo)
	fmt.Println(config)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))

}
