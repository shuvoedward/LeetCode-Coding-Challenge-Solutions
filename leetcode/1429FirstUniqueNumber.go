package leetcode

/*

Problem Description

You need to design a data structure that maintains
a queue of integers and can efficiently identify
the first unique (non-duplicate) integer in the queue.

The FirstUnique class should support three operations:

    Constructor FirstUnique(int[] nums): Initialize the data structure with an initial array of integers.
	These integers form the starting queue.

    Method showFirstUnique(): Return the value of the first unique integer in the queue.
	A unique integer is one that appears exactly once in the entire queue.
	If no unique integer exists, return -1.
	The "first" unique integer refers to the one that appears earliest in the queue order.

    Method add(int value): Add a new integer to the end of the queue.
	This may affect which integers are considered unique.

*/

type FirstUnique struct {
	nodes map[int]*firstUniqueNode
	head  *firstUniqueNode
	tail  *firstUniqueNode
}

type firstUniqueNode struct {
	val  int
	prev *firstUniqueNode
	next *firstUniqueNode
}

func Constructor(nums []int) FirstUnique {
	// build the linked list
	fu := FirstUnique{
		nodes: map[int]*firstUniqueNode{},
		head:  &firstUniqueNode{},
		tail:  &firstUniqueNode{},
	}

	fu.head.next = fu.tail
	fu.tail.prev = fu.head

	for _, num := range nums {
		fu.Add(num)
	}

	return fu

}

func (fu *FirstUnique) showFirstUnique() int {
	if fu.head.next == fu.tail {
		return -1
	}
	return fu.head.next.val
}
func (fu *FirstUnique) Add(value int) {
	// if exists, change linked list, delete from map
	// not exists, just add, map and linked list

	if node, exists := fu.nodes[value]; exists {
		prev := node.prev
		next := node.next
		prev.next = next
		next.prev = prev
		delete(fu.nodes, value)
	} else {
		node := &firstUniqueNode{val: value}
		prev := fu.tail.prev
		prev.next = node
		node.prev = prev
		fu.tail.prev = node
		node.next = fu.tail
		fu.nodes[value] = node
	}

}
