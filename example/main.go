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

func canIdrink(age int) bool {
	/*
		if koreanAge := age + 2; koreanAge < 18 {
			return false
		} else {
			return true
		}*/

	/*switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
	*/
	/*
		if age < 18 {
			return false
		} else if age == 18 {
			return true
		} else if age > 50 {
			return false
		}
		return false*/

	switch korAge := age + 2; korAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false

}

func main() {
	arr := []string{"name", "age", "phoneNumber"}
	arr = append(arr, "chiki")
	fmt.Println(arr)

}
