package main

import (
	"fmt"
)

type person struct {
	first_name                string
	last_name                 map[string]string
	favorite_ice_cream_flavor []string
}
type vehicul struct {
	doors int
	color string
}
type truck struct {
	vehicul
	fourWheel bool
}
type sedan struct {
	vehicul
	luxury bool
}

func main() {
	//Exercise_01(SUCCESS)
	n1 := person{
		first_name: "Hamza",
		last_name: map[string]string{
			"Lekhal":   "The black",
			"Bouazza":  "Father of glory",
			"Mostaqim": "The straighter",
			"Kfaiti":   "The fapper",
		},
		favorite_ice_cream_flavor: []string{
			"Chocolate", "Strawberry", "Orange", "Watermellon",
		},
	}
	for _, n := range n1.last_name {
		fmt.Printf("%v\n", n)
	}
	//Exercise_03
	c_01 := truck{
		vehicul:   vehicul{doors: 4, color: "Red"},
		fourWheel: false,
	}
	c_02 := sedan{
		vehicul: vehicul{doors: 5, color: "White"},
		luxury:  true,
	}
	fmt.Println(c_01)
	fmt.Println(c_02)
	switch c_01.fourWheel {
	case false:
		fmt.Printf("The car has %v and it's color is %v but it does not have Four Wheels.\n", c_01.vehicul.doors, c_01.vehicul.color)
	default:
		fmt.Println("UNDER_CONSTRUCTION")
	}
	switch c_02.luxury {
	case false:
		fmt.Printf("The sedan have %v and it's color is %v but it looks cheap !", c_02.doors, c_02.color)
	case true:
		fmt.Printf("The sedan have %v and it's color is %v and it is considered as Luxury car\n", c_02.doors, c_02.color)
	}
	//Exercise_04
	my_infos := struct {
		bloodtype         string
		height            float32
		weight            int
		eye_color         string
		favorite_clothing []string
	}{
		bloodtype: "O-",
		height:    1.89,
		weight:    97,
		eye_color: "Green",
		favorite_clothing: []string{
			"Black Leather Jacket", "Para Jeans", "Classical Watch", "Sun Glasses", "Addidas Sneakers",
		},
	}
	fmt.Printf("Hamza has a blootype of %v \nwhile it's height is %vm and it's wieght is %v kg his eye colors are %v and his favorite clothings are %v", my_infos.bloodtype, my_infos.height, my_infos.weight, my_infos.eye_color, my_infos.favorite_clothing)
}



package main

import (
	"fmt"
)

//Introduction de Struct
type person_name struct {
	first_name string
	last_name  string
}
type shop_list struct {
	person_name
	shop_item   string
	shop_price  float32
	credit_card bool
}

func main() {
	fmt.Println("D")
	n1 := person_name{
		first_name: "Hamza",
		last_name:  "Lekhal",
	}
	n2 := person_name{
		first_name: "Mohammed",
		last_name:  "Ali",
	}
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Printf("First person name is %v\nSecond Person name is %v\n", n1.first_name, n2.first_name)
	shop1 := shop_list{
		shop_item:  "Iphone 13 Pro Max",
		shop_price: 899.99,
	}
	shop2 := shop_list{
		shop_item:  "BackPack 50l",
		shop_price: 89.99,
	}
	shop3 := shop_list{
		shop_item:  "Addidas Jogging Sneakers",
		shop_price: 69.99,
	}
	shop_01 := shop_list{
		person_name: person_name{
			first_name: "Hamza",
		},
		shop_item:   "Nvidia RTX 3070",
		shop_price:  1200,
		credit_card: false,
	}
	shop_02 := shop_list{
		person_name: person_name{
			first_name: "Mohammed",
		},
		shop_item:   "Full HD Smart TV",
		shop_price:  499.99,
		credit_card: true,
	}
	t := shop1.shop_price + shop2.shop_price + shop3.shop_price
	fmt.Printf("You have bought : %v and %v and %v\n Your total cost is %v$\n", shop1.shop_item, shop2.shop_item, shop3.shop_item, t)
	//
	fmt.Printf("Dear %v,\nYou are about to buy %v for %v$ however you cannot pay with a Credit Card due to some restrictions.\n", shop_01.person_name.first_name, shop_01.shop_item, shop_01.shop_price)
	fmt.Printf("Dear %v,\nYou are about to buy %v for %v$ You can either pay with cash or with Credit Card.\n", shop_02.person_name.first_name, shop_02.shop_item, shop_02.shop_price)
	//Annonymous_Struct
	car := struct {
		car_name  string
		car_model int
	}{
		car_name:  "Mercedes",
		car_model: 1985,
	}
	bike := struct {
		bike_name  string
		bike_model int
	}{
		bike_name:  "Bennilli",
		bike_model: 2001,
	}
	fmt.Printf("The car's name is %v with a model of %v\n", car.car_name, car.car_model)
	fmt.Printf("The bike's name is %v with a model of %v\n", bike.bike_name, bike.bike_model)
	//Exercise_01(SUCCESS)
	ice_cream_buyer := struct {
		first_name                string
		last_name                 string
		favorite_ice_cream_flavor string
	}{
		first_name:                "Hamza",
		last_name:                 " Lekhal",
		favorite_ice_cream_flavor: "Chocolate",
	}
	ice_cream_buyer_01 := struct {
		first_name                string
		last_name                 string
		favorite_ice_cream_flavor string
	}{
		first_name:                "Sarah",
		last_name:                 " Luv",
		favorite_ice_cream_flavor: "Vanilla",
	}
	ice_01 := ice_cream_buyer.first_name + ice_cream_buyer.last_name
	ice_02 := ice_cream_buyer_01.first_name + ice_cream_buyer_01.last_name
	fmt.Printf("%v's favorite ice cream flavor is %v\n", ice_01, ice_cream_buyer.favorite_ice_cream_flavor)
	fmt.Printf("%v's favorite ice cream flavor is %v", ice_02, ice_cream_buyer_01.favorite_ice_cream_flavor)
	//Exercise_02
	/*account_infos := map[string]string{ 
		first_name string
		last_name string
		first_name_01 string
		last_name_01 string
		first_name_02 string
		last_name_02 string
		favorite_drinks string
		favorite_drinks_01 string
		favorite_drinks_02 string
	}{
		first_name "Rebecca",
		last_name "Jones",
		first_name_01 "Jessica",
		last_name_01 "Night",
		first_name_02 "Johnny",
		last_name_02 "Sins",
		favorite_drinks "Apple Juice"
		favorite_drinks_01 "Tequila whiskey"
		favorite_drink_02 "Coffee"

	}*/
