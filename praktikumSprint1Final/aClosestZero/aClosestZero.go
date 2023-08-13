package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const outOfLimit = 1_000_000_001

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	buf := make([]byte, 11_000_000)
	scanner.Buffer(buf, 11_000_000)

	scanner.Scan()
	scanner.Scan()
	values := strings.Split(scanner.Text(), " ")

	slice := make([]int, len(values))
	for idx, value := range values {
		slice[idx], _ = strconv.Atoi(value)
	}

	dist := outOfLimit
	for i := 0; i < len(slice); i++ {
		if slice[i] == 0 {
			dist = 0
		} else {
			dist += 1
		}

		slice[i] = dist
	}

	dist = outOfLimit
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == 0 {
			dist = 0
		} else {
			dist += 1
		}

		if slice[i] > dist {
			slice[i] = dist
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(slice); i++ {
		writer.WriteString(strconv.Itoa(slice[i]))
		writer.WriteByte(' ')
	}
	writer.Flush()
}
