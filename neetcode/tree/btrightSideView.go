package tree

func RightSideViewBFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		rightSide := 0
		qLen := len(q)

		for range qLen {
			node := q[0]
			q = q[1:]

			if node != nil {
				rightSide = node.Val
				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}
			}
		}
		res = append(res, rightSide)
	}
	return res
}

func RightSideViewDFS(root *TreeNode) []int {
	var res []int

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}

	dfs(root, 0)
	return res
}

/*
if depth == len(res) {
			res = append(res, node.Val)
		}
This line ensures that only the first node visited at
each depth is added to the result.
Since the traversal is right-first depth-first search (DFS),
the first node visited at any given depth is the rightmost
node at that level.

Here's how to say it more clearly:

"If the current depth equals the length of the result list,
it means we haven't added any node from this depth level yet.
Since we're traversing the right child before the left child,
the first node we see at each depth is the rightmost node at
that level. So, we append its value to the result."

This trick relies on the fact that in DFS, the recursive
depth increases as we go down the tree, and the list res
only grows one element per level — only once per depth.

Let me know if you want a visual example to solidify it.
*/

// My solution
// In the queue, the last element for each level is the right side node
func RightSideView(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		l := len(queue)
		for i := range l {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if i == l-1 {
				res = append(res, node.Val)
			}
		}
	}

	return res
}
