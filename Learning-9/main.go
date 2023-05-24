package main

import (
	"log"
)

type User struct {
	Fname string
	Lname string
}

func main() {
	mymap := make(map[string]User)

	me := User{
		Fname: "Chris",
		Lname: "Williamson",
	}
	you := User{
		Fname: "Sandra",
		Lname: "Bullock",
	}

	mymap["me"] = me
	mymap["you"] = you

	log.Println(mymap)
	log.Println(mymap["me"])
	log.Println(mymap["me"].Fname)
}
