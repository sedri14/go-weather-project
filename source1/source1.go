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
	minTemp float64
	maxTemp float64
	humidity int
	wind int
	chanceOfRain int
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

func getWeatherSummary(doc *goquery.Document) {
	humidity := getHumidity(doc)
	fmt.Println(humidity)

	min, max := getMinMaxTemp(doc)
	fmt.Println(min, max)

	chance := getChance(doc)
	fmt.Println(chance)

	wind := getWind(doc)
	fmt.Println(wind)

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

func HelloFromSource1()  {	
	doc := getPageContent()
    getWeatherSummary(doc)
}

