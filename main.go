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

	if args == "\\n" {
		fmt.Println()
	} else if args != "" {

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

					fmt.Print(charMap[int(splitLines[j][k])][i])
				}
				fmt.Println()
			}
		}
	}
}
