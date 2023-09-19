package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	data []int
	head int
	tail int
	size int
}

func CreateQueue(size int) *Queue {
	data := make([]int, size)
	return &Queue{data: data}
}

func (q *Queue) Push(num int) {
	if cap(q.data) == q.size {
		fmt.Println("error")
	} else {
		q.data[q.tail] = num
		q.tail = (q.tail + 1) % len(q.data)
		q.size++
	}
}

func (q *Queue) Pop() {
	if q.size == 0 {
		fmt.Println("None")
	} else {
		fmt.Println(q.data[q.head])
		q.head = (q.head + 1) % len(q.data)
		q.size--
	}
}

func (q *Queue) Peek() {
	if q.size == 0 {
		fmt.Println("None")
	} else {
		fmt.Println(q.data[q.head])
	}
}

func (q *Queue) Size() {
	fmt.Println(q.size)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Scan()
	size, _ := strconv.Atoi(s.Text())

	queue := CreateQueue(size)

	for n > 0 {
		s.Scan()
		values := strings.Split(s.Text(), " ")

		switch values[0] {
		case "push":
			num, _ := strconv.Atoi(values[1])
			queue.Push(num)
		case "pop":
			queue.Pop()
		case "peek":
			queue.Peek()
		case "size":
			queue.Size()
		}

		n--
	}
}
