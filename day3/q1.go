package day3

import (
	"strconv"
	"strings"
)

func Q1(input string) int {
	output := 0

	for {
		if len(input) == 0 {
			break
		}
		index := strings.Index(input, "mul(")
		if index == -1 {
			break
		}
		curr := ""

		for i := index + 4; i < len(input); i++ {
			b := input[i]
			if !valid(b) {
				curr = ""
				break
			}
			if input[i] == ')' {
				break
			}
			curr = curr + string(b)
		}

		output += calc(curr)
		input = input[index+len(curr)+1:]
	}
	return output
}
func calc(input string) int {
	if len(input) == 0 {
		return 0
	}

	split := strings.Split(input, ",")
	if len(split) != 2 {
		return 0
	}
	num1, _ := strconv.Atoi(split[0])
	num2, _ := strconv.Atoi(split[1])
	return num1 * num2
}

func valid(s byte) bool {
	if s >= '0' && s <= '9' {
		return true
	}
	if s == ')' {
		return true
	}
	if s == ',' {
		return true
	}
	return false

}
