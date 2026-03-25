package leetcode

import "fmt"

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
	}

	return dummy.Next
}

func BuildLinkedList(list []int) *ListNode {
	head := &ListNode{}
	cur := head
	for i := 0; i < len(list); i++ {
		cur.Val = list[i]
		if i < len(list)-1 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}

	return head
}

func PrintLinkedList(head *ListNode, title string) {
	fmt.Printf("%s\n", title)
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}

func RunswapPairsTest() {
	tests := []struct {
		head *ListNode
		exp  *ListNode
	}{
		{head: BuildLinkedList([]int{1, 2, 3, 4}), exp: BuildLinkedList([]int{2, 1, 4, 3})},
	}

	for i, tt := range tests {
		got := swapPairs(tt.head)
		fmt.Printf("\ntest no %d\n", i+1)
		PrintLinkedList(tt.exp, "exp")
		fmt.Println()
		PrintLinkedList(got, "got")
	}

}
