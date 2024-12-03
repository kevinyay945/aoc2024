package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var leftPoint []int
	var rightPoint []int
	for scanner.Scan() {
		txt := scanner.Text()
		split := strings.Split(txt, "   ")
		//fmt.Println(split[1])
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])
		leftPoint = append(leftPoint, left)
		rightPoint = append(rightPoint, right)
	}
	sort.Ints(leftPoint)
	sort.Ints(rightPoint)

	output := 0
	for i := range leftPoint {
		r := rightPoint[i]
		l := leftPoint[i]
		if r > l {
			output = output + r - l
		} else {
			output = output + l - r
		}
	}

	fmt.Println(output)

}
