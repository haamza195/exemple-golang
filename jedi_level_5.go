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
