package linkedlist

func HasCycle(head *ListNode) bool {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}
