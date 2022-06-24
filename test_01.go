package main

import "fmt"

func main() {
	fmt.Println("Hello world.")
	//Closure_01
	n := numSeq()
	fmt.Printf("The closure value is %v\n", n())
	fmt.Printf("The next closure value is %v\n", n())
	fmt.Printf("The next closure value is %v\n", n())
	//Closure_01
	pos, neg := addr(), addr()
	for i := 0; i < 10; i++ {
		fmt.Printf("The pos value is %v |||| The neg value is %v\n", pos(i), neg(-2*i))
	}
}

//Closure_01
func numSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//Closure_02
func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
