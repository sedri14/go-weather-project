package weatherData

import (
	"fmt"
	"math"

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



func summaryAverage(forecasts []common.Forecast) common.WeatherSummary {
	var sumMinTemp float64
	var sumMaxTemp float64
	var sumHumidity float64
	var sumWind float64
	var chanceOfRain float64
	
	var summary common.WeatherSummary
	numForcasts := float64(len(forecasts))
	for _,f := range forecasts{
		summary = f.Summary
		sumMinTemp += float64(summary.MinTemp)
		sumMaxTemp += float64(summary.MaxTemp)
		sumHumidity += float64(summary.Humidity)
		sumWind += float64(summary.Wind)
		chanceOfRain += float64(summary.ChanceOfRain)
	}

	avgSummary := common.WeatherSummary {
		MinTemp: int(math.Round(sumMinTemp / numForcasts)),
		MaxTemp: int(math.Round(sumMaxTemp / numForcasts)),
		Humidity: int(math.Round(sumHumidity / numForcasts)),
		Wind: int(math.Round(sumWind / numForcasts)),
		ChanceOfRain: int(math.Round(chanceOfRain / numForcasts)),
	}
	
	return avgSummary
}

func temperatureAverage(forecasts []common.Forecast) float64 {
	var sumTemp float64
	
	numForcasts := float64(len(forecasts))
	for _,f := range forecasts{
		sumTemp += f.AvgTemp
	}

	return sumTemp / numForcasts
}

func willItRainAverage(forecasts []common.Forecast) []int {
	
	numForcasts := float64(len(forecasts))
	numDays := len(forecasts[0].WillItRain)
	var avgSlice = make([]int,numDays)

	for _,f := range forecasts{
		for i,chance := range f.WillItRain {
			avgSlice[i] += chance
		}
	}

	for i,sum := range avgSlice {
		avgSlice[i] = int(math.Round(float64(sum) / numForcasts))
	}

	return avgSlice
}

func tempArrayAverage(forecasts []common.Forecast) []common.MinMaxPair {
	
	numForcasts := float64(len(forecasts))
	numDays := len(forecasts[0].TempArray)
	var avgSlice = make([]common.MinMaxPair,numDays)

	for _,f := range forecasts{
		for i,pair := range f.TempArray {
			avgSlice[i].Min += pair.Min
			avgSlice[i].Max += pair.Max
		}
	}

	for i,pair := range avgSlice {
		avgSlice[i].Min = int(math.Round(float64(pair.Min) / numForcasts))
		avgSlice[i].Max = int(math.Round(float64(pair.Max) / numForcasts))
	}

	return avgSlice
}

func printChanceOfRain (forecasts []common.Forecast, days int) {
	fmt.Printf("Chance of rain in the next %d days: ", days)
	for _,chance := range willItRainAverage(forecasts) {
		fmt.Printf("%d%%, ", chance)
	}
	fmt.Println()
}

func printMinMaxTemp (forecasts []common.Forecast, days int) {
	fmt.Printf("Max and min temperatures over the next %d days: ", days)
	for _,pair := range tempArrayAverage(forecasts) {
		fmt.Printf("(Max: %d, Min: %d), ", pair.Max, pair.Min)
	}
	fmt.Println()
}

func printNextRainDay(forecasts []common.Forecast) {
	fmt.Println("Next day likely to rain: ")
	for i,forecast := range forecasts {
		nextRainDay := forecast.NextRainDay
		if (nextRainDay == "") {
			nextRainDay = "None"
		}
		fmt.Printf("Source #%d: %s\n", i+1, nextRainDay)
	}
}

func GetForecast(days int) {

	var forecasts []common.Forecast
	forecasts = append(forecasts, source1.GetForecast(days))
	forecasts = append(forecasts, source2.GetForecast(days))
	forecasts = append(forecasts, source3.GetForecast(days))

	common.PrintSummary(summaryAverage(forecasts))
	fmt.Println("--- Additional Info ---")
	fmt.Printf("Average temperature today: %.1f˚C\n", temperatureAverage(forecasts))
	printChanceOfRain(forecasts, days)
	printMinMaxTemp(forecasts, days)
	printNextRainDay(forecasts)
}