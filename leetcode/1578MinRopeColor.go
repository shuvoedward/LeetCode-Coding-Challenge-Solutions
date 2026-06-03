package leetcode

import "fmt"

func minCost(colors string, neededTime []int) int {
	curMax := neededTime[0]
	total := curMax
	result := 0

	l := 0
	for r := 1; r < len(colors); r++ {
		if colors[l] != colors[r] {
			l = r
			result += total - curMax
			total = neededTime[r]
			curMax = total
		} else {
			total += neededTime[r]
			curMax = max(curMax, neededTime[r])
		}
	}

	result += total - curMax

	return result
}

/*

cnt := len(colors)
    sum := 0
    if cnt == 1 {
        return sum
    }
     x := 0

    for i := 1; i < cnt; i++ {
         if colors[x] == colors[i] {
            if neededTime[x] <= neededTime[i] {
                sum = sum + neededTime[x]
                x = i
            } else {
                sum = sum + neededTime[i]
            }

         } else {
            x = i
         }
    }
    return sum


*/

func TestMinCost() {
	tests := []struct {
		colors     string
		neededTime []int
		exp        int
	}{
		{colors: "aabaa", neededTime: []int{1, 2, 3, 4, 1}, exp: 2},
		{colors: "abc", neededTime: []int{1, 2, 3}, exp: 0},
		{colors: "abaac", neededTime: []int{1, 2, 3, 4, 5}, exp: 3},
		{colors: "aaac", neededTime: []int{1, 2, 3, 4}, exp: 3},
		{colors: "aaaaaa", neededTime: []int{1, 2, 3, 4, 5, 6}, exp: 15},
	}

	for _, tt := range tests {
		got := minCost(tt.colors, tt.neededTime)
		fmt.Printf("exp: %d, got: %d\n", tt.exp, got)
	}
}
