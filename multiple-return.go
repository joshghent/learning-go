package main

import (
	"fmt"
)

func vals(num1 int, num2 int) (int, int) {
	return num1 * 2, num2 * 3
}

func main() {
	a, b := vals(2, 3)

	fmt.Println(a)
	fmt.Println(b)
}
