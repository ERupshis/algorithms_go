package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var n int
	fmt.Scanf("%d \n", &n)
	scannerBuf := make([]byte, n)
	scanner.Buffer(scannerBuf, n)

	scanner.Scan()
	scanner.Scan()
	words := strings.Split(string(scanner.Text()), " ")

	var longestWord string
	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}

	fmt.Printf("%s\n%d", longestWord, len(longestWord))
}
