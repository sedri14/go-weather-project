package main

import (
	"fmt"

	"example.com/common"
	"example.com/weatherData"
)

func main() {
    //weatherData.HelloFromWeatherData()
	var f common.Forecast
	f = weatherData.GetForecast(4)
	fmt.Println("Main>>> ")
	fmt.Println(f)
}