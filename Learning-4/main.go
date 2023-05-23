package main

import (
	"fmt"
	"math/rand"
)

func init() {

	fmt.Println("This is where initializationt starts")
}
func main() {
	x := rand.Intn(250)

	fmt.Printf("%v \n", x)

	switch {

	case x < 100:
		fmt.Println("less than 100")
	case x > 100 && x < 200:
		fmt.Println("less than 200")
	case x > 200 && x < 250:
		fmt.Println("less than 250")
	default:
		fmt.Println("greater than 250")
	}
}
