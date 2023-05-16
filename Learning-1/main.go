package main

import (
	"fmt"

	addon "github.com/cwilliamson29/Golang_Learning_Addon"
)

func main() {
	s1 := addon.Bark()
	s2 := addon.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(addon.Bark())
	fmt.Println(addon.Barks())

}
