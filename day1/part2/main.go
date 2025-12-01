package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	count := 0
	position := 50

	fmt.Println("Starting position: ", position)
	for _, n := range readFile(convert) {
		move := position + n

		revolutions := move / 100
		if revolutions < 0 {
			revolutions = -revolutions
		}
		if move < 0 && position != 0 {
			revolutions++
		}

		position = move % 100
		if position < 0 {
			position += 100
		}

		if revolutions > 0 && position == 0 {
			revolutions--
		}

		if position == 0 {
			count++
		}

		count += revolutions

		fmt.Println(
			"Move: ", n,
			"Position: ", position,
			"Count: ", count,
			"Revolutions: ", revolutions,
		)
	}

	fmt.Println(count)
}

func readFile[T any](convert func(string) T) []T {
	f, err := os.Open("day1/part2/input.txt")
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

func convert(line string) int {
	number, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(err)
	}

	if line[0] == 'L' {
		number = -number
	}

	return number
}
