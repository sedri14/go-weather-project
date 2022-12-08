package source1

//--------------timeanddate.com--------------

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://www.timeanddate.com/weather/israel/tel-aviv/ext"

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
	//fmt.Println(humidity)
	min, max := getMinMaxTemp(doc)
	//fmt.Println(min, max)
	chance := getChance(doc)
	//fmt.Println(chance)
	wind := getWind(doc)
	//fmt.Println(wind)

	return WeatherSummary {min, max, humidity, wind, chance}
}

func getHumidity(doc *goquery.Document) int {
	humidityStr := doc.Find("#wt-ext>tbody>tr:first-child>td:nth-child(8)").First().Text()
	h := humidityStr[:len(humidityStr)-1]
	intHumidity, err := strconv.Atoi(h)
	check(err)

	return intHumidity
}

func getMinMaxTemp(doc *goquery.Document) (int, int) {
	minMaxStr := doc.Find("#wt-ext>tbody>tr:first-child>td:nth-child(3)").First().Text()
	r := regexp.MustCompile(`[0-9]+`)
	split := r.FindAllString(minMaxStr, 2)

	max, errMax := strconv.Atoi(split[0])
	check(errMax)
	min, errMin := strconv.Atoi(split[1])
	check(errMin)

	return min,max
}


func getChance(doc *goquery.Document) int{
	chanceStr := doc.Find("#wt-ext>tbody>tr:first-child>td:nth-child(9)").First().Text()
	c := chanceStr[:len(chanceStr)-1]
	intChance, err := strconv.Atoi(c)
	check(err)

	return intChance
}

func getWind(doc *goquery.Document) int {
	windStr := doc.Find("#wt-ext>tbody>tr:first-child>td:nth-child(6)").First().Text()
	r := regexp.MustCompile(`[0-9]+`)
	split := r.FindAllString(windStr, 1)

	wind, err := strconv.Atoi(split[0])
	check(err)

	return wind
}

func TempArray(days int, doc *goquery.Document) []MinMaxPair {
	r := regexp.MustCompile(`[0-9]+`)
	var minMaxPairs []MinMaxPair

	doc.Find("#wt-ext>tbody>tr>td:nth-child(3)").EachWithBreak(func(index int, item *goquery.Selection) bool {
		pair := item.Text()
		fmt.Println(pair)

		split := r.FindAllString(pair, 2)
		max, errMax := strconv.Atoi(split[0])
		check(errMax)
		min, errMin := strconv.Atoi(split[1])
		check(errMin)
		minMaxPairs = append(minMaxPairs, MinMaxPair{min, max})

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

//TODO: fix this. decide a unified format for date //TODO:
//Next rain day - a function that returns the next day with a chance for rain over 50% in a specified city
//if no rainy day is found in near future, bool result is false
func NextDayRain(doc *goquery.Document) (string, bool) {
	r := regexp.MustCompile(`[0-9]+`)
	//rMonth := regexp.MustCompile(`([a-zA-Z]+)`)

	var dateStr string

	doc.Find("#wt-ext>tbody>tr").EachWithBreak(func(index int, item *goquery.Selection) bool {
		rainChanceStr := item.Find("td:nth-child(9)").Text()
		intChance, err := strconv.Atoi(r.FindString(rainChanceStr))
		check(err)
		fmt.Println(intChance)

		if intChance > 50 {
			dateStr = item.Find("th").Text()
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

func WillItRain(doc *goquery.Document) {
	
}

func HelloFromSource1()  {	
	doc := getPageContent()
    // w := getWeatherSummary(doc)
	// fmt.Println(w)
	// t := TempArray(3, doc)
	// fmt.Println(t)
	// a := AverageTemp(3,doc)
	// fmt.Println(a)
	r,ok := NextDayRain(doc)
	fmt.Println(r, ok)
}

