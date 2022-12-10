package source3

//--------------www.meteoprog.com--------------

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

const URL = "https://www.meteoprog.com/review/Telaviv/"

type Data struct {
		Date string
		Max int
		Min int
		Date_short string
		Humidity int
		Wind_speed int
		Pop int
}

type WeatherSummary struct {
	minTemp int
	maxTemp int
	humidity int
	wind int
	chanceOfRain int
}

type MinMaxPair struct {
	min int
	max int
}

func getHtml(url string) *http.Response {
	response, error := http.Get(url)
	if error != nil {
		fmt.Println(error)
	}

	if response.StatusCode != 200 {
		fmt.Println("Status code:", response.StatusCode)
	}

	return response
}

func getPageBodyJson() []Data {
	r := regexp.MustCompile(`var data = (.*);`)
	response := getHtml(URL)
	body, err := io.ReadAll(response.Body)
	check(err)
	bodyStr := r.FindStringSubmatch(string(body))
	var jsonData []Data
	err = json.Unmarshal([]byte(bodyStr[1]), &jsonData)
	check(err)

	return jsonData
}

// func getPageContent() *goquery.Document{
// 	response := getHtml(URL)
// 	defer response.Body.Close()

// 	doc, error := goquery.NewDocumentFromReader(response.Body)
// 	check(error)

// 	return doc
// }

func check(error error) {
	if (error != nil) {
		fmt.Println(error)
	}
}

func getWeatherSummary(data []Data) WeatherSummary {
	humidity := data[0].Humidity
	min, max := data[0].Min, data[0].Max
	chance := data[0].Pop
	wind := data[0].Wind_speed

	return WeatherSummary {min, max, humidity, wind, chance}
}

func TempArray(days int, data []Data) []MinMaxPair {
	var minMaxPairs []MinMaxPair

	for i:=0; i < days; i++ {
		minMaxPairs = append(minMaxPairs, MinMaxPair{data[i].Min, data[i].Max})
	}

	return minMaxPairs
}

func AverageTemp(days int, data[]Data) float64 {
	tempArr := TempArray(days, data)

	var maxSum int
	var minSum int
	for _,pair := range tempArr {
		maxSum = maxSum + pair.max
		minSum = minSum + pair.min 
	}

	maxAvg := float64(maxSum) / float64(len(tempArr))
	minAvg := float64(minSum)/ float64(len(tempArr))
	avg := (minAvg + maxAvg) / 2.0

	return avg
}

func NextDayRain(data []Data) (string, bool) {

	for _,item := range data {
		if item.Pop > 50 {
			return item.Date, true
		}
	}

	return "", false
}

func WillItRain(days int, data []Data) []int{
	var chances []int

	for i:=0; i < days; i++ {
		chances = append(chances, data[i].Pop)
	}

	return chances
}


func HelloFromSource3() {	
	jsonData := getPageBodyJson()
	fmt.Println(WillItRain(8,jsonData))
}