package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const InputFile = "day5/part2/input.txt"

func main() {
	ranges := clip(readFile())

	sum := 0
	for _, r := range ranges {
		sum += r[1] - r[0] + 1
	}

	fmt.Println(sum)
}

// clip removes overlapping ranges in sorted slice
func clip(ranges [][]int) [][]int {
	slices.SortFunc(ranges, func(a, b []int) int {
		return a[0] - b[0]
	})

	result := make([][]int, 0, len(ranges))
	result = append(result, ranges[0])

	last := ranges[0]
	for _, r := range ranges[1:] {
		if r[0] <= last[1] {
			r[0] = last[1] + 1
		}
		if last[1] < r[1] {
			last = r
			result = append(result, last)
		}
	}
	return result
}

func readFile() [][]int {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ranges := make([][]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "-") { // range
			parts := strings.Split(line, "-")
			p1, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			p2, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, []int{p1, p2})
		} else if line == "" {
			return ranges
		}

	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return ranges
}
