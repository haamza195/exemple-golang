package main

import (
	"fmt"
	"strings"
)

func main() {

	var stars string = strings.Repeat("*", 10)
	return fmt.Sprintf("%s \nWelcome ! \n%s", stars, stars)
}
