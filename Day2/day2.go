package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day2.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	helper(string(data))
}

func part1(input [][]int) {
	numSafe := 0

	increasing := false
	decreasing := false
	unsafe := false

	for _, arr := range input {
		if arr[0] > arr[1] {
			decreasing = true
		} else {
			increasing = true
		}
		for i := 1; i < len(arr); i++ {
			diff := math.Abs(float64(arr[i] - arr[i-1]))
			if diff > 3 || diff < 1 {
				unsafe = true
				break
			}
			if increasing && arr[i] < arr[i-1] {
				unsafe = true
				break
			}
			if decreasing && arr[i] > arr[i-1] {
				unsafe = true
				break
			}
		}
		if !unsafe {
			numSafe++
		}
		unsafe, increasing, decreasing = false, false, false
	}
	fmt.Println(numSafe)
}

func part2(input [][]int) {
	numSafe := 0

	for _, arr := range input {
		if isSafe(arr) {
			numSafe++
		} else {
			for i := range arr {
				temp := append([]int(nil), arr...)
				newArr := append(temp[:i], temp[i+1:]...)
				if isSafe(newArr) {
					numSafe++
					break
				}
			}
		}
	}
	fmt.Println(numSafe)
}

func isSafe(arr []int) bool {
	increasing := false
	decreasing := false
	if arr[0] > arr[1] {
		decreasing = true
	} else {
		increasing = true
	}

	for i := 1; i < len(arr); i++ {
		diff := math.Abs(float64(arr[i] - arr[i-1]))
		if diff > 3 || diff < 1 {
			return false
		}
		if increasing && arr[i] < arr[i-1] {
			return false
		}
		if decreasing && arr[i] > arr[i-1] {
			return false
		}
	}
	return true
}

func helper(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	result := make([][]int, len(lines))

	for i, line := range lines {
		numStrings := strings.Fields(line)

		numbers := make([]int, len(numStrings))

		for j, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Printf("error converting string to number: %v", err)
				continue
			}
			numbers[j] = num
		}
		result[i] = numbers
	}
	part2(result)
}
