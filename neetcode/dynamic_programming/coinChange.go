package dynamicprogramming

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	seenAmount := map[int]int{}

	var dfs func(amount int) int
	dfs = func(amount int) int {
		if amount < 0 {
			return -1
		}

		if amount == 0 {
			return 0
		}

		if stepCount, exists := seenAmount[amount]; exists {
			return stepCount
		}

		step := math.MaxInt

		for _, coin := range coins {
			res := dfs(amount - coin)

			if res == -1 {
				continue
			}

			step = min(step, res+1)
		}

		if step == math.MaxInt {
			seenAmount[amount] = -1
			return -1
		}

		seenAmount[amount] = step
		return step
	}

	return dfs(amount)
}

func Test_CoinChange() {
	tests := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{coins: []int{1, 4, 3}, amount: 7, expected: 2},
		{coins: []int{1, 5, 10}, amount: 12, expected: 3},
		{coins: []int{2}, amount: 3, expected: -1},
	}

	for i, tt := range tests {
		output := coinChange(tt.coins, tt.amount)
		fmt.Printf("test: %d, expected: %d, got: %d\n", i, tt.expected, output)
	}
}
