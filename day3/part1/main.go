package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const InputFile = "day3/part1/input.txt"

func main() {
	sum := 0
	for _, bank := range readFile(convert) {
		sum += maxJoltage(bank)
	}

	fmt.Println(sum)
}

func maxJoltage(bank []int) int {
	n1 := slices.Max(bank[0 : len(bank)-1])
	i := slices.Index(bank, n1)
	n2 := slices.Max(bank[i+1:])

	return n1*10 + n2
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
	numbers := make([]int, len(line))
	for i := range line {
		numbers[i] = int(line[i]) - '0'
	}

	return numbers
}
