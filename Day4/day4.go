package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("day4.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	grid := parseInput(string(data))

	part1(grid)
}

func part1(grid [][]rune) {
	matches := 0
	rows := len(grid)
	cols := len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if col <= cols-4 {
				checkRight(grid, row, col, &matches)
			}

			if row <= rows-4 {
				checkDown(grid, row, col, &matches)

				if col <= cols-4 {
					checkDownRight(grid, row, col, &matches)
				}

				if col >= 3 {
					checkDownLeft(grid, row, col, &matches)
				}
			}
		}
	}
	fmt.Println(matches)

}

func part2(input string) {

}

// ***** Helper functions to simplify loop logic ***** //

func parseInput(input string) [][]rune {
	// Split input into lines
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(strings.TrimSpace(line))
	}

	return grid
}

func checkRight(grid [][]rune, row, col int, matches *int) {

	if grid[row][col] == 'X' &&
		grid[row][col+1] == 'M' &&
		grid[row][col+2] == 'A' &&
		grid[row][col+3] == 'S' {
		*matches++
	}

	if grid[row][col] == 'S' &&
		grid[row][col+1] == 'A' &&
		grid[row][col+2] == 'M' &&
		grid[row][col+3] == 'X' {
		*matches++
	}
}

func checkDown(grid [][]rune, row, col int, matches *int) {

	if grid[row][col] == 'X' &&
		grid[row+1][col] == 'M' &&
		grid[row+2][col] == 'A' &&
		grid[row+3][col] == 'S' {
		*matches++
	}

	if grid[row][col] == 'S' &&
		grid[row+1][col] == 'A' &&
		grid[row+2][col] == 'M' &&
		grid[row+3][col] == 'X' {
		*matches++
	}
}

func checkDownRight(grid [][]rune, row, col int, matches *int) {

	if grid[row][col] == 'X' &&
		grid[row+1][col+1] == 'M' &&
		grid[row+2][col+2] == 'A' &&
		grid[row+3][col+3] == 'S' {
		*matches++
	}

	if grid[row][col] == 'S' &&
		grid[row+1][col+1] == 'A' &&
		grid[row+2][col+2] == 'M' &&
		grid[row+3][col+3] == 'X' {
		*matches++
	}
}

func checkDownLeft(grid [][]rune, row, col int, matches *int) {

	if grid[row][col] == 'X' &&
		grid[row+1][col-1] == 'M' &&
		grid[row+2][col-2] == 'A' &&
		grid[row+3][col-3] == 'S' {
		*matches++
	}

	if grid[row][col] == 'S' &&
		grid[row+1][col-1] == 'A' &&
		grid[row+2][col-2] == 'M' &&
		grid[row+3][col-3] == 'X' {
		*matches++
	}
}
