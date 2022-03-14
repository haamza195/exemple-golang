package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("This is a test go file.")
	Hello()
}
func Hello() string {
	return quote.Hello()
}
