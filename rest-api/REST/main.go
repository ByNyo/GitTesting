package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type DayWeather struct {
	Weather []HourWeather `json:"weather"`
	Source  []Source      `json:"sources"`
}

type HourWeather struct {
	Timestamp                  string  `json:"timestamp"`
	Source_id                  int     `json:"source_id"`
	Precipitation              float64 `json:"precipitation"`
	Pressure_msl               float64 `json:"pressure_msl"`
	Sunshine                   float64 `json:"sunshine"`
	Temperature                float64 `json:"temperature"`
	Wind_direction             int     `json:"wind_direction"`
	Wind_speed                 float64 `json:"wind_speed"`
	Cloud_cover                int     `json:"cloud_cover"`
	Dew_point                  float64 `json:"dew_point"`
	Relative_humidity          int     `json:"relative_humidity"`
	Visibility                 int     `json:"visibility"`
	Wind_gust_direction        int     `json:"wind_gust_direction"`
	Wind_gust_speed            float64 `json:"wind_gust_speed"`
	Condition                  string  `json:"condition"`
	Precipation_probability    float64 `json:"precipation_probability"`
	Precipation_probability_6h float64 `json:"precipation_probability_6h"`
	Solar                      float64 `json:"solar"`
	Icon                       string  `json:"icon"`
}

type Source struct {
	Id               int     `json:"id"`
	Dwd_station_id   string  `json:"dwd_station_id"`
	Observation_type string  `json:"observation_type"`
	Latitude         float64 `json:"lat"`
	Longitude        float64 `json:"lon"`
	Height           float64 `json:"height"`
	Station_name     string  `json:"station_name"`
	Wmo_station_id   string  `json:"wmo_station_id"`
	First_record     string  `json:"first_record"`
	Last_record      string  `json:"last_record"`
	Distance         float64 `json:"distance"`
}

func main() {
	url := "https://api.brightsky.dev/weather?lat=52&lon=7.6&date=2020-04-21"
	var today DayWeather
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &today)
	if err != nil {
		log.Fatal(err)
	}

	mdata, err := json.MarshalIndent(today, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = saveJSONFile("weather.json", mdata)
	if err != nil {
		log.Fatal(err)
	}
}

func saveJSONFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Println("File created!")
	return nil
}
