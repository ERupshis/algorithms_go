package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	//read X

	scanner.Scan()
	firstStr := strings.Split(scanner.Text(), " ")
	var first []int

	for _, val := range firstStr {
		tmp, _ := strconv.Atoi(val)
		first = append(first, tmp)
	}

	scanner.Scan()
	second, _ := strconv.Atoi(scanner.Text())

	var res []int
	for i := len(first) - 1; i >= 0; i-- {
		sum := first[i] + second
		res = append(res, sum%10)

		second = sum / 10
	}

	for second != 0 {
		res = append(res, second%10)
		second /= 10
	}

	writer := bufio.NewWriter(os.Stdout)
	for i := len(res) - 1; i >= 0; i-- {
		writer.WriteByte(byte('0' + res[i]))
		writer.WriteByte(' ')
	}
	writer.Flush()
}
