package main

import (
	"fmt"
	"log"
	"os"

	// "strings"

	"git.learn.01founders.co/abmutungi/ascii-art-test.git/am"
)

func main() {
	args := os.Args[1]

	splitLines := am.SplitLines(args)

	lines, err := am.ReadLines("ascii.txt")
	if err != nil {
		log.Fatalf("ReadLines: %s", err)
	}

	charMap := make(map[int][]string)

	start := 32

	for i := 0; i < len(lines); i++ {
		if len(charMap[start]) == 9 {
			start++
		}
		charMap[start] = append(charMap[start], lines[i])
	}

	for j, val := range splitLines {
	for i := 1; i < 9; i++ {
			for k := 0; k < len(val); k++ {
				// if i == 8 && k == len(val)-1 {
				// 	fmt.Println("hey")
				// }
				fmt.Print(charMap[int(splitLines[j][k])][i])
			}

			// if j == len(splitLines)-1 {
			// 	fmt.Println()
			// }
			fmt.Println()
		}
	}
}

// for i := 1; i < 9; i++ {
// 	for j := range args {
// 		fmt.Print(charMap[int(args[j])][i])
// 	}
// 	fmt.Println()
// }

/*
package main

import (
	"fmt"
)

func main() {

	myMap := make(map[int][]string)

	temp:= []string{"ARNOLD"} //type of this is slice of string
	another := []string{"MJR"}

	myMap[10] = temp

	fmt.Println(myMap[10])


	//////////////////

	myMap[20] = another

	fmt.Println(myMap[20])

	fmt.Println(myMap)

	myMap[20] = []string{"Nik"}

	fmt.Println(myMap)


	//how do u add to an existing slice


	temp = append(temp, "secondstring")

	fmt.Println(temp)

	//how do u add to a map??

	myMap[10] = append (myMap[10], "secondstring")
	fmt.Println(myMap[10])



	////////////

	start:= 35
	for i:=0; i<50;i++{
		myMap[start] = append (myMap[start], "line")
		if i== 10 || i==25{
			start ++
		}


	}


	for _, line := range myMap[10]{
		fmt.Println(line)
	}

	fmt.Println(myMap)
}
*/
