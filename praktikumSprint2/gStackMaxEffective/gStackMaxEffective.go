package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	data []int
	max  []int
}

func (s *Stack) push(x int) {
	s.data = append(s.data, x)

	var maxValue int
	if len(s.max) == 0 {
		maxValue = x
	} else if x > s.max[len(s.max)-1] {
		maxValue = x
	} else {
		maxValue = s.max[len(s.max)-1]
	}

	s.max = append(s.max, maxValue)
}

func (s *Stack) pop() {
	if len(s.data) == 0 {
		fmt.Println("error")
	} else {
		s.data = s.data[:len(s.data)-1]
		s.max = s.max[:len(s.max)-1]
	}
}

func (s *Stack) getMax() {
	if len(s.data) == 0 {
		fmt.Println("None")
		return
	}

	fmt.Println(s.max[len(s.max)-1])
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	stack := Stack{}
	for n != 0 {
		scanner.Scan()
		values := strings.Split(scanner.Text(), " ")

		switch values[0] {
		case "push":
			num, _ := strconv.Atoi(values[1])
			stack.push(num)
		case "pop":
			stack.pop()
		case "get_max":
			stack.getMax()
		}

		n--
	}
}
