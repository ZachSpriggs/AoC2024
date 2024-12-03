package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("day3.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	part2(string(data))
}

func part1(input string) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	sum := 0

	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += mul(num1, num2)
	}

	fmt.Println(sum)
}

func part2(input string) {

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	re2 := regexp.MustCompile(`do\(\)`)
	re3 := regexp.MustCompile(`don't\(\)`)
	mulMatches := re.FindAllStringSubmatch(input, -1)
	mulPositions := re.FindAllStringIndex(input, -1)
	dos := re2.FindAllStringIndex(input, -1)
	donts := re3.FindAllStringIndex(input, -1)
	sum := 0

	for i, match := range mulMatches {
		if shouldCountMul(mulPositions[i][0], dos, donts) {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			result := mul(num1, num2)
			sum += result
		}
	}

	fmt.Println(sum)
}

func mul(num1 int, num2 int) int {
	return num1 * num2
}

func shouldCountMul(mulPos int, dos [][]int, donts [][]int) bool {
	doPos := -1
	dontPos := -1

	for _, do := range dos {
		if do[0] < mulPos && do[0] > doPos {
			doPos = do[0]
		}
	}

	for _, dont := range donts {
		if dont[0] < mulPos && dont[0] > dontPos {
			dontPos = dont[0]
		}
	}

	if doPos == -1 && dontPos == -1 {
		return true
	}

	return doPos > dontPos
}
