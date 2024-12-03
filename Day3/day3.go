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
	part1(string(data))
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

}

func mul(num1 int, num2 int) int {
	return num1 * num2
}
