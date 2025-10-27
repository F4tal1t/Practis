package main

import (
	"errors"
	"fmt"
)

func check(num int) (string, error) {
	if num < 0 {
		return "", ErrInvalidFile
	}
	return "Number is positve", nil

}

var ErrInvalidFile = errors.New("Invalid file")

func read(filename string) error {
	if filename == "" {
		return errors.New("File is empty")
	}
	return nil
}
func file() (string, error) {
	e := read("")
	if e != nil {
		return "", fmt.Errorf("File read error:%w", e)
	}
	return "File exist", nil
}
func main() {
	_, e := check(-2)
	if e != nil {
		fmt.Println("Error:", e)
	}
	_, er := file()
	if er != nil {
		fmt.Println(er)
		err := errors.Unwrap(er)
		if errors.Is(err, ErrInvalidFile) {
			fmt.Println(err)
		}

	}

}
