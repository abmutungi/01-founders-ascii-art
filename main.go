package main

import (
	"fmt"
	"strings"

	"git.learn.01founders.co/abmutungi/ascii-art-test.git/am"
)

func main() {
	sliceOfStr := strings.Split(am.ReadFile(), " ")
	fmt.Println(sliceOfStr)
}
