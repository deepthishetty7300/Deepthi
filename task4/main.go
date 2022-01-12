package main

import (
	"fmt"
	"log"
)

func myFunc(a interface{}) {
	var b string
	fmt.Printf("%T\n", a)
	switch a.(type) {
	case int:
		log.Println("enter a string")
		return
	case string:
		b = a.(string)
	default:
		fmt.Println("\nType not found")
	}
	fmt.Println("value of b is:", b)
}
func main() {
	var a interface {
	} = 897

	myFunc(a)
}
