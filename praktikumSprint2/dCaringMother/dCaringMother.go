package main

/*
Comment it before submitting
type ListNode struct {
    data   string
    next *ListNode
}
*/

func Solution(head *ListNode, elem string) int {
	res := -1

	idx := 0
	for head != nil {
		if head.data == elem {
			res = idx
			break
		}

		idx++
		head = head.next
	}

	return res
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*idx :=*/ Solution(&node0, "node2")
	// result is : idx == 2
}
