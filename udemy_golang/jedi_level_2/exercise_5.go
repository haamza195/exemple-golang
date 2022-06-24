package main

import "fmt"

func main() {
	a := `This is a raw string literal \n `
	b := "\n This is an interpreted string"
	fmt.Printf("%s %s", a, b)
}
