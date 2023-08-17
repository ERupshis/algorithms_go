package main

import "fmt"

type Stack struct {
	data []byte
}

func (s *Stack) push(x byte) {
	s.data = append(s.data, x)
}

func (s *Stack) top() byte {
	return s.data[len(s.data)-1]
}

func (s *Stack) pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) empty() bool {
	return len(s.data) == 0
}

func isCorrectPair(left, right byte) bool {
	if left == '(' {
		return left+1 == right
	} else {
		return left+2 == right
	}
}

func main() {
	var seq string
	fmt.Scan(&seq)

	stack := Stack{}

	correct := true
	for _, ch := range []byte(seq) {
		if ch == '[' || ch == '{' || ch == '(' {
			stack.push(ch)
			continue
		} else if stack.empty() || !isCorrectPair(stack.top(), ch) {
			correct = false
			break
		}
		stack.pop()
	}

	if stack.empty() && correct {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}
