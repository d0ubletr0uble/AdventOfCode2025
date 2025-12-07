package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const InputFile = "day7/part2/input.txt"

func main() {
	data := readFile(convert)

	idx := strings.IndexRune(string(data[0]), 'S')
	beams := []int{idx}

	// track how many timelines each beam carries
	timelines := map[int]int{idx: 1}

	fmt.Println(string(data[0]))

	for i := 1; i < len(data); i++ {
		for _, j := range intersect(beams, splitters(data[i])) {
			beams = slices.DeleteFunc(beams, func(beam int) bool {
				return beam == j
			})
			if !slices.Contains(beams, j-1) {
				beams = append(beams, j-1)
			}
			if !slices.Contains(beams, j+1) {
				beams = append(beams, j+1)
			}
			timelines[j-1] += timelines[j]
			timelines[j+1] += timelines[j]
			timelines[j] = 0
		}

		// beam visualization
		for _, b := range beams {
			data[i][b] = '|'
		}

		fmt.Println(string(data[i]))
	}

	fmt.Println("timelines =", addVals(timelines))
}

func addVals(m map[int]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func intersect(a []int, b []int) []int {
	result := make([]int, 0)
	for _, element := range a {
		if slices.Contains(b, element) {
			result = append(result, element)
		}
	}
	return result
}

func splitters(runes []rune) []int {
	idx := make([]int, 0)
	for i := range runes {
		if runes[i] == '^' {
			idx = append(idx, i)
		}
	}
	return idx
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

func convert(line string) []rune {
	return []rune(line)
}
