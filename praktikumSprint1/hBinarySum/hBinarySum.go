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
	num1 := scanner.Text()
	scanner.Scan()
	num2 := scanner.Text()

	if num1 == "0" && num2 == "0" {
		fmt.Print(0)
		return
	}

	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	var res []byte
	var transition uint8
	for len(num1) != 0 {
		var int1, int2 uint8
		int1 = num1[len(num1)-1] - '0'
		if len(num2) != 0 {
			int2 = num2[len(num2)-1] - '0'
		}

		tmp := transition + int1 + int2

		res = append(res, tmp%2+'0')
		transition = tmp / 2

		num1 = num1[:len(num1)-1]
		if len(num2) != 0 {
			num2 = num2[:len(num2)-1]
		}
	}

	if transition != 0 {
		res = append(res, transition+'0')
		transition = 0
	}

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Print(res[i] - '0')
	}
}
