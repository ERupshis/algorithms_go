package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	word1 := scanner.Text()

	scanner.Scan()
	word2 := scanner.Text()

	mapWord1 := make(map[byte]int)
	for _, ch := range []byte(word1) {
		mapWord1[ch]++
	}

	mapWord2 := make(map[byte]int)
	for _, ch := range []byte(word2) {
		mapWord2[ch]++
	}

	for k, v := range mapWord2 {
		if v != mapWord1[k] {
			fmt.Print(string(k))
			break
		}
	}
}
