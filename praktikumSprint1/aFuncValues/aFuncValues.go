package main

import "fmt"

func main() {
	var a, x, b, c int
	fmt.Scan(&a, &x, &b, &c)

	fmt.Println(a*x*x + b*x + c)

	return
}
