package main

import "fmt"

type Car interface {
	StartEngine() string
	Drive() string
}

type Sedan struct {
	Brand  string
	Length int
}

type SUV struct {
	Brand  string
	Length int
}

func (s Sedan) StartEngine() string {
	return fmt.Sprintf("Enginer Momen:%s", s.Brand)
}

func (s Sedan) Drive() string {
	return fmt.Sprintf("Enginer Lengh:%d", s.Length)
}

func (s SUV) StartEngine() string {
	return fmt.Sprintf("Enginer Momen:%s", s.Brand)
}

func (s SUV) Drive() string {
	return fmt.Sprintf("Enginer Lengh:%d", s.Length)
}

func DefineCar(c Car) {

	//fmt.Println(c.StartEngine())
	//fmt.Println(c.Drive())
}

func main() {
	/*var car Car
	car = Sedan{Brand: "Momen", Length: 5}
	DefineCar(car)
	car = SUV{Brand: "MVv", Length: 12}
	DefineCar(car)

	*/
	//	var output interface{}
	//	output = 45
	//	Y, ok := output.(int)
	//	if !ok {
	//		fmt.Println("Unable to convert output to int")
	//	}
	//	fmt.Println(Y)
	//	output = "GFG"
	//	str, ok := output.(int)
	//	if !ok {
	//		fmt.Println("Unable to convert output to string")
	//	}
	//	fmt.Println(str)
	//
	
}
