package main

import (
	"fmt"
)

func main() {
	line1 := `"    /\     "`
	line2 := `"   /  \    "`
	line3 := `"  / /\ \   "`
	line4 := `" / ____ \  "`

	storeLine := []string{"A", line1, line2, line3, line4}

	// fmt.Println(storeLine)

	for i := 0; i < 4; i++ {
		if i == 0 {
		}
		fmt.Println(storeLine[i])
	}
}

// map

// 	ArnoldsCharacterMap:= make(map[int][]string)

// 	ArnoldsCharacterMap[65] = storeLine

// 	fmt.Println(ArnoldsCharacterMap)

// 	storeLine2 := []string{"A", line1,line2,line3,line4}

// }
