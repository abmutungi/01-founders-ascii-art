package am

import (
	"bufio"
	"fmt"
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

// func SplitNewLine(s string) string {

// 	var slice []string

// 	for i := 0; i < len(s); i++ {

// 		slice = append(slice, string(s[i]))

// 		if slice[i] == "n" && slice[i-1] == "\\" {
// 			fmt.Println("hey")
// 		}
// 	}
// 	return strings.Join(slice, "")
// }

func SplitLines(s string) [][]byte {

	var count int

	for i := 0; i < len(s); i++ {
		if s[i] == 'n' && s[i-1] == '\\' {
			count++
		}
	}
	splitString := []byte(s)
	splitLines := make([][]byte, count+1)

	j := 0

	for i := 0; i < len(splitLines); i++ {
		for j < len(splitString) {
			if splitString[j] == 'n' && splitString[j-1] == '\\' {
				j++
				splitLines[i] = splitLines[i][:len(splitLines[i])-1]
				break
			}
			splitLines[i] = append(splitLines[i], splitString[j])
			j++
		}
	}
	fmt.Println(splitLines)
	return splitLines
}
