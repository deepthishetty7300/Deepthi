package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	var s string
	var num int
	for {
		_, err := fmt.Scan(&s)
		num, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Enter a valid number")
		} else {
			var slice []int
			str := []string{}

			for num > 0 {
				slice = append(slice, num%10)
				num = num / 10
			}
			length := len(slice) - 1
			for i := length; i >= 0; i-- {

				fmt.Print(slice[i], " ")
			}
			fmt.Println(" ", reflect.TypeOf(slice))
			for i := range slice {
				number := slice[i]
				text := strconv.Itoa(number)
				str = append(str, text)
			}
			result := strings.Join(str, ",")
			fmt.Print(result)
			fmt.Println(" ", reflect.TypeOf(result))
			break
		}
	}
}
