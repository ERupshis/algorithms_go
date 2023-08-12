package main

import "fmt"

// time: O(n)
func linearSieve(n int) []int {
	leastPrimes := make([]int, n+1)
	var primes []int

	for i := 2; i <= n; i++ {
		if leastPrimes[i] == 0 {
			leastPrimes[i] = i
			primes = append(primes, i)
		}

		for _, prime := range primes {
			x := prime * i

			if prime > leastPrimes[i] || x > n {
				break
			}

			leastPrimes[x] = prime
		}
	}

	return primes
}

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Printf("%v", linearSieve(num))
}
