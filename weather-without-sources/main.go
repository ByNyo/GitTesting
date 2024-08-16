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
)

type DayWeather struct {
	Weather []HourWeather `json:"weather"`
}

type HourWeather struct {
	Timestamp                string  `json:"timestamp"`
	SourceId                 int     `json:"source_id"`
	Precipitation            float64 `json:"precipitation"`
	PressureMsl              float64 `json:"pressure_msl"`
	Sunshine                 float64 `json:"sunshine"`
	Temperature              float64 `json:"temperature"`
	WindDirection            int     `json:"wind_direction"`
	WindSpeed                float64 `json:"wind_speed"`
	CloudCover               int     `json:"cloud_cover"`
	DewPoint                 float64 `json:"dew_point"`
	RelativeHumidity         int     `json:"relative_humidity"`
	Visibility               int     `json:"visibility"`
	WindGustDirection        int     `json:"wind_gust_direction"`
	WindGustSpeed            float64 `json:"wind_gust_speed"`
	Condition                string  `json:"condition"`
	PrecipationProbability   float64 `json:"precipation_probability"`
	PrecipationProbability6h float64 `json:"precipation_probability_6h"`
	Solar                    float64 `json:"solar"`
	Icon                     string  `json:"icon"`
}

type OWCity struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
}

type City struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

var apiKey = "6fb9b01da8337245c3fbd5ac0dac4d62"
var URL = "https://api.brightsky.dev/weather?lat=52&lon=7.6&date=2020-04-21"
var _url = "https://api.brightsky.dev/weather?"
var date = "2020-04-21"
var latitude float64 = 52.5
var longitude float64 = 13.4

// Date Format year-month-day
func SetDate(year, month, day int) {
	if month <= 12 && month >= 1 && year >= 2010 && year <= time.Now().Year() && day >= 0 && day <= 31 {
		date = fmt.Sprintf("%d-02%d-02%d", year, month, day)
		reloadURL()
	} else {
		layout := "2006-01-02"
		date = time.Now().Format(layout)
		// Make a message in app for user: "Incorrect Date" (something like that)
	}
}

func SetLocationByCityName(name string, cities map[string]City) {
	cities = ReadCities(cities)
	if city, exists := cities[strings.ToLower(name)]; exists { // When the city exists
		fmt.Println("City:", city)
	} else { // When the city doesn't exist
		SaveCityByName(name, cities)
	}
}

func ReadCities(cities map[string]City) map[string]City {
	file, err := os.Open("resources/cities.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cities)
	return cities
}

func SaveCityByName(name string, cities map[string]City) {
	var owcities []OWCity
	resp, err := http.Get("http://api.openweathermap.org/geo/1.0/direct?q=" + name + "&limit=1&appid=" + apiKey)
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
	cities[strings.ToLower(foundcity.Name)] = City{Lat: foundcity.Lat, Lon: foundcity.Lon}
	data, err := json.MarshalIndent(cities, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = saveJSONFile("resources", "cities.json", data)
	if err != nil {
		log.Fatal(err)
	}
}

func SetLocation(Latitude float64, Longitude float64) {
	if Latitude <= 54 && Latitude >= 48 && Longitude <= 14 && Longitude >= 6 {
		latitude = Latitude
		longitude = Longitude
		reloadURL()
	} else {
		// Make a message in app for user: "Not in range" (something like that)
	}
}

func SetDateAndLocation(year, month, day int, Latitude float64, Longitude float64) {
	SetDate(year, month, day)
	SetLocation(Latitude, Longitude)
}

func reloadURL() {
	u, err := url.Parse(URL)
	if err != nil {
		log.Fatal(err)
	}
	values := u.Query()
	values.Set("lat", strconv.FormatFloat(latitude, 'f', 2, 64))
	values.Set("lon", strconv.FormatFloat(longitude, 'f', 2, 64))
	values.Set("lat", strconv.FormatFloat(latitude, 'f', 2, 64))
	values.Set("date", date)
	u.RawQuery = values.Encode()
	URL = u.Redacted()
}

func ShowWeatherFromTime(day DayWeather, t time.Time) {
	fmt.Println(day.Weather[t.Hour()])
	// Make use of it later
}

var today DayWeather

func main() {
	Cities := make(map[string]City)
	SetLocationByCityName("Berlin", Cities)
	SetDateAndLocation(2024, 8, 15, Cities["Berlin"].Lat, Cities["Berlin"].Lon)
	RequestWeather()
	ShowWeatherFromTime(today, time.Now())
	SaveWeather() // Saves weather in weather.json
	doEvery(30*time.Minute, ShowWeather)
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

func ShowWeather(time.Time) {
	ShowWeatherFromTime(today, time.Now())
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		RequestWeather()
		f(x)
	}
}

func RequestWeather() {
	response, err := http.Get(URL)
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

// jq commands:
// 2: cat weather.json | jq
// 3.1: cat weather.json | jq '.weather[1]'
// 3.2: cat weather.json | jq '.weather[] | select(.wind_speed>13.3)'
// 3.3: cat weather.json | jq '.weather[] | [.temperature, .wind_speed, .relative_humidity]'
// 3.4: cat weather.json | jq '[.weather[].temperature] | add / length'
// 3.5 Sorting: cat weather.json | jq '.weather[].temperature' | sort -n
// 3.5 Finding: cat weather.json | jq '[.weather[].temperature] | max'
