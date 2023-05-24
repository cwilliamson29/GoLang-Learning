package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []int

	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 2)

	log.Println(mySlice)

	sort.Ints(mySlice)

	log.Println(mySlice)

	//rand.Intn()
}
