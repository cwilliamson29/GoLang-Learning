package main

import "fmt"

var x = 42

const y = 84

func main() {
	z := 21

	fmt.Println(x, y, z)
	fmt.Printf("%v, %v, %v \n", x, y, z)
}
