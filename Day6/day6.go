package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	row, col int
}

func main() {
	data, err := os.ReadFile("day6.txt")
	if err != nil {
		fmt.Println("error reading file", err)
		return
	}
	part1(string(data))
}

func part1(input string) {
	grid := parseInput(input)
	visited := make(map[Position]bool)

	startPos := findStart(grid)

	count := simulatePatrol(grid, startPos, visited)
	fmt.Println("Positions: ", count)
}

func part2(input string) {

}

func parseInput(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(strings.TrimSpace(line))
	}
	return grid
}

func findStart(grid [][]rune) Position {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '^' || grid[i][j] == '>' ||
				grid[i][j] == 'v' || grid[i][j] == '<' {
				return Position{i, j}
			}
		}
	}
	return Position{}
}

func simulatePatrol(grid [][]rune, start Position, visited map[Position]bool) int {
	currentPos := start
	currentDir := grid[start.row][start.col]

	moves := map[rune]Position{
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
		'<': {0, -1},
	}
	turnRight := map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}

	for {
		visited[currentPos] = true

		move := moves[currentDir]
		nextPos := Position{
			row: currentPos.row + move.row,
			col: currentPos.col + move.col,
		}

		if nextPos.row < 0 || nextPos.row >= len(grid) ||
			nextPos.col < 0 || nextPos.col >= len(grid[0]) {
			break
		}

		if grid[nextPos.row][nextPos.col] == '#' {
			currentDir = turnRight[currentDir]
		} else {
			currentPos = nextPos
			grid[currentPos.row][currentPos.col] = currentDir
		}
	}
	return len(visited)
}
