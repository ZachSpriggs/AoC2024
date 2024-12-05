package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	before int
	after  int
}

func main() {
	data, err := os.ReadFile("day5.txt")
	if err != nil {
		fmt.Println("error reading file", err)
		return
	}
	part1(string(data))
}

func part1(input string) {
	sections := strings.Split(input, "\n\n")
	rules := parseRules(sections[0])
	updates := parseUpdates(sections[1])

	sum := 0

	for _, update := range updates {
		if isValid(update, rules) {
			midIdx := len(update) / 2
			sum += update[midIdx]
		}
	}
	fmt.Println(sum)
}

func part2(input string) {

}

// Helper functions

func parseRules(input string) []Rule {
	rules := make([]Rule, 0)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		parts := strings.Split(line, "|")
		before, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		after, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		rules = append(rules, Rule{before: before, after: after})
	}
	return rules
}

func parseUpdates(input string) [][]int {
	updates := make([][]int, 0)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		update := make([]int, 0)
		numStrs := strings.Split(line, ",")
		for _, numStr := range numStrs {
			num, _ := strconv.Atoi(strings.TrimSpace(numStr))
			update = append(update, num)
		}
		updates = append(updates, update)
	}
	return updates
}

func isValid(update []int, rules []Rule) bool {
	positions := make(map[int]int)
	for i, num := range update {
		positions[num] = i
	}

	for _, rule := range rules {
		beforePos, beforeExists := positions[rule.before]
		afterPos, afterExists := positions[rule.after]

		if beforeExists && afterExists {
			if beforePos > afterPos {
				return false
			}
		}
	}
	return true
}
