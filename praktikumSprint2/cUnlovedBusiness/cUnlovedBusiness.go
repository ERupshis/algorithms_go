package main

/*
Comment it before submitting
type ListNode struct {
    data   string
    next *ListNode
}
*/

func getNodeByIndex(head *ListNode, idx int) *ListNode {
	res := head
	for res != nil && idx > 0 {
		res = res.next
		idx--
	}

	return res
}

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		return head.next
	}

	beforeRemove := getNodeByIndex(head, idx-1)
	toRemove := beforeRemove.next

	beforeRemove.next = toRemove.next
	return head
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*newHead :=*/ Solution(&node0, 1)
	// result is : node0 -> node2 -> node3
}
