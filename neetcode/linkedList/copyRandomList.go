package linkedlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	oldToCopy := map[*Node]*Node{nil: nil}

	cur := head

	for cur != nil {
		oldToCopy[cur] = cur
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		copy := oldToCopy[cur]
		copy.Next = oldToCopy[cur.Next]
		copy.Random = oldToCopy[cur.Random]
		cur = cur.Next
	}

	return oldToCopy[head]
}
