package main

import (
	"fmt"
	"log"
	"git.learn.01founders.co/abmutungi/ascii-art-test.git/am"
)

func main() {
	lines, err := am.ReadLines("ascii.txt")
	if err != nil {
		log.Fatalf("ReadLines: %s", err)
	}
	fmt.Println(lines)
}
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
