package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const InputFile = "day6/part2/input.txt"

func main() {
	lines, operations := readFile()
	numbers := transposeConvert(lines)

	sum := 0
	for i := range numbers {
		if operations[i] == '*' {
			sum += mul(numbers[i]...)
		} else { // +
			sum += add(numbers[i]...)
		}
	}

	fmt.Println(sum)
}

func add(ints ...int) int {
	sum := 0
	for _, n := range ints {
		sum += n
	}
	return sum
}

func mul(ints ...int) int {
	result := 1
	for _, n := range ints {
		result *= n
	}
	return result
}

func transposeConvert(matrix []string) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]int, strings.Count(matrix[0], " ")+1)

	k := 0
	for i := 0; i < cols; i++ {
		number := ""
		for j := 0; j < rows; j++ {
			number += string(matrix[j][i])
		}

		n, err := strconv.Atoi(strings.TrimSpace(number))
		if err != nil {
			k++
		} else {
			result[k] = append(result[k], n)
		}
	}

	return result[:k+1]
}

func readFile() (lines []string, operations string) {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return lines[:len(lines)-1], strings.ReplaceAll(lines[len(lines)-1], " ", "")
}
