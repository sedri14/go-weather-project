package source1

import (
	"fmt"
)

// Hello returns a greeting for the named person.
func HelloFromSource1() string {	
    message := fmt.Sprintf("Hello from source1")
    return message
}