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
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	matrix := make([][]int16, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		strValues := strings.Split(scanner.Text(), " ")
		matrix[i] = make([]int16, m)

		for j := 0; j < m; j++ {
			num, _ := strconv.Atoi(strValues[j])
			matrix[i][j] = int16(num)
		}
	}

	res := make([][]int16, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int16, n)
		for j := 0; j < n; j++ {
			res[i][j] = matrix[j][i]
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			strVal := strconv.Itoa(int(res[i][j]))
			writer.WriteString(strVal)
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
		writer.Flush()
	}

}
