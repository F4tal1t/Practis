package main

import "fmt"

func main() {
	//for loop
	/*
		fmt.Println("Basic For Loop")
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}

		//for as while loop
		x := 5
		for x < 10 {
			fmt.Println(x)
			x++
		}
		//range for loop
		num := []int{1, 2, 3, 4, 5}
		for idx, val := range num {
			fmt.Println(idx, val)
		}
	*/
	//switch case
	day := 7
	switch day {
	case 1, 2, 3:
		fmt.Println("Week Start")
	case 4, 5, 6:
		fmt.Println("Week End")
	default:
		fmt.Println("Week Enddd")
	}
	var x interface{} = "am"
	switch x.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("Dony knoow")
	}
}
