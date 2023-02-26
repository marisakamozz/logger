package logger

import (
	"fmt"
)

var Version string = "1.0"

func Log(message string, count int) {
	for i := 0 ; i < count; i++ {
		fmt.Println("[LOG] " + message)
	}
}
