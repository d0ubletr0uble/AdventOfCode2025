package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	intervals := readFile(convert)[0]

	sum := 0
	for _, interval := range intervals {
		for i := interval[0]; i <= interval[1]; i++ {
			if isInvalid(i) {
				sum += i
			}
		}
	}

	fmt.Println(sum)
}

func isInvalid(n int) bool {
	str := strconv.Itoa(n)
	for i := 1; i <= len(str)/2; i++ {
		if strings.Repeat(str[:i], len(str)/i) == str {
			return true
		}
	}

	return false
}

func readFile[T any](convert func(string) T) []T {
	f, err := os.Open("day2/part2/input.txt")
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

func convert(line string) [][2]int {
	intervals := make([][2]int, 0)

	parts := strings.Split(line, ",")
	for _, part := range parts {
		numbers := strings.Split(part, "-")
		n1, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}

		intervals = append(intervals, [2]int{n1, n2})
	}

	return intervals
}
