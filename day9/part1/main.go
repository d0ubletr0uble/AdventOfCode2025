package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const InputFile = "day9/part1/input.txt"

func main() {
	data := readFile(convert)
	fmt.Println(maxArea(data))
}

func maxArea(data [][2]int) int {
	result := 0
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			result = max(result, area(data[i], data[j]))
		}
	}
	return result
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
