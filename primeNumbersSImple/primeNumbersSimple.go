package main

import "fmt"

func isPrime(n int) bool {
	i := 2
	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i++
	}

	return true
}

func primeNumbers(n int) []int {
	var res []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Printf("%v", primeNumbers(n))
}
