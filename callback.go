package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func sphere(vol func(radius float64) float64) {
	fmt.Printf("The volume of the sphere is ::: %v ", vol(7))
}
func main() {
	fmt.Println("Hello world.")
	s_01 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //As long as it's the same variable type we can call it in a func_parameter.
	s := sum(s_01...)                        //Including the same parameter type.
	fmt.Printf("The sum is %v\n", s)
	s_02 := even(sum, s_01...)
	fmt.Printf("The even value is %v\n", s_02)
	s_03 := odd(sum, s_01...)
	fmt.Printf("The odd number is %v\n", s_03)
	//Callback_02
	volume_of_sphere := func(radius float64) float64 {
		result := 4 / 3 * math.Pi * radius * radius * radius
		return result
	}
	sphere(volume_of_sphere)
	//
	var seqLen int
	fmt.Scan(&seqLen)
	n := make(square, seqLen)
	n.ReadInput(os.Stdin)
	switch {
	case n.ExamineProgression(ConditionalMap["AP"]):
		fmt.Println("AP")
	case n.ExamineProgression(ConditionalMap["GP"]):
		fmt.Println("GP")
	default:
		fmt.Println("RANDOM ! ! !")
	}
}

//Callback_01
func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
func even(f func(xi ...int) int, yi ...int) int {
	var yi_01 []int
	for _, y := range yi {
		if y%2 == 0 {
			yi_01 = append(yi_01, y)
		}
	}
	return f(yi_01...)
}
func odd(f func(xi ...int) int, bi ...int) int {
	var b_01 []int
	for _, b := range bi {
		if b%2 != 0 {
			b_01 = append(b_01, b)
		}
	}
	return f(b_01...)
}

//
var ConditionalMap = map[string]func(int, int) int{
	"AP": func(i1 int, i2 int) int {
		return i1 - i2
	},
	"GP": func(i1, i2 int) int {
		return i1 / i2
	},
}

type square []int

func (s square) ReadInput(r io.Reader) {
	for j := 0; j < len(s); j++ {
		fmt.Fscan(r, &s[j])
	}
}
func (s square) ExamineProgression(condition func(a, b int) int) bool {
	{
		if len(s) < 3 {
			return false
		}
		d := condition(s[1], s[0])
		for i := 2; i < len(s)-1; i++ {
			if condition(s[i+1], s[0]) != d {
				return false
			}
		}
	}
	return true
}
