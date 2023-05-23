package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	switch {
	case x < 4 && y < 4:
		fmt.Println("both are less than four")
	case x > 6 && y > 6:
		fmt.Println("both are greater than six")
	case x >= 4 && y <= 4:
		fmt.Println("both are inclusive 4 and 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
	}
}
