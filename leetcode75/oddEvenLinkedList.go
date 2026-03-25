package leetcode75

func OddEvenList(head *ListNode) *ListNode {
	// need atleast three element, if not already in correct position
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even

	// even pointer reaches end faster than odd pointer
	// for even list reaches nil, for odd len of list, reaches the last node if both pointer start at beginning
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}

/*
Need atleast 3 elements

even goes to end faster than odd pointer, so for loop condition checks for even

*/
