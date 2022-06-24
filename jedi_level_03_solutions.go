	/*Exercise_1 (SUCCESS)
	for x := 0; x <= 10000; x++ {
		fmt.Printf("Generating...%d \n", x)
	}*/
	//Execise_2 (SUCCESS)
	/*var x int
	for x = 65; x <= 90; x++ {
		fmt.Println(x)
		for n := 0; n < 3; n++ {
			fmt.Printf("%#U\n", x)
		}

	}*/
	//Exercise_3 (SUCCESS)
	/*l := 1995
	for y := 2022; y <= 2022; y-- {
		t := y - l
		fmt.Printf("You have lived for %d years.", t)
		break
	}*/
	//Exercise_4 (SUCCESS)
	/*for {
		y := 2022
		l := 1995
		t := y - l
		fmt.Printf("You have lived for %d years.", t)
		break
	}*/
	//Exercise_5 (FAILED!) I couldn't understand the exercise.
	/*for x := 10; x <= 100; x++ {
		fmt.Printf("The divided by 4 of value %d \t has a remainder of %d\n", x, x%4)
	}*/
	//Exercise_6 (SUCCESS)
	/*var name string
	var age int
	fmt.Printf("Enter your name: %s\n", name)
	fmt.Scanf("%s\n", &name)
	fmt.Printf("Enter your age: ")
	fmt.Scanf("%d", &age)
	if name == "Hamza" && age == 27 {
		fmt.Printf("Welcome %s , your current age is %d", name, age)
	} else {
		fmt.Println("Incorrect username.")
	}*/
	//Exercise_7 (SUCCESS)
	/*var name string
	var age int
	fmt.Println("Enter your name:")
	fmt.Scanf("%s\n", &name)
	fmt.Println("Enter your age:")
	fmt.Scanf("%d", &age)
	if name == "Hamza" && age == 27 {
		fmt.Printf("Your name is %s and your current age is %d", name, age)
	} else if name == "Ali" && age == 28 {
		fmt.Printf("Your name is %s and your current age is %d", name, age)
	} else {
		fmt.Println("Unknown user in the database.")
	}*/
	//Exercise_8 (SUCCESS)
	/*a := 10
	switch {
	case a == 10:
		fmt.Println("The number is ten.")
		fallthrough
	case a == 20:
		fmt.Println("The number is twenty.")
		fallthrough
	case a == 30:
		fmt.Println("The number is thirty.")
	}*/
	//Exercise_9 (SUCCESS)
	/*var FavSport string
	FavSport = "Jogging"
	switch {
	case FavSport == "Football":
		fmt.Println("Your favorite sport is Football.")
	case FavSport == "Surfing":
		fmt.Println("Your favorite sport is Surfing.")
	case FavSport == "Jogging":
		fmt.Println("Your favorite sport is jogging.")
	}*/
