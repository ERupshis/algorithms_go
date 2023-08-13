package main

import "fmt"

const base = 4

func main() {
	var num int
	fmt.Scan(&num)

	for num > 1 {
		if num%base != 0 {
			fmt.Println("False")
			break
		}

		num /= base
	}

	if num == 1 {
		fmt.Println("True")
	}
}
