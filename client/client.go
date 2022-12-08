package main

import (
    "fmt"

    "example.com/weatherData"
)

func main() {
    // Get a greeting message and print it.
    message := weatherData.HelloFromWeatherData()
    fmt.Println(message)
}