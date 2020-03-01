package main

import "fmt"

func calc(num1, num2 int) int {
	var result = num1 + num2
	return result
}

func main() {
	var num1 int = 12
	var num2 int = 3
	fmt.Println(calc(num1,num2))
}