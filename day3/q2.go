package day3

import (
	"strings"
)

func Q2(input string) int {
	output := 0

	needCalc := true
	for {
		//fmt.Println(input)
		if len(input) == 0 {
			break
		}
		index := strings.Index(input, "mul(")
		notDo := strings.Index(input, "don't()")
		do := strings.Index(input, "do()")
		if index == -1 {
			break
		}
		if notDo == -1 {
			notDo = len(input)
		}
		if do == -1 {
			do = len(input)
		}
		if index > notDo || index > do {
			if notDo > do {
				needCalc = true
				input = input[do+len("do()"):]
				continue
			}
			if do > notDo {
				needCalc = false
				input = input[notDo+len("don't()"):]
				continue
			}
		}
		if !needCalc {
			input = input[index+len("mul("):]
			continue
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
		input = input[index+len(curr)+4+1:]
	}
	return output
}
