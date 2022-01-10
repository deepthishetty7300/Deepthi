package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	var slice []int
	for num > 0 {
		slice = append(slice, num%10)
		num = num / 10
	}
	length := len(slice) - 1
	for i := length; i >= 0; i-- {

		fmt.Print(slice[i], " ")
	}
}
