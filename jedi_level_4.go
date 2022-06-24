	//Exercise_01(SUCCESS)
	ex_01 := []int{10, 12, 13, 14, 15}
	for e, f := range ex_01 {
		fmt.Printf("range of %d has a value of %d\n", e, f)
	}
	//Exercise_02(SUCESS)
	ex_02 := []int{100, 110, 120, 130, 140, 150, 160, 170, 180, 190}
	for _, f_02 := range ex_02[:5] {
		fmt.Printf(" %d\n", f_02)

	}
	for _, f_002 := range ex_02[5:10] {
		fmt.Printf("%d\n", f_002)
	}
	//Exercise_03(SUCCESS)
	ex_03 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	ex_03 = append(ex_03, 52)
	fmt.Println(ex_03)
	ex_03 = append(ex_03, 53, 54, 55)
	fmt.Println(ex_03)
	ex_003 := []int{56, 57, 58, 59, 60}
	ex_03 = append(ex_03, ex_003...)
	fmt.Println(ex_03)
	//Exercise_04(SUCCESS)
	ex_04 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	ex_04 = Delete_Slice(ex_04, 3, 4, 5)
	fmt.Println(ex_04)
	//Exercise_05(SUCCESS)
	ex_05 := []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawai", "Idaho", "Ilinois", "Indiana", "Lowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Marryland", "Massachusett", "Michigan", "Minnesota", "Missisipi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dokata", "Ohaio", "Oklahoma", "Oregon", "Pennyslavania", "Rhode Island", "South Carolina", "South Dakota", "Tenessese", "Texas", "Utah", "Vermont", "Virginia", "Washington DC", "West Virginia", "Wisconsin", "Wyoming"}
	for ex_005 := 0; ex_005 < len(ex_05); ex_005++ { //(IF USING <= :: panic: runtime error: index out of range [50] with length 50||CORRECTION :: Just use < instead.)
		fmt.Printf("ID_STATE :: %d has a name of %s\n", ex_005, ex_05[ex_005])
	}
	//Exercise_06(SUCCESS)
	ex_06 := []string{"James", "Bond", "Not Stirred"}
	ex_006 := []string{"Miss", "MoneyPenny", "Hello James"}
	ex_0006 := [][]string{ex_06, ex_006}
	fmt.Println(ex_0006)
	//Exercise_07(SUCCESS)
	ex_07 := map[string]string{
		"Hamza": "Programming",
		"Sarah": "Jogging",
		"Ali":   "Soccer",
	}
	for key, n := range ex_07 {
		fmt.Printf("The name :: %s likes to %s\n", key, n)
	}
	//Exercise_08(Continuation of exercise 07)
	ex_07["Mustapha"] = "Tennis"
	for newkey, new_n := range ex_07 {
		fmt.Printf("The name :: %s likes to %s\n", newkey, new_n)
	}
	//Exercise_09(Continuation to exercise 07)
	delete(ex_07, "Hamza")
	for del_key, del_n := range ex_07 {
		fmt.Printf("The name :: %s likes to %s\n", del_key, del_n)
	}
}

//This Function is related to exercise 4
func Delete_Slice(s []int, index int, index_01 int, index_02 int) []int {
	return append(s[:index], s[index+3:]...)
}