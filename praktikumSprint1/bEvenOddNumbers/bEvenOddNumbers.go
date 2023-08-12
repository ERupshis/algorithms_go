package main

import (
	"fmt"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	a = abs(a)
	b = abs(b)
	c = abs(c)

	if (a%2 == 0) && (b%2 == 0) && (c%2 == 0) || (a%2 == 1) && (b%2 == 1) && (c%2 == 1) {
		fmt.Println("WIN")
	} else {
		fmt.Println("FAIL")
	}
}
