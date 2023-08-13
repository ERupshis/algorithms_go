package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if num == 0 {
		fmt.Print(num)
		return
	}

	var res []int
	for num != 0 {
		res = append(res, num%2)
		num /= 2
	}

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i])
	}
}
