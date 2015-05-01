package log

import (
	"fmt"
)

func Info(msg ...interface{}) {
	fmt.Printf(fmt.Sprintf("Info: %s ", fmt.Sprint(msg...)))
	fmt.Println()

}

func Error(msg ...interface{}) {
	fmt.Printf(fmt.Sprintf("Error: %s ", fmt.Sprint(msg...)))
	fmt.Println()

}
