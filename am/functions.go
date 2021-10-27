package am

import (
	"fmt"
	"os"
)

func ReadFile() string {
	fileOne, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input")
	}
	return string(fileOne)
}
