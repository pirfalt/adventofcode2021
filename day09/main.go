package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/pirfalt/adventofcode2021/input"
)

func one(matix [][]int) int {
	sum := 0
	for rowi, row := range matix {
		rl := len(matix)
		cl := len(row)
		for coli, v := range row {
			lOk := rowi >= 1
			rOk := rowi < rl-1
			uOk := coli >= 1
			dOk := coli < cl-1
			if lOk && matix[rowi-1][coli] <= v {
				continue
			}
			if rOk && matix[rowi+1][coli] <= v {
				continue
			}
			if uOk && matix[rowi][coli-1] <= v {
				continue
			}
			if dOk && matix[rowi][coli+1] <= v {
				continue
			}

			sum += v + 1
		}
	}
	// log.Println(low)
	return sum
}

type point struct {
	x, y int
}

func two(matix [][]int) int {
	points := map[point]int{}
	for y, r := range matix {
		for x, v := range r {
			points[point{x, y}] = v
		}
	}

	bastins := []int{}
	for p, v := range points {
		// Neigbors
		u := point{p.x, p.y - 1}
		d := point{p.x, p.y + 1}
		l := point{p.x - 1, p.y}
		r := point{p.x + 1, p.y}

		// Check if low
		if nv, ok := points[u]; ok && nv <= v {
			continue
		}
		if nv, ok := points[d]; ok && nv <= v {
			continue
		}
		if nv, ok := points[l]; ok && nv <= v {
			continue
		}
		if nv, ok := points[r]; ok && nv <= v {
			continue
		}

		// Collect neigbors
		found := map[point]struct{}{
			p: {},
		}
		collectNeigbors(points, found, p)

		// Save basin size
		bastins = append(bastins, len(found))
	}
	sort.Ints(bastins)
	last3 := bastins[len(bastins)-3:]

	res := 1
	for _, b := range last3 {
		res *= b
	}
	return res
}

func collectNeigbors(points map[point]int, found map[point]struct{}, p point) {
	v := points[p]
	// log.Println(p, v)

	// Find neigbors
	u := point{p.x, p.y - 1}
	d := point{p.x, p.y + 1}
	l := point{p.x - 1, p.y}
	r := point{p.x + 1, p.y}

	// Collect
	if nv, ok := points[u]; ok && nv > v && nv < 9 {
		found[u] = struct{}{}
		collectNeigbors(points, found, u)
	}
	if nv, ok := points[d]; ok && nv > v && nv < 9 {
		found[d] = struct{}{}
		collectNeigbors(points, found, d)
	}
	if nv, ok := points[l]; ok && nv > v && nv < 9 {
		found[l] = struct{}{}
		collectNeigbors(points, found, l)
	}
	if nv, ok := points[r]; ok && nv > v && nv < 9 {
		found[r] = struct{}{}
		collectNeigbors(points, found, r)
	}
}

func parseInput(in io.Reader) [][]int {
	buf, _ := io.ReadAll(in)
	lines := strings.Split(strings.TrimSpace(string(buf)), "\n")

	o := [][]int{}
	for _, l := range lines {
		ss := strings.Split(l, "")
		ii := input.ParseInts(ss)
		o = append(o, ii)
	}
	return o
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	matrix := parseInput(input)

	// One
	log.Printf("one: %v", one(matrix))

	// Two
	log.Printf("two: %v", two(matrix))
}
