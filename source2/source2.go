package source2

//--------------www.weather-atlas.com--------------

import (
	"fmt"
	"net/http"

	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://www.weather-atlas.com/en/israel/tel-aviv-yafo-long-term-weather-forecast"

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

func getPageContent() *goquery.Document{
	response := getHtml(URL)
	defer response.Body.Close()

	doc, error := goquery.NewDocumentFromReader(response.Body)
	check(error)

	return doc
}

func check(error error) {
	if (error != nil) {
		fmt.Println(error)
	}
}

func getWeatherSummary(doc *goquery.Document) WeatherSummary {
	humidity := getHumidity(doc)
	min := getMinTemp(doc)
	//max := getMaxTemp(doc) //TODO: check if works at morning hours
	chance := getChance(doc)
	wind := getWind(doc)

	return WeatherSummary {min, 0, humidity, wind, chance}
}

func getHumidity(doc *goquery.Document) int {
	r := regexp.MustCompile(`[0-9]+`)
	humidityStr := doc.Find(`div[itemtype="https://schema.org/Table"]>div:first-child>div:nth-child(2)>div.row>div:first-child>ul>li:nth-child(2)`).First().Text()
	h := r.FindString(humidityStr)
	intHumidity, err := strconv.Atoi(h)
	check(err)

	return intHumidity
}

func getMaxTemp(doc *goquery.Document) int {
	r := regexp.MustCompile(`[0-9]+`)
	maxTempStr := doc.Find(`div[itemtype="https://schema.org/Table"]>div:first-child>div:first-child>div:nth-child(2)>ul>li:first-child`).First().Text()
	m := r.FindString(maxTempStr)
	fmt.Println(maxTempStr)
	intMax, err := strconv.Atoi(m)
	check(err)

	return intMax	
}

func getMinTemp(doc *goquery.Document) int {
	r := regexp.MustCompile(`[0-9]+`)
	maxTempStr := doc.Find(`div[itemtype="https://schema.org/Table"]>div:first-child>div:first-child>div:nth-child(2)>ul>li:nth-child(2)`).First().Text()
	m := r.FindString(maxTempStr)
	intMin, err := strconv.Atoi(m)
	check(err)

	return intMin	
}

func getWind(doc *goquery.Document) int {
	windStr := doc.Find(`div[itemtype="https://schema.org/Table"]>div:first-child>div:nth-child(2)>div.row>div:first-child>ul>li:first-child`).First().Text()
	r := regexp.MustCompile(`[0-9]+`)
	w := r.FindString(windStr)

	wind, err := strconv.Atoi(w)
	check(err)

	return wind
}

func getChance(doc *goquery.Document) int{
	r := regexp.MustCompile(`[0-9]+`)
	chanceStr := doc.Find(`div[itemtype="https://schema.org/Table"]>div:first-child>div:nth-child(2)>div.row>div:nth-child(2)>ul>li:first-child`).First().Text()
	c := r.FindString(chanceStr)
	intChance, err := strconv.Atoi(c)
	check(err)

	return intChance
}

func TempArray(days int, doc *goquery.Document) []MinMaxPair {
	r := regexp.MustCompile(`[0-9]+`)
	var minMaxPairs []MinMaxPair

	doc.Find(`div[itemtype="https://schema.org/Table"]>div>div:first-child>div:nth-child(2)>ul`).EachWithBreak(func(index int, item *goquery.Selection) bool {
		maxStr := item.Find(`li:first-child`).First().Text()
		minStr := item.Find(`li:nth-child(2)`).First().Text()

		max := r.FindString(maxStr)
		min := r.FindString(minStr)

		intMax, errMax := strconv.Atoi(max)
		check(errMax)
		intMin, errMin := strconv.Atoi(min)
		check(errMin)
		minMaxPairs = append(minMaxPairs, MinMaxPair{intMin, intMax})

		return index != days - 1
	})

	return minMaxPairs
}

func AverageTemp(days int, doc *goquery.Document) float64 {
	tempArr := TempArray(days, doc)

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

func NextDayRain(doc *goquery.Document) (string, bool) {
	r := regexp.MustCompile(`[0-9]+`)
	//rMonth := regexp.MustCompile(`([a-zA-Z]+)`)

	var dateStr string

	doc.Find(`div[itemtype="https://schema.org/Table"]>div`).EachWithBreak(func(index int, item *goquery.Selection) bool {
		rainChanceStr := item.Find("div>div:nth-child(2)>div.row>div:nth-child(2)>ul>li:first-child").Text()
		intChance, err := strconv.Atoi(r.FindString(rainChanceStr))
		check(err)
		//fmt.Println(intChance)

		if intChance > 50 {
			dateStr = item.Find(`div:first-child>div:first-child>div:first-child`).Text()
			fmt.Println(dateStr)
			return false
		}

		return true
	})

	if (dateStr == "") {
		return "", false
	} else {
		return dateStr, true
	}
}

//Will it rain? - a function that gets a city and a number and returns the chance of rain in this city in the next x days
func WillItRain(days int, doc *goquery.Document) []int{
	r := regexp.MustCompile(`[0-9]+`)
	var chances []int

	doc.Find(`div[itemtype="https://schema.org/Table"]>div`).EachWithBreak(func(index int, item *goquery.Selection) bool {
		rainChanceStr := item.Find("div>div:nth-child(2)>div.row>div:nth-child(2)>ul>li:first-child").Text()
		intChance, err := strconv.Atoi(r.FindString(rainChanceStr))
		check(err)
		fmt.Println(intChance)
		check(err)
		chances = append(chances, intChance)

		return index != days - 1
	})

	return chances
}

func HelloFromSource2() {	
	doc := getPageContent()
	//fmt.Println(getWeatherSummary(doc))
	//fmt.Println(TempArray(4,doc))
	//fmt.Println(AverageTemp(4,doc))
	fmt.Println(NextDayRain(doc))
	//WillItRain(5,doc)

}