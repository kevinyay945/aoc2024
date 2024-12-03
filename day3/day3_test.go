package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestQ1(t *testing.T) {
	input := readfile("./input.txt")
	allString := strings.Join(input, "")
	q1Sum := Q1(allString)

	fmt.Println("Q1", q1Sum)
}
func TestQ2(t *testing.T) {
	input := readfile("./input.txt")
	allString := strings.Join(input, "")
	sum := Q2(allString)

	fmt.Println("Q2", sum)
}

func readfile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output []string
	for scanner.Scan() {
		txt := scanner.Text()
		output = append(output, txt)
	}
	return output
}
