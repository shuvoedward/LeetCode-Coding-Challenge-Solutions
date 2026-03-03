package linkedlist

// Two pointers
func RemomveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	left := dummy
	right := head

	for n > 0 {
		// right == the nth node
		right = right.Next
		n--
		// n reaches 0
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return dummy.Next
}
