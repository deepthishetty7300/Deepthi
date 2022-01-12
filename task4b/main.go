package main

import (
	"errors"
	"fmt"
)

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by 0!")
	}
	return x / y, nil
}

func main() {
	answer, err := divide(5, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}
}
