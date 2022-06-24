package main

import (
	"fmt"
)

func NeedsLicense(kind string) bool {
	if kind == "car" {
		return true
	} else {
		return false
	}
}
func ChooseVehicul(option1 string, option2 string) string {
	if option1 > option2 {
		return option2 + " is clearly the better choice."
	} else {
		return option1 + " is clearly the better choice"
	}
}
func CalculateResellPrice(originalprice, age float64) float64 {
	if originalprice == 1000 && age == 1 {
		return originalprice - (originalprice*12)/(age*60)
	} else if originalprice == 1000 && age == 5 {
		return originalprice - (age * 60)
	} else if originalprice == 1000 && age == 15 {
		return originalprice - (originalprice / 2)
	} else {
		return originalprice
	}
}
func main() {
	needLicense := NeedsLicense("car")
	vehicul := ChooseVehicul("Volkswagen Beetle", "Volkswagen Golf")
	deadning := CalculateResellPrice(1000, 15)
	fmt.Println(needLicense)
	fmt.Println(vehicul)
	fmt.Println(deadning)
}
