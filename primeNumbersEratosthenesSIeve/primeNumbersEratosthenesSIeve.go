package main

import "fmt"

func eratosthenesSieve(n int) []int {
	if n <= 1 {
		return nil
	}

	var primes = make([]bool, n+1)
	for idx, _ := range primes {
		primes[idx] = true
	}

	primes[0] = false
	primes[1] = false

	for num := 2; num < n; num++ {
		if primes[num] {
			for elem := num * num; elem <= n; elem += num {
				primes[elem] = false
			}
		}
	}

	var res []int
	for idx, _ := range primes {
		if primes[idx] {
			res = append(res, idx)
		}
	}

	return res
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Printf("%v", eratosthenesSieve(n))
}
