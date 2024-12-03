package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day1.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	part1(string(data))
}

func part1(input string) {
	leftArr := make([]int, 0)
	rightArr := make([]int, 0)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("error converting left num: ", err)
			continue
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("error converting right num: ", err)
			continue
		}
		leftArr = append(leftArr, left)
		rightArr = append(rightArr, right)
	}

	slices.Sort(leftArr)
	slices.Sort(rightArr)

	sum := 0

	for i := 0; i < len(leftArr); i++ {
		diff := leftArr[i] - rightArr[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	// fmt.Println(sum)
	part2(leftArr, rightArr)
}

func part2(leftArray []int, rightArray []int) {
	counts := make(map[int]int)

	for _, num := range rightArray {
		counts[num]++
	}

	score := 0
	for _, num := range leftArray {
		if count, exists := counts[num]; exists {
			score += num * count
		}
	}

	fmt.Println(score)
}
