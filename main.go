package main

import "fmt"

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("corner", "developer", "moon", "suwon")
}
