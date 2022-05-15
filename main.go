package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("Done...")
	length = len(name)
	uppercase = strings.ToUpper(name)
	fmt.Println("함수")
	return
}

func main() {
	totalLength, up := lenAndUpper("corner")
	fmt.Println("----------")
	fmt.Println(totalLength, up)
}
