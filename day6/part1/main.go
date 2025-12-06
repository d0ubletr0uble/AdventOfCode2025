package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const InputFile = "day6/part1/input.txt"

func main() {
	numbers, operations := readFile()
	numbers = transpose(numbers)

	sum := 0
	for i := range numbers {
		if operations[i] == '*' {
			sum += mul(numbers[i]...)
		} else {
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

func transpose(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]int, cols)
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			result[i] = append(result[i], matrix[j][i])
		}
	}

	return result
}

func readFile() (allNumbers [][]int, operations []rune) {
	re := regexp.MustCompile(`\d+`)

	f, err := os.Open(InputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, "*") && !strings.Contains(line, "+") { // numbers
			parts := re.FindAllString(line, -1)
			numbers := make([]int, 0, len(parts))
			for _, p := range parts {
				number, err := strconv.Atoi(p)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, number)
			}
			allNumbers = append(allNumbers, numbers)
		} else { // operations
			for _, char := range line {
				if char != ' ' {
					operations = append(operations, char)
				}
			}
		}
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return allNumbers, operations
}
