package leetcode

import (
	"fmt"
	"slices"
)

func bagOfTokensScore(tokens []int, power int) int {
	slices.Sort(tokens)

	left, right := 0, len(tokens)-1
	score := 0

	for left <= right {
		if power < tokens[left] && score == 0 {
			return score
		}

		if power >= tokens[left] {
			score++
			power -= tokens[left]
			left++
		} else if score > 0 {
			if left < right {
				score--
				power += tokens[right]
			}
			right--
		}
	}

	return score
}

func bagOfTokensScoreOpt(tokens []int, power int) int {
	score, maxScore := 0, 0
	slices.Sort(tokens)
	left, right := 0, len(tokens)-1
	for left <= right {
		if power >= tokens[left] {
			power -= tokens[left]
			score++
			maxScore = max(score, maxScore)
		} else if score > 0 {
			score--
			power += tokens[right]
			right--
		} else {
			break
		}
	}

	return maxScore
}

func TestBagOfTokens() {
	tests := []struct {
		tokens []int
		power  int
		exp    int
	}{
		{tokens: []int{68, 85, 34, 25, 60}, power: 44, exp: 2},
	}

	for _, tt := range tests {
		got := bagOfTokensScore(tt.tokens, tt.power)
		fmt.Printf("exp: %d, got: %d\n", tt.exp, got)
	}
}
