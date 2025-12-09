package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const InputFile = "day9/part2/input.txt"

func main() {
	data := readFile(convert)
	lines := toLines(data) // lines of the polygon
	fmt.Println(maxArea(data, lines))
}

func toLines(data [][2]int) [][2][2]int {
	lines := make([][2][2]int, len(data))
	for i := 0; i < len(data)-1; i++ {
		lines[i] = [2][2]int{data[i], data[i+1]}
	}
	lines[len(data)-1] = [2][2]int{data[len(data)-1], data[0]}

	return lines
}

func maxArea(data [][2]int, lines [][2][2]int) int {
	result := 0
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if isWithinPolygon(data[i], data[j], lines) {
				result = max(result, area(data[i], data[j]))
			}
		}
	}
	return result
}

func isWithinPolygon(p1, p2 [2]int, lines [][2][2]int) bool {
	leftX := min(p1[0], p2[0])
	topY := min(p1[1], p2[1])
	rightX := max(p1[0], p2[0])
	bottomY := max(p1[1], p2[1])

	for _, line := range lines {
		isLeft := line[0][0] <= leftX && line[1][0] <= leftX
		isAbove := line[0][1] <= topY && line[1][1] <= topY
		isRight := line[0][0] >= rightX && line[1][0] >= rightX
		isBelow := line[0][1] >= bottomY && line[1][1] >= bottomY

		if !isLeft && !isAbove && !isRight && !isBelow {
			return false
		}
	}

	return true
}

func area(p1, p2 [2]int) int {
	deltaX := p1[0] - p2[0]
	if deltaX < 0 {
		deltaX = -deltaX
	}
	deltaY := p1[1] - p2[1]
	if deltaY < 0 {
		deltaY = -deltaY
	}
	deltaX++
	deltaY++

	return deltaX * deltaY
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

func convert(line string) [2]int {
	parts := strings.Split(line, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return [2]int{x, y}
}
