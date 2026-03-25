package leetcode

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([][]rune, numRows)
	curRow := 0
	goingDown := false

	for _, ch := range s {
		rows[curRow] = append(rows[curRow], ch)

		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	result := make([]rune, 0, len(s))
	for _, row := range rows {
		result = append(result, row...)
	}

	return string(result)
}

func TestZigZagConvert() {
	tests := []struct {
		s       string
		numRows int
		exp     string
	}{
		{s: "PAYPALISHIRING", numRows: 4, exp: "PINALSIGYAHRPI"},
		{s: "PAYPALISHIRING", numRows: 3, exp: "PAHNAPLSIIGYIR"},
		{s: "A", numRows: 1, exp: "A"},
	}

	for _, tt := range tests {
		got := convert(tt.s, tt.numRows)
		fmt.Printf("exp: %s, got: %s, pass: %t\n", tt.exp, got, tt.exp == got)
	}
}
