package main

import "fmt"

func main() {
	a := 1

	if a == 5 {
		fmt.Println("Is 5")
		return
	} else {
		fmt.Println("Not 5")
	}
}
