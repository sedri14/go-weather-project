package common

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