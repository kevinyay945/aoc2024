package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestQ1(t *testing.T) {
	file, err := os.Open("q1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]int
	for scanner.Scan() {
		txt := scanner.Text()
		split := strings.Split(txt, " ")
		//fmt.Println(split[1])
		report := make([]int, len(split))
		for i, s := range split {
			report[i], _ = strconv.Atoi(s)
		}
		input = append(input, report)
	}

	fmt.Println("Q1", Q1(input))

	fmt.Println("Q2", Q2(input))
}
