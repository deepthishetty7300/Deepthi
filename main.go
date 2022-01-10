package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)

	s := strconv.Itoa(num)
	firstCharacter := s[0:1]

	i, _ := strconv.Atoi(firstCharacter)

	if i == 0 {
		fmt.Println("Enter a Valid number")
		main()

	} else {

		var s []int
		for num > 0 {
			s = append(s, num%10)
			num = num / 10
		}
		length := len(s) - 1
		for i := length; i > -1; i-- {
			fmt.Print(s[i], " ")
		}
	}

}
