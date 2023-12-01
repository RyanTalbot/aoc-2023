package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var sumArray []int
	var sum int

	input, _ := os.ReadFile("./day-01/input.txt")
	lines := strings.Fields(string(input))

	for line := range lines {
		reg := regexp.MustCompile(`[^0-9]+`).ReplaceAllString(lines[line], "")
		sumArray = append(sumArray, sumFirstAndLast(reg))
	}

	for number := range sumArray {
		sum += sumArray[number]
	}

	fmt.Println(sumArray)
	fmt.Println(sum)

}

func sumFirstAndLast(str string) int {
	var first, last int

	if len(str) > 0 {
		if len(str) == 1 {
			first = parse(str, 0)
			last = parse(str, 0)
		} else {
			first = parse(str, 0)
			last = parse(str, len(str)-1)
		}

	}

	merge := fmt.Sprintf("%d%d", first, last)
	result, _ := strconv.Atoi(merge)

	return result
}

func parse(str string, pos int) int {
	element := str[pos]
	stringified := string(element)
	result, _ := strconv.Atoi(stringified)

	return result
}
