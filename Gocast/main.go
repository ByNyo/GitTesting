package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var weatherAPIURL string
var geoAPIURL string
var geoAPIKey string
var latitude = 52.5170365
var longitude = 13.3888599
var date string
var cityName string
var today DayWeather

type DayWeather struct {
	Hours []HourWeather `json:"weather"`
}

type HourWeather struct {
	TimeStamp                  string  `json:"timestamp"`
	SourceID                   int     `json:"source_id"`
	Precipitation              float64 `json:"precipitation"`
	PressureMSL                float64 `json:"pressure_msl"`
	Sunshine                   float64 `json:"sunshine"`
	Temperature                float64 `json:"temperature"`
	WindDirection              int     `json:"wind_direction"`
	WindSpeed                  float64 `json:"wind_speed"`
	CloudCover                 int     `json:"cloud_cover"`
	DewPoint                   float64 `json:"dew_point"`
	RelativeHumidity           int     `json:"relative_humidity"`
	Visibility                 int     `json:"visibility"`
	WindGustDirection          int     `json:"wind_gust_direction"`
	WindGustSpeed              float64 `json:"wind_gust_speed"`
	Condition                  string  `json:"condition"`
	PrecipitationProbability   float64 `json:"precipitation_probability"`
	PrecipitationProbability6h float64 `json:"precipitation_probability_6h"`
	Solar                      float64 `json:"solar"`
	Icon                       string  `json:"icon"`
}

type OWCity struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Country   string  `json:"country"`
}

type City struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func SetDateAndLocationByCityName(year, month, day int, cityName string, cities map[string]City) {
	SetDate(year, month, day)
	SetLocationByCityName(cityName, cities)
}

func SetDateAndLocation(year, month, day int, lat, lon float64) {
	SetDate(year, month, day)
	SetLocation(lat, lon)
}

func SetDate(year, month, day int) {
	_date := fmt.Sprintf("%v-%.2v-%.2v", year, month, day)
	_, err := checkDate(_date)
	if err != nil { // What happens when date is invalid
		fmt.Println("Error: ", err)
		// TODO: Throw error at user
	}
	if year >= 2010 && year <= time.Now().Year() && month >= 1 && month <= 12 && day >= 1 && day <= 31 {
		date = _date
	} else { // What happens when the date is out of range
		date = fmt.Sprintf("%v-%.2v-%.2v", time.Now().Year(), int(time.Now().Month()), time.Now().Day())
		// TODO: Throw error at user
	}
	reloadWeatherURL()
}

func SetLocation(lat, lon float64) {
	if lat <= 54 && lat >= 48 && lon <= 14 && lon >= 6 {
		latitude = lat
		longitude = lon
		reloadWeatherURL()
	} else { // When location is not in range
		// Throw error at user
	}
}

func SetLocationByCityName(name string, cities map[string]City) {
	cities = ReadCities(name, cities)
	if city, exists := cities[strings.ToLower(name)]; exists { // When the city exists
		SetLocation(city.Lat, city.Lon)
	} else { // When the city doesn't exist
		SaveCityByName(name, cities)
		SetLocationByCityName(name, cities)
		// Maybe throw error
	}
}

func ReadCities(name string, cities map[string]City) map[string]City {
	file, err := os.Open("resources/cities.json")
	if err != nil {
		SaveCityByName(name, cities)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cities)
	if err != nil {
		fmt.Println(err)
	}
	return cities
}

func SaveCityByName(name string, cities map[string]City) string {
	cityName = name
	reloadGeoURL()
	var owcities []OWCity
	resp, err := http.Get(geoAPIURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &owcities)
	if err != nil {
		log.Fatal(err)
	}
	var foundcity OWCity
	for _, owcity := range owcities {
		if owcity.Country == "DE" {
			foundcity = owcity
		}
	}
	cities[strings.ToLower(name)] = City{Lat: foundcity.Latitude, Lon: foundcity.Longitude}
	data, err := json.MarshalIndent(cities, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	saveJSONFile("resources", "cities.json", data)
	return foundcity.Name
}

func reloadWeatherURL() {
	u, err := url.Parse(weatherAPIURL)
	if err != nil {
		log.Fatal(err)
	}

	v := u.Query()
	v.Set("lat", strconv.FormatFloat(latitude, 'f', 2, 64))
	v.Set("lon", strconv.FormatFloat(longitude, 'f', 2, 64))
	v.Set("date", date)
	u.RawQuery = v.Encode()
	weatherAPIURL = u.String()
}

func reloadGeoURL() {
	u, err := url.Parse(geoAPIURL)
	if err != nil {
		log.Fatal(err)
	}

	v := u.Query()
	v.Set("q", cityName)
	v.Set("appid", geoAPIKey)
	u.RawQuery = v.Encode()
	geoAPIURL = u.Redacted()
}

func main() {
	Cities := make(map[string]City)
	godotenv.Load()
	godotenv.Load(".env.dev")
	weatherAPIURL = os.Getenv("WEATHERAPI_URL")
	geoAPIURL = os.Getenv("GEOAPI_URL")
	geoAPIKey = os.Getenv("GEOAPI_KEY")
	SetLocationByCityName("Berlin", Cities)
	SetDate(2010, 2, 28)
	reloadWeatherURL()
	RequestWeather()
	SaveWeather()
}

func checkDate(s string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return date, err
	}
	return date, nil
}

func saveJSONFile(directory, filename string, data []byte) error {
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	filePath := fmt.Sprintf("%s/%s", directory, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func RequestWeather() {
	resp, err := http.Get(weatherAPIURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &today)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveWeather() {
	data, err := json.MarshalIndent(today, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = saveJSONFile("resources", "weather.json", data)
	if err != nil {
		log.Fatal(err)
	}
}
