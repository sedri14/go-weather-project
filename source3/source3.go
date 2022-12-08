package source3

import (
	"fmt"
)

// Hello returns a greeting for the named person.
func HelloFromSource3() string {	
    message := fmt.Sprintf("Hello from source3")
    return message
}