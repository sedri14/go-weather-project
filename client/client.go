package main

import (
	"fmt"

	"example.com/common"
	"example.com/weatherData"
)

// type Forecast struct {
// 	Summary WeatherSummary
// 	TempArray []MinMaxPair
// 	AvgTemp float64
// 	NextRainDay string
// 	WillItRain []int
// }

// type WeatherSummary struct {
// 	minTemp int
// 	maxTemp int
// 	humidity int
// 	wind int
// 	chanceOfRain int
// }

// type MinMaxPair struct {
// 	min int
// 	max int
// }

func main() {
    //weatherData.HelloFromWeatherData()
	var f common.Forecast
	f = weatherData.GetForecast(4)
	fmt.Println("Main>>>The result is: ")
	fmt.Println(f)
}