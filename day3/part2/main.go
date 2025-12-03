package main

import (
	"bufio"
	"fmt"
	"os"
)

const InputFile = "day3/part2/input.txt"

func main() {
	sum := 0
	for _, bank := range readFile(convert) {
		sum += maxJoltage(bank)
	}

	fmt.Println(sum)
}

func maxJoltage(bank []int) int {
	var maxN, n int
	i := -1

	for numbersLeft := 12; numbersLeft > 0; numbersLeft-- {
		end := len(bank) - numbersLeft + 1
		n, i = maxAndIndex(i+1, bank[:end])
		maxN = (maxN * 10) + n
	}

	return maxN
}

func maxAndIndex(start int, numbers []int) (int, int) {
	maxN := numbers[start]
	maxI := start

	for i := start; i < len(numbers); i++ {
		if numbers[i] > maxN {
			maxN = numbers[i]
			maxI = i
		}
	}

	return maxN, maxI
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
