package logger

import (
	"fmt"
)

var Version string = "1.0"

func Log(message string) {
	fmt.Println("[LOG] " + message)
}
