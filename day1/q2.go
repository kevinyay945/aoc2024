package main
import (
"bufio"
"fmt"
"os"
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

	scoreMap := map[int]int{}
	for _, p := range rightPoint {
		scoreMap[p] = scoreMap[p] + 1
	}

	output := 0
	for _, p := range leftPoint {
		output = output + scoreMap[p]*p
	}

	fmt.Println(output)

}

