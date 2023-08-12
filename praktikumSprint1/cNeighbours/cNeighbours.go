package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const outOfLimit = 1001

func readNumber(s *bufio.Scanner) int {
	s.Scan()
	values := strings.Split(s.Text(), " ")
	num, _ := strconv.Atoi(values[0])
	return num
}

func readNumbers(s *bufio.Scanner) []int {
	s.Scan()
	values := strings.Split(s.Text(), " ")

	var res = make([]int, len(values))
	for idx, val := range values {
		res[idx], _ = strconv.Atoi(val)
	}
	return res
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	row := readNumber(scanner)
	col := readNumber(scanner)

	matrix := make([][]int, row+2)
	for i := 0; i <= row+1; i++ {
		matrix[i] = make([]int, col+2)
		for j := 0; j <= col+1; j++ {
			matrix[i][j] = outOfLimit
		}
	}

	for i := 1; i <= row; i++ {
		values := readNumbers(scanner)

		for j := 1; j <= len(values); j++ {
			matrix[i][j] = values[j-1]
		}
	}

	y := readNumber(scanner)
	x := readNumber(scanner)

	yShift := []int{-1, 0, 1, 0}
	xShift := []int{0, -1, 0, 1}

	var res []int
	for i := 0; i < len(xShift); i++ {
		tmp := matrix[y+1+yShift[i]][x+1+xShift[i]]
		if tmp != outOfLimit {
			res = append(res, tmp)
		}
	}

	sort.Ints(res)
	for _, value := range res {
		fmt.Printf("%d ", value)
	}
}
