package main

import (
	"fmt"
	"reflect"
)

func main() {
	//fmt.Println("This is a test")
	var x = 50
	fx := float64(x) + 10
	fmt.Printf("x has a value of : %v \n type of x is : %s \n", x, reflect.TypeOf(x))
	fmt.Printf("\n fx has a value of : %v \n type of fx is : %s \n", fx, reflect.TypeOf(fx))
}
