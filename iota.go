package main

import (
	"fmt"
)

type BitSize float64

const (
	KB BitSize = 1 << (10 * (iota + 1))
	MB
	GB
)

func (b BitSize) String() string {
	switch {
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)

	}
	return fmt.Sprintf("%.2fB", b)
}
func main() {
	fmt.Println("Data bytes converter")
	fmt.Println(1001*KB, 900*MB, 7*GB)
	fmt.Println(BitSize(9837498900))
}
