package main

import (
	"log"

	input "github.com/pirfalt/adventofcode2021/input"
)

func one(in []string) int {
	ii := input.ParseInts(in)
	for _, i0 := range ii {
		for _, i1 := range ii {
			if i0+i1 == 2020 {
				return i0 * i1
			}
		}
	}
	return -1
}

func two(in []string) int {
	ii := input.ParseInts(in)
	for _, i0 := range ii {
		for _, i1 := range ii {
			for _, i2 := range ii {
				if i0+i1+i2 == 2020 {
					return i0 * i1 * i2
				}
			}
		}
	}
	return -1
}

func main() {
	input := input.ReadLines("./input.txt")
	log.Printf("one: %v", one(input))
	log.Printf("two: %v", two(input))
}
