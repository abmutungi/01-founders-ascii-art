package am

import (
	"bufio"
	"os"
)

// func ReadFile() string {
// 	fileOne, err := os.ReadFile("ascii.txt")
// 	if err != nil {
// 		fmt.Println("Invalid input")
// 	}
// 	return string(fileOne)
// }

// func MakeMap() {
// 	m := make(map[string]int)

// 	start := 32
// 	line := 1

// }
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}