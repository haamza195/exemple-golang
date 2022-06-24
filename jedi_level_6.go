package main

import "fmt"

func main() {
	fmt.Println("Jedi Level 06 Exercises.")
	//Exercise_01
	a, b := bar(27, "Hamza\n")
	fmt.Printf("Exercise 01 results \n %v\n,%v %v", foo(5), a, b)
	//Exercise_02
	a_01 := foo_02(1, 2, 3, 4, 5, 6, 7, 8, 9)
	a_02 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	a_03 := bar_02(a_02)
	fmt.Println("Exercise 02 results\n", a_01)
	fmt.Println(a_03)
	//Exercise_03
	fmt.Println("Exercise 03\n")
	numbers(60, 80)
	//defer numbers(90, 100)
	show()
	//Exercise 04
	fmt.Println("Exercise 04\n")
	a_04 := infos{
		first_name: "Hamza",
		last_name:  "Lekhal",
		age:        27,
	}
	a_04.infos_01()
}

//Exercise_01 ( SUCCESS)
func foo(n int) int {
	return n
}
func bar(n int, p string) (int, string) {
	return n, p
}

//Exercise_02(SUCESS)
func foo_02(n ...int) []int {
	return n
}
func bar_02(n []int) []int {
	return n
}

//Exercise_03(SUCESS)
func numbers(n int, p int) int {
	result := n * p
	fmt.Printf("Result is : %v\n", result)
	return result
}
func show() {
	fmt.Println("This outplaced the defered function.\n")
}

//Exercise_04(SUCESS)
type infos struct {
	first_name string
	last_name  string
	age        int
}

func (a infos) infos_01() {
	fmt.Printf("Your first name is :%v\n", a.first_name)
	fmt.Printf("Your last name is :%v\n", a.last_name)
	fmt.Printf("Your current age is :%v\n", a.age)
}
