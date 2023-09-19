package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val  int
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Put(num int) {
	newNode := &Node{val: num, next: nil}
	if q.size == 0 {
		q.head = newNode
	} else {
		q.tail.next = newNode
	}

	q.tail = newNode
	q.size++
}

func (q *Queue) Get() {
	if q.size == 0 {
		fmt.Println("error")
	} else {
		fmt.Println(q.head.val)
		q.head = q.head.next
		q.size--
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

	queue := Queue{}

	for n > 0 {
		s.Scan()
		values := strings.Split(s.Text(), " ")

		switch values[0] {
		case "put":
			num, _ := strconv.Atoi(values[1])
			queue.Put(num)
		case "get":
			queue.Get()
		case "size":
			queue.Size()
		}

		n--
	}
}
