package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	subTrees := map[string][]*TreeNode{}
	result := []*TreeNode{}

	var dfs func(root *TreeNode) string

	dfs = func(root *TreeNode) string {
		if root == nil {
			return "nil"
		}

		s := strings.Join([]string{strconv.Itoa(root.Val), dfs(root.Left), dfs(root.Right)}, ",")
		if node, exists := subTrees[s]; exists && len(node) == 1 {
			result = append(result, node[0])
		}

		subTrees[s] = append(subTrees[s], root)
		return s
	}

	dfs(root)
	return result
}

func findDuplicateSubtrees2(root *TreeNode) []*TreeNode {
	tripletToID := make(map[string]int)
	idCounts := make(map[int]int)
	res := make([]*TreeNode, 0)

	var tripletGen func(*TreeNode) int
	tripletGen = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftNodeID := tripletGen(node.Left)
		rightNodeID := tripletGen(node.Right)

		triplet := fmt.Sprintf("%v,%v,%v", node.Val, leftNodeID, rightNodeID)
		if _, ok := tripletToID[triplet]; !ok {
			tripletToID[triplet] = len(tripletToID) + 1
		}

		id := tripletToID[triplet]
		idCounts[id]++
		if idCounts[id] == 2 {
			res = append(res, node)
		}

		return id
	}

	tripletGen(root)
	return res
}
