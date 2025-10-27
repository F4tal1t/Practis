package main

import "fmt"

func divide(a, b int) int {
	defer func() {
		if resp := recover(); resp != nil {
			fmt.Printf("Recovered from panic: %v\n", resp)
		}
	}()
	if b == 0 {
		panic("Division by zero")
	}
	return a / b
}
func main() {

	fmt.Println(divide(4, 2))
	fmt.Println(divide(4, 0))
	fmt.Println(divide(4, 4))
}
