package main

import "fmt"

func findDigitSum(num int) int {

	sum := 0
	rem := 0
	for num > 0 {
		rem = num % 10
		sum = sum + rem
		num = num / 10
	}
	if sum > 9 {
		return findDigitSum(sum)
	}
	return sum
}
func main() {

	fmt.Println(findDigitSum(0122))

}
