package am

import (
	"fmt"
	"os"
)

func ReadFile() string {
	fileOne, err := os.ReadFile("ascii.txt")
	if err != nil {
		fmt.Println("Invalid input")
	}
	return string(fileOne)
}

func MakeMap() {
	m := make(map[string]int)

	start := 32
	line := 1

}
