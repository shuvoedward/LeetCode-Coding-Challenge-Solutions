package leetcode

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	if n >= 1 {
		dp[1] = 1
	}

	for nodes := 2; nodes <= n; nodes++ {
		total := 0
		for root := 1; root <= nodes; root++ {
			left := root - 1
			right := nodes - root
			total += dp[left] * dp[right]
		}
		dp[nodes] = total
	}

	return dp[n]
}
