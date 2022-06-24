package main

import "fmt"

type day string

const (
	Monday = iota + 1
	Tuesday
	Wednsday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println("Days name during a week.")
}
