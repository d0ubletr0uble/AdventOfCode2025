package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const InputFile = "day8/part2/input.txt"

type Point struct {
	X, Y, Z int
}

type Distance struct {
	Pair [2]Point
	Dist float64
}

func (p Point) Distance(p2 Point) float64 {
	return math.Sqrt(
		math.Pow(float64(p.X-p2.X), 2) +
			math.Pow(float64(p.Y-p2.Y), 2) +
			math.Pow(float64(p.Z-p2.Z), 2),
	)
}

func main() {
	data := readFile(convert)
	circuits := make([][]Point, 0, len(data))
	for _, p := range data {
		circuits = append(circuits, []Point{p})
	}
	dist := distances(data)
	slices.SortFunc(dist, func(a, b Distance) int {
		return int(a.Dist - b.Dist)
	})

	for i := range dist {
		circuits = connect(dist[i].Pair[0], dist[i].Pair[1], circuits)
		if len(circuits) == 1 {
			fmt.Println(dist[i].Pair[0].X * dist[i].Pair[1].X)
			break
		}
	}
}

func connect(p1 Point, p2 Point, circuits [][]Point) [][]Point {
	if inSameCircuit(p1, p2, circuits) {
		return circuits // connection ignored
	}

	var p1c, p2c int

	for i := range circuits {
		if slices.Contains(circuits[i], p1) {
			p1c = i
		}
		if slices.Contains(circuits[i], p2) {
			p2c = i
		}
	}

	circuits[p1c] = append(circuits[p1c], circuits[p2c]...)
	circuits = slices.Delete(circuits, p2c, p2c+1)

	return circuits
}

func distances(list []Point) []Distance {
	d := make([]Distance, 0)
	for i := range list {
		for j := i + 1; j < len(list); j++ {
			dist := list[i].Distance(list[j])
			d = append(d, Distance{Pair: [2]Point{list[i], list[j]}, Dist: dist})
		}
	}

	return d
}

func inSameCircuit(p1, p2 Point, circuits [][]Point) bool {
	p1i := slices.IndexFunc(circuits, func(c []Point) bool {
		return slices.Contains(c, p1)
	})
	p2i := slices.IndexFunc(circuits, func(c []Point) bool {
		return slices.Contains(c, p2)
	})

	return p1i == p2i
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

func convert(line string) Point {
	parts := strings.Split(line, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	z, err := strconv.Atoi(parts[2])
	if err != nil {
		panic(err)
	}

	return Point{
		X: x,
		Y: y,
		Z: z,
	}
}
