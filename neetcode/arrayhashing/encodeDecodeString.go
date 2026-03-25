package arrayhashing

import (
	"strconv"
	"strings"
)

type Solution struct{}

/*
Only a pointer (s *Solution) receiver can modify the struct's fields directly.
For example, if the Solution struct had fields,
a pointer receiver could change their values.

without pointer receiver (s Solution) would create a copy of it and would not
make any changes to the origin Solution
*/
func (s *Solution) Encode(strs []string) string {
	/*
	   In Go, strconv.Itoa is a function from the strconv package that converts an integer to its string representation.

	   The name Itoa stands for Integer to ASCII:

	   Itoa: Converts an integer (int) to a string.
	   The result is the string representation of the number.
	   import "strconv"

	   func main() {
	       num := 42
	       str := strconv.Itoa(num) // Convert integer 42 to string "42"
	       fmt.Println(str)         // Output: "42"
	   }


	*/

	res := ""
	for _, str := range strs {
		res += strconv.Itoa(len(str)) + "#" + str

	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	// encoded = "4#abcd2#gh3#ijk" , 4##abc 4#a4bc
	// [abcd, gh, ijk]
	res := []string{}

	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++

		}
		length, _ := strconv.Atoi(encoded[i:j]) // second return value is error

		i = j + 1
		res = append(res, encoded[i:i+length])
		i += length

	}
	return res

}
func encode(strs []string) string {
	// 5#Hello5#World#2My
	// s[2 : 5+1]
	if len(strs) == 0 {
		return ""
	}

	var encodeStr strings.Builder

	for _, str := range strs {
		l := len(str)
		encodeStr.WriteString(strconv.Itoa(l))
		encodeStr.WriteString("#")
		encodeStr.WriteString(str)
	}

	return encodeStr.String()
}

func decode(encoded string) []string {
	if len(encoded) == 0 {
		return []string{}
	}

	result := []string{}
	i := 0
	for i < len(encoded) {
		s := string(encoded[i])
		l, _ := strconv.Atoi(s)
		result = append(result, encoded[i+2:i+l+2])
		i = i + l + 2
	}

	return result
}
