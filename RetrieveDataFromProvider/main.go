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
	"time"
)

type DayWeather struct {
	Weather []HourWeather `json:"weather"`
	Source  []Source      `json:"sources"`
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

type Source struct {
	ID              int     `json:"id"`
	DwdStationId    string  `json:"dwd_station_id"`
	ObservationType string  `json:"observation_type"`
	Latitude        float64 `json:"lat"`
	Longitude       float64 `json:"lon"`
	Height          float64 `json:"height"`
	StationName     string  `json:"station_name"`
	WmoStationId    string  `json:"wmo_station_id"`
	FirstRecord     string  `json:"first_record"`
	LastRecord      string  `json:"last_record"`
	Distance        float64 `json:"distance"`
}

type Cities struct {
	city []City
}

type City struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
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

func SetLocationByCityName(name string) {
	var city Cities
	resp, err := http.Get("http://api.openweathermap.org/geo/1.0/direct?q=" + name + "&limit=5&appid=" + apiKey)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &city.city)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("City:", city.city[0].Name)
	fmt.Println("Latitude:", city.city[0].Lat)
	fmt.Println("Longitude:", city.city[0].Lon)
	SetLocation(city.city[0].Lat, city.city[0].Lon)
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
	values.Set("date", date)
	u.RawQuery = values.Encode()
	URL = u.Redacted()
}

func ShowWeatherFromTime(day DayWeather, t time.Time) {
	fmt.Println(day.Weather[t.Hour()])
}

func main() {
	SetLocationByCityName("München")
	SetDateAndLocation(2024, 8, 15, 52.5, 13.4)

	var today DayWeather
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
	/*data, err := json.MarshalIndent(today, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = saveJSONFile("weather.json", data)
	if err != nil {
		log.Fatal(err)
	}*/
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

// jq commands:
// 2: cat weather.json | jq
// 3.1: cat weather.json | jq '.weather[1]'
// 3.2: cat weather.json | jq '.weather[] | select(.wind_speed>13.3)'
// 3.3: cat weather.json | jq '.weather[] | [.temperature, .wind_speed, .relative_humidity]'
// 3.4: cat weather.json | jq '[.weather[].temperature] | add / length'
// 3.5 Sorting: cat weather.json | jq '.weather[].temperature' | sort -n
// 3.5 Finding: cat weather.json | jq '[.weather[].temperature] | max'
