package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/pirfalt/adventofcode2021/input"
)

func one(slots []int) int {
	o, _ := run(slots)
	return o
}

func two(slots []int) int {
	_, t := run(slots)
	return t
}

func run(slots []int) (one, two int) {
	moves1 := make([]int, len(slots))
	moves2 := make([]int, len(slots))
	for target := range slots {
		for i, v := range slots {
			dist := abs(target - i)

			// One
			moves1[target] += dist * v

			// Two
			fule := 0
			for m := 1; m <= dist; m++ {
				fule += m
			}
			moves2[target] += fule * v
		}
	}

	// log.Println(moves)
	sort.Ints(moves1)
	sort.Ints(moves2)
	return moves1[0], moves2[0]
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func parseInput(in io.Reader) []int {
	dd, _ := io.ReadAll(in)
	ss := strings.Split(strings.TrimSpace(string(dd)), ",")
	positions := input.ParseInts(ss)
	sort.Ints(positions)
	// log.Println(positions)

	max := positions[len(positions)-1]
	slots := make([]int, max+1)

	for _, p := range positions {
		slots[p]++
	}
	// log.Println(slots)
	return slots
}

func main() {
	inputF, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputF.Close()

	slots := parseInput(inputF)

	// One
	log.Printf("one: %v", one(slots))

	// Two
	log.Printf("two: %v", two(slots))
}
