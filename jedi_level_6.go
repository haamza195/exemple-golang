package main

import "fmt"

func main() {
	fmt.Println("Jedi Level 06 Exercises.")
	//Exercise_01
	a, b := bar(27, "Hamza")
	fmt.Printf("Exercise 01 results \n %v\n,%v %v", foo(5), a, b)
}

//Exercise_01 ( SUCCESS)
func foo(n int) int {
	return n
}
func bar(n int, p string) (int, string) {
	return n, p
}

//Exercise_02
func foo_02(n ...int) []int {
	return n
}
