package leetcode75

import (
	"strconv"
	"strings"
)

func DecodeString(s string) string {
	// 3[a2[c]]
	// accaccacc
	// 3[a]2[bc]
	// aaabcbc

	numStack := []int{}
	strStack := []string{}

	currentNum := 0
	currentString := ""

	for _, char := range s {
		if char >= '0' && char <= '9' {
			digit, _ := strconv.Atoi(string(char))
			currentNum = currentNum*10 + digit
		} else if char == '[' {
			numStack = append(numStack, currentNum)
			strStack = append(strStack, currentString)

			currentNum = 0
			currentString = ""
		} else if char == ']' {
			k := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevString := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			var builder strings.Builder
			builder.WriteString(prevString)

			for range k {
				builder.WriteString(currentString)
			}

			currentString = builder.String()

		} else {
			currentString += string(char)
		}
	}

	return currentString
}
