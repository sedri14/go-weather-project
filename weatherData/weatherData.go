package weatherData

import (
	"fmt"

	"example.com/common"
	"example.com/source1"
	"example.com/source2"
	"example.com/source3"
)

// type Forecast struct {
// 	Summary WeatherSummary
// 	TempArray []MinMaxPair
// 	AvgTemp float64
// 	NextRainDay string
// 	WillItRain []int
// }

func HelloFromWeatherData() {
	// source1.HelloFromSource1()
	//source2.HelloFromSource2()
	//source3.HelloFromSource3()
    //message := fmt.Sprintf(messageSource1)
    //return message
	//source1.getPageContent
}

func GetForecast(days int) common.Forecast {

	forecastSource1 := source1.GetForecast(days)
	forecastSource2 := source2.GetForecast(days)
	forecastSource3 := source3.GetForecast(days)

	fmt.Println("source1:")
	common.PrintForecast(forecastSource1)
	fmt.Println("source2:")
	common.PrintForecast(forecastSource2)
	fmt.Println("source3:")
	common.PrintForecast(forecastSource3)


	return forecastSource1

}