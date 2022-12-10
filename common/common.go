package common

import (
	"fmt"
)

type Forecast struct {
	Summary     WeatherSummary
	TempArray   []MinMaxPair
	AvgTemp     float64
	NextRainDay string
	WillItRain  []int
}

type WeatherSummary struct {
	MinTemp      int
	MaxTemp      int
	Humidity     int
	Wind         int
	ChanceOfRain int
}

type MinMaxPair struct {
	Min int
	Max int
}

func PrintForecast (forcast Forecast) {
	fmt.Println("Forecast start: --------------")
	fmt.Println("Summary: " )
	fmt.Println(forcast.Summary)
	fmt.Println("TempArray: " )
	fmt.Println(forcast.TempArray)
	fmt.Println("AvgTemp: " )
	fmt.Println(forcast.AvgTemp)
	fmt.Println("NextRainDay: " )
	fmt.Println(forcast.NextRainDay)
	fmt.Println("WillItRain: " )
	fmt.Println(forcast.WillItRain)
	fmt.Println("Forecast end: --------------")
}