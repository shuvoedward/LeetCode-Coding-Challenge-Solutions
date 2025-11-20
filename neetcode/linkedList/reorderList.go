package linkedlist

func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	second := slow.Next
	slow.Next = nil
	// reverse the secondlist
	var prev *ListNode
	for second != nil {
		temp := second.Next
		second.Next = prev
		prev = second
		second = temp
	}

	second = prev

	first := head
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}
