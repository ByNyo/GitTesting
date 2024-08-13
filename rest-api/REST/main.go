package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Weathers struct {
	Weathers []Weather
}

type Weather struct {
	timestamp                  string  `json:timestamp`
	source_id                  int     `json:source_id`
	precipitation              int     `json:precipitation`
	pressure_msl               float64 `json:pressure_msl`
	sunshine                   int     `json:sunshine`
	temperature                float64 `json:temperature`
	wind_direction             int     `json:wind_direction`
	wind_speed                 float64 `json:wind_speed`
	cloud_cover                int     `json:cloud_cover`
	dew_point                  float64 `json:dew_point`
	relative_humidity          int     `json:relative_humidity`
	visibility                 int     `json:visibility`
	wind_gust_direction        int     `json:wind_gust_direction`
	wind_gust_speed            float64 `json:wind_gust_speed`
	condition                  string  `json:condition`
	precipation_probability    any     `json:precipation_probability`
	precipation_probability_6h any     `json:precipation_probability_6h`
	solar                      int     `json:solar`
	icon                       string  `json:icon`
}

func saveJSONResponse(fileName string, key interface{}) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(key)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	url := "https://api.brightsky.dev/weather?lat=52&lon=7.6&date=2020-04-21"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		data := string(bodyBytes)
		fmt.Printf("Data from API: %+v\n", data)
		saveJSONResponse("weather.json", data)
		fmt.Println("Done!")
	}
}
