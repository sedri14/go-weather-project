package weatherData

import (
"example.com/source1"
)

// Hello returns a greeting for the named person.
func HelloFromWeatherData() string {
    // Return a greeting that embeds the name in a message.
	message := source1.HelloFromSource1()
    //message := fmt.Sprintf("Hello from weather data")
    return message
}