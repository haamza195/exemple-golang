member_db := map[string]int{
	"Jake":     32,
	"Sarah":    20,
	"Ali":      19,
	"Hamza":    27,
	"Mustapha": 29,
}
memrank_db := map[string]string{
	"Jake":     "Member",
	"Sarah":    "Member",
	"Ali":      "Moderator",
	"Hamza":    "Administrator",
	"Mustapha": "Super Moderator",
}
fmt.Printf("The member list in the DB is %v\n The rank in the DB is %v\n", member_db, memrank_db)
x, ok := memrank_db["Hamza"]
fmt.Println(x, ok) //Search for Administrator
if x, ok := memrank_db["Hamza"]; ok {

	fmt.Printf("The member %v has a rank of %v", member_db["Hamza"], x)
}