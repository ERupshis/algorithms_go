package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	var res []int

	for i := 2; i*i <= num; i++ {
		for num%i == 0 {
			res = append(res, i)
			num /= i
		}
	}

	if num != 1 {
		res = append(res, num)
	}

	for _, val := range res {
		fmt.Printf("%d ", val)
	}
}
