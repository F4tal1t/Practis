package main

import (
	"fmt"
	"strconv"
)

func Hello(a int, b int) int {
	fmt.Printf("Hello World %v\n", a+b)
	return a + b
}

// Variadic Functions
func Add(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

// Functions with Multiple Results
func swap(x, y string) (string, string) {
	return y, x
}
func Convert(str string) (int, error) {
	return strconv.Atoi(str)
}
func Area(a, b float64) (area float64, per float64) {
	area = a * b
	per = 2 * (b + a)
	return
}
func main() {
	ans := Hello(4, 5)
	fmt.Println(ans)
	println(Add(1, 2, 6, 7, 7, 8))
	println(swap("hello", "world"))

	fmt.Println(Area(4.3, 5.6))
}
