package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	// totalLength, upperName := lenAndUpper("corner")
	// fmt.Println(totalLength, upperName)
	totalLength, _ := lenAndUpper("corner")
	fmt.Println(totalLength)
}
