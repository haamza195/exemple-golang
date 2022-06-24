package main

import "fmt"

func main() {
	fmt.Println("Exercise 4")
	x := 554346
	fmt.Printf("x's value in decimal is %d\n x's value in binary is %b\n x's value in hex is %#x \n", x, x, x)
	y := x << 1
	fmt.Printf("The shifted value in decimal is %d\n In binary is %b\n In hex is %#x", y, y, y)
}
