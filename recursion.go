package main

import (
	"fmt"
)

func main() {
	fmt.Println("Count Down starts in :")
	CountDown(3)
	n := 50
	fmt.Printf("Sum is %v\n", sum_01(n))
	fmt.Printf("Sum 02 is %v", sum_02(4))
}

//Exemples.
func CountDown(number int) {
	//fmt.Println("Count Down starts in :")
	if number > 0 {
		fmt.Println(number)
		CountDown(number - 1)
	} else {
		fmt.Println("Count Down Stops.")
	}
}

func sum_01(number int) int {
	if number == 0 {
		return 0
	} else {
		return number + sum_01(number-1)
	}
}
func sum_02(number int) int {
	if number == 0 {
		return 1
	} else {
		return number * sum_02(number-1)
	}
}
