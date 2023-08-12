package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const outOfLimit = -274

//func readNumber(s *bufio.Scanner) int {
//	s.Scan()
//	values := strings.Split(s.Text(), " ")
//	res, _ := strconv.Atoi(values[0])
//	return res
//}

func readNumbers(s *bufio.Scanner) []int {
	s.Scan()
	values := strings.Split(s.Text(), " ")

	res := make([]int, 0, len(values))
	for _, value := range values {
		tmp, _ := strconv.Atoi(value)
		res = append(res, tmp)
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var n int
	fmt.Scanf("%d\n", &n)
	//n := readNumber(scanner)

	slice := make([]int, n+2)
	slice[0] = outOfLimit
	slice[n+1] = outOfLimit

	buf := make([]byte, n*5)
	scanner.Buffer(buf, n*5)
	for idx, value := range readNumbers(scanner) {
		slice[idx+1] = value
	}

	res := 0
	for i := 1; i <= n; i++ {
		if slice[i-1] < slice[i] && slice[i] > slice[i+1] {
			res++
		}
	}

	fmt.Println(res)

}
