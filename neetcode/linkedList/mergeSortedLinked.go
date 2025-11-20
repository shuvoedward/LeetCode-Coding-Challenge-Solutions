package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	resultHead := result

	cur1 := list1
	cur2 := list2

	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			result.Next = cur1
			cur1 = cur1.Next
		} else {
			result.Next = cur2
			cur2 = cur2.Next
		}
		result = result.Next
	}

	if cur1 != nil {
		result.Next = cur1
	}

	if cur2 != nil {
		result.Next = cur2
	}

	return resultHead.Next
}
