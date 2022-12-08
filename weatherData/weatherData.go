package weatherData

import (
	"fmt"
"example.com/source1"
"example.com/source2"
"example.com/source3"
)

// Hello returns a greeting for the named person.
func HelloFromWeatherData() string {
    // Return a greeting that embeds the name in a message.
	messageSource1 := source1.HelloFromSource1()
	messageSource2 := source2.HelloFromSource2()
	messageSource3 := source3.HelloFromSource3()
    message := fmt.Sprintf(messageSource1 + messageSource2 + messageSource3)
    return message
}