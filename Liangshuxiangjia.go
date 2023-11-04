package main

import "fmt"

func main() {
	num1 := 1
	num2 := 2
	sum := add(num1, num2)
	fmt.Println("和为: ", sum)
}

func add(num1, num2 int) int {
	return num1 + num2
}
