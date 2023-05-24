package main

import (
	"fmt"
	"log"
)

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%v is ODD\n", i)
		log.Println("asdfp")
	}
}
