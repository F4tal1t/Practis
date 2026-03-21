package main

import (
	"fmt"
)

func demo(i int) {
	fmt.Println("hell", i)
}

func main() {

	fmt.Println("Balls")
	defer demo(3)
	defer demo(1)

	demo(2)

	//go to is skipped
}
