package leetcode75

func longestZigZag(root *TreeNode) int {
	maxLength := 0

	var dfs func(root *TreeNode, goLeft bool, length int)

	dfs = func(root *TreeNode, goLeft bool, length int) {
		if root == nil {
			return
		}

		if length > maxLength {
			maxLength = length
		}

		if goLeft {
			dfs(root.Left, false, length+1)
			dfs(root.Right, true, 1)
		} else {
			dfs(root.Right, true, length+1)
			dfs(root.Left, false, 1)
		}
	}

	dfs(root, true, 0)
	dfs(root, false, 0)

	return maxLength
}
