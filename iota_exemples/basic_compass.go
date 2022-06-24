package main

import "fmt"

type direction int

const (
	north direction = iota
	south
	west
	east
)

func (compass direction) string() string {
	return [...]string{"North", "South", "West", "East"}[compass]
}

func main() {
	fmt.Println("Basic Compass.")
	var compass direction = south
	fmt.Print(compass)
	switch compass {
	case north:
		fmt.Println("The compass is showing north direction")
	case south:
		fmt.Println("The compass is showing south direction")
	case west:
		fmt.Println("The compass is showing west direction")
	case east:
		fmt.Println("The compass is showing east direction")
	}
}
