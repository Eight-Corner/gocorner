package main

import "fmt"

func sayBye() {
	fmt.Println("Bye")
}

func SayHello() {
	fmt.Println("Hello")
}

func superAdd(numbers ...int) int {
	/*for index, number := range numbers {
		//	range가 number 안에서 조건에 따라 반복문을 실행 하도록 해줍니다.
		//	자바스크립트의 forEach처럼 말이죠!
		fmt.Println(number, index)
	}*/

	/*	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}*/
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := superAdd(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(total)
}
