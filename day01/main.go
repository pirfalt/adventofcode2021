package main

import (
	"log"
	"math"

	input "github.com/pirfalt/adventofcode2021/input"
)

func one(in []string) int {
	ii := input.ParseInts(in)

	incCount := 0
	prev := math.MaxInt
	for _, i := range ii {
		if i > prev {
			incCount++
		}
		prev = i
	}
	return incCount
}

func two(in []string) int {
	ii := input.ParseInts(in)

	incCount := 0
	for i := 3; i < len(ii); i++ {
		p0, p1, p2, p3 := ii[i-3], ii[i-2], ii[i-1], ii[i]
		p := p0 + p1 + p2
		c := p1 + p2 + p3
		if c > p {
			incCount++
		}
	}
	return incCount
}

func main() {
	input := input.ReadLines("./input.txt")
	log.Printf("one: %v", one(input))
	log.Printf("two: %v", two(input))
}
