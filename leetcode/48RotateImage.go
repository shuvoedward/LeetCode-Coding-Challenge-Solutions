package leetcode

import "fmt"

func rotateImage(image [][]int) [][]int {
	l, r := 0, len(image)-1

	for l < r {
		for i := range r - l {
			top, bottom := l, r
			topLeft := image[top][l+i]

			image[top][l+i] = image[bottom-i][l]
			image[bottom-i][l] = image[bottom][r-i]
			image[bottom][r-i] = image[top+i][r]
			image[top+i][r] = topLeft
		}
		l++
		r--
	}

	return image
}

func TestRotateImage() {
	tests := []struct {
		image [][]int
		exp   [][]int
	}{
		{image: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
			exp: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}

	for _, tt := range tests {
		got := rotateImage(tt.image)
		fmt.Printf("exp: %v, got: %v\n", tt.exp, got)
	}
}
