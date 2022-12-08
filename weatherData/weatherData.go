package weatherData

import "fmt"

// Hello returns a greeting for the named person.
func HelloFromWeatherData() string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hello from weather data")
    return message
}