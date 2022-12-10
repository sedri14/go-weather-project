package weatherData

import (
	//"fmt"

	"example.com/common"
	"example.com/source1"
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
	// forecastSource2 := source1.GetForecast(days)
	// forecastSource3 := source1.GetForecast(days)

	return forecastSource1

}