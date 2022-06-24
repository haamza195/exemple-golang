package main

import (
	"fmt"
	"math"
)

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
	//Exercise_05
	fmt.Println("Exercise 05\n")
	a_05 := circle{
		radius: 6.898123,
	}
	a_06 := square_01{
		lenght: 20,
		width:  30,
	}
	fmt.Printf("The area of circle is %v\n", a_05.area())
	fmt.Printf("The area of the square is %v \n", a_06.area())
	//Exercise_06
	fmt.Println("Exercise 06 - Annonymous func\n")
	a_07 := func(n int) int {
		return n * (n - 1)
	}
	fmt.Printf("The annonymous func is %v\n", a_07(5))
	//Exercise_08
	fmt.Println("Exercise 08\n")
	fmt.Printf("The returned func value is %v \n", ret(5))
	//Exercise_09
	fmt.Println("Exercise 09\n")
	volume_of_circle_02 := func(radius float64) float64 {
		return math.Pi * (radius * radius)
	}
	area_02(volume_of_circle_02)
	//Exercise_10
	fmt.Println("Exercise 10 \n")
	nextseq := seq()
	fmt.Printf("Closure value is :: %v\n", nextseq())
	fmt.Printf("Closure value is :: %v\n", nextseq())
	fmt.Printf("Closure value is :: %v\n", nextseq())
	fmt.Printf("Closure value is :: %v\n", nextseq())
	fmt.Printf("Closure value is :: %v\n", nextseq())
	fmt.Println("Resetting...\n")
	newseq := seq()
	fmt.Printf("Closure value is :: %v", newseq())
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

//Exercise_05(SUCCESS)
type circle struct {
	radius float64
}
type square_01 struct {
	lenght float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)

}
func (s square_01) area() float64 {
	return s.lenght * s.width

}

type shape interface {
	area()
}

func area_circ(pi shape) {
	pi.area()
}

//Exercise_06(SUCCES - LOOK AT MAIN FUNC)--------------------------------------------------------------------------------------------------------------------------------
//Exercise_07(FAILED.)
//Exercise_08 (INCLUDING EXERCISE_11)
func ret(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * ret(n-1)
	}
}

//Exercise_09(SUCCESS)
func area_02(volume_01 func(radius float64) float64) {
	fmt.Printf("The volume area of the second circle is %v\n", volume_01(7.897912))
}

//Exercise_10(SUCCESS)
func seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//Summary -- I still need more practices over in order to learn func and master it.
