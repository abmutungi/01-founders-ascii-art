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

/*
package main

import (
	"fmt"
)

func main() {
	line1:= `"    /\     "`
	line2:= `"   /  \    "`
	line3:= `"  / /\ \   "`
	line4:= `" / ____ \  "`


	storeLine := []string{, line1,line2,line3,line4}

	//fmt.Println(storeLine)

	for i:= 0; i<4;i++{
	if i==0{

	}
	fmt.Println(storeLine[i])
	}



	//map

	ArnoldsCharacterMap:= make(map[int][]string)

	ArnoldsCharacterMap[65] = storeLine


	fmt.Println(ArnoldsCharacterMap)


	storeLine2 := []string{"A", line1,line2,line3,line4}


}
*/
