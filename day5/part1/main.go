package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const InputFile = "day5/part1/input.txt"

func main() {
	ranges, ids := readFile()

	count := 0
	for _, id := range ids {
		if isFresh(id, ranges) {
			count++
		}
	}

	fmt.Println(count)
}

func isFresh(id int, ranges [][]int) bool {
	for _, r := range ranges {
		if r[0] <= id && id <= r[1] {
			return true
		}
	}

	return false
}

func readFile() (ranges [][]int, ids []int) {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

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
		} else if line != "" { // id
			id, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			ids = append(ids, id)
		}

	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return ranges, ids
}
