package main

import "fmt"

func findDigitSum(num int) int {
	}
	if sum > 9 {
		return findDigitSum(sum)
	}
	return sum
}
func main() {

	fmt.Println(findDigitSum(2462))

}
