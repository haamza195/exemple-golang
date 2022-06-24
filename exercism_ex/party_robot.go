package main

import "fmt"

func Welcome(name string) string {
	return fmt.Sprintf("Welcome %s", name)
}
func Happybirthday(name string, age int) string {
	return fmt.Sprintf("\nHappy birthday, %s\nYou are now %d years old!", name, age)
}
func AssignTable(name string, table int, seatmate string, direction string, distance float64) string {
	return fmt.Sprintf("\nWelcome to my party again, %s!\nYour table number is %03d. Your table is %s,which is %.1f meters away from your neighbor %s.", name, table, direction, distance, seatmate)
}
func main() {
	welcome := Welcome("Hamza")
	age := Happybirthday("Hamza", 27)
	table := AssignTable("Hamza", 17, "Ali", "on the upper corner", 3.13123123)
	fmt.Printf(welcome)
	fmt.Printf(age)
	fmt.Printf(table)
}
