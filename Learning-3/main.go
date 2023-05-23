package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)

	fmt.Printf("%v \n", x)

	if x <= 100 {
		fmt.Printf("Value is %v of 0 - 100 \n", x)
	} else if x <= 200 {
		fmt.Printf("Value is %v of 101 - 200 \n", x)
	} else {
		fmt.Printf("Value is %v of 201 - 250 \n", x)
	}

	y := rand.Intn(3)
	fmt.Printf("%v is random int of 3\n", y)
}
