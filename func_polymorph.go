package main

import "fmt"

type a_drive interface {
	driving_01()
	car_review()
}

//Driving methods
type person_001 struct {
	first_name string
	last_name  string
}
type drive_01 struct {
	person_001
	license bool
}

//
type vehicul_01 struct {
	vehicul_name  string
	vehicul_model int
}
type vehicul_new struct {
	vehicul_01
	is_Used bool
}

//Function_polymorph
func (driver_01 drive_01) driving_01() {
	switch driver_01.license {
	case true:
		fmt.Printf("The attendant %v %v is able to drive a car.\n", driver_01.person_001.first_name, driver_01.person_001.last_name)
	default:
		fmt.Printf("The attendant %v %v is not being able to drive a car.\n", driver_01.person_001.first_name, driver_01.person_001.last_name)
	}

}
func (car_rev vehicul_new) car_review() {
	car_age := 2022 - car_rev.vehicul_model
	switch car_rev.is_Used {
	case true:
		if car_age >= 10 {
			fmt.Printf("The car %v is %v years old which makes it used and old.\n", car_rev.vehicul_name, car_age)
		} else if car_age <= 5 {
			fmt.Printf("The car %v is %v years old which makes it newly used.\n", car_rev.vehicul_name, car_age)
		}
	case false:
		if car_age >= 10 {
			fmt.Printf("The car %v is %v years old has not been used by it's owner but it is kinda old.\n", car_rev.vehicul_name, car_age)
		} else if car_age <= 5 {
			fmt.Printf("The car %v is %v years old has not been used by it's owner and it is kinda new!!!\n", car_rev.vehicul_name, car_age)
		}
	}

}

func main() {
	fmt.Println("This is a polymorth func re-exercise.")
	//This is for driving methods
	person_27 := drive_01{
		person_001: person_001{first_name: "Hamza", last_name: "Lekhal"},
		license:    false,
	}
	person_28 := drive_01{
		person_001: person_001{first_name: "Mohammed", last_name: "Ali"},
		license:    true,
	}
	//This is for vehicul reviewing.
	vehicul_a := vehicul_new{
		vehicul_01: vehicul_01{vehicul_name: "Mercedes", vehicul_model: 2001},
		is_Used:    true,
	}
	vehicul_b := vehicul_new{
		vehicul_01: vehicul_01{vehicul_name: "Audi", vehicul_model: 2018},
		is_Used:    false,
	}
	//Printing out method interfaces.
	person_27.driving_01()
	person_28.driving_01()
	vehicul_a.car_review()
	vehicul_b.car_review()
}
