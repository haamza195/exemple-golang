package main

import "fmt"

type year float64

const (
	one year = 2022 + iota
	two
	three
	four
	five
	six
)

func main() {
	fmt.Println("Exercise 6")
	fmt.Println(one, two, three, four, five, six)
}
