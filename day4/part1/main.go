package main

import (
	"bufio"
	"fmt"
	"os"
)

const InputFile = "day4/part1/input.txt"

func main() {
	grid := readFile(convert)

	count := 0
	// i=0,j=0 is the top left corner
	for i := range grid { // i vertical down
		for j := range grid[i] { // j horizontal to the right
			if grid[i][j] == 1 && noOfNeighbours(i, j, grid) < 4 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func noOfNeighbours(i int, j int, grid [][]int) int {
	count := 0

	count += getPosition(i-1, j-1, grid) // top left
	count += getPosition(i-1, j, grid)   // top
	count += getPosition(i-1, j+1, grid) // top right

	count += getPosition(i, j-1, grid) // left
	count += getPosition(i, j+1, grid) // right

	count += getPosition(i+1, j-1, grid) // bottom left
	count += getPosition(i+1, j, grid)   // bottom
	count += getPosition(i+1, j+1, grid) // bottom right

	return count
}

// getPosition for out-of-bounds read returns 0
func getPosition(i int, j int, grid [][]int) int {
	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[i])-1 {
		return 0
	}
	return grid[i][j]
}

func readFile[T any](convert func(string) T) []T {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines := make([]T, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, convert(scanner.Text()))
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func convert(line string) []int {
	rolls := make([]int, len(line))
	for i := range line {
		if line[i] == '@' {
			rolls[i] = 1
		}
	}

	return rolls
}
