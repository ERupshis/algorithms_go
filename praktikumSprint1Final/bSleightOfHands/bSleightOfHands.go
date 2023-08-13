package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	uniqueButtons := make(map[byte]int)
	for i := 0; i < 4; i++ {
		scanner.Scan()
		for _, elem := range scanner.Text() {
			uniqueButtons[byte(elem)]++
		}
	}

	delete(uniqueButtons, '.')

	k *= 2
	var res int
	for _, v := range uniqueButtons {
		if v <= k {
			res++
		}
	}

	fmt.Println(res)
}
