
func addnum() int {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	i := 0
	r := len(numbers)
	for {
		sum += numbers[i]
		i++
		if i >= r {
			break
		}
		fmt.Printf("Sum is %d", sum)

	}
	return 0
}
func addnum2() int {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	fmt.Printf("Sum is %d", sum)
	return 0
}

add_num := addnum()
fmt.Println(add_num)
add_num_2 := addnum2()
fmt.Println(add_num_2)
numrange := []int{10, 20, 30, 40}
for i, k := range numrange {
	fmt.Printf("Range of %d is :: %d\n", i, k)
}
names := []string{"Ali", "Mustapha", "Kamal", "Mohammed", "Hamza"}
for k, elements := range names {
	fmt.Printf("Case %d has a name of %s \n", k, elements)

}
d := [...]int{10, 20, 30, 40, 50}
for a, e := range d {
	fmt.Printf("%d has elements of %d \n", a, e)
}
didi := [...]int{11, 12, 13, 14, 15, 16}
for d := 0; d < len(didi); d++ {
	fmt.Println(didi[d])
}
mimi := [...]string{"Hamza", "Adam", "Sarah"}
for c := 0; c < len(mimi); c++ {
	fmt.Println(mimi[c])
}
//Multi_Dimensional Slices
arr_1 := []string{"Becky", "Emily", "Henry"}
arr_2 := []string{"Cat", "Dog", "Rabbit"}
mult := [][]string{arr_1, arr_2}
fmt.Println(mult)

xixi := []int{80, 90, 100, 110, 120}
for o, xelements := range xixi {
	fmt.Printf("%d has numbers of %d", o, xelements)
}
pipi := []int{130, 140, 150, 160, 170}
for xi := 0; xi <= len(pipi); xi++ {
	y := 0
	fmt.Printf("%d has elements of %d", y, xi)
}