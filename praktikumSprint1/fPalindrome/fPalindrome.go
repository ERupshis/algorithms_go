package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	runes := []rune(line)

	l := 0
	r := len(runes) - 1
	for {
		for !unicode.IsLetter(runes[l]) && !unicode.IsDigit(runes[l]) {
			l++
		}

		for !unicode.IsLetter(runes[r]) && !unicode.IsDigit(runes[r]) {
			r--
		}

		if l >= r {
			break
		}

		if unicode.ToLower(runes[l]) == unicode.ToLower(runes[r]) {
			l++
			r--
		} else {
			fmt.Println("False")
			break
		}
	}

	if l >= r {
		fmt.Println("True")
	}

}
