package main

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/pirfalt/adventofcode2021/input"
)

func one(in io.Reader, steps int) int {
	bb, _ := io.ReadAll(in)
	ss := strings.Split(strings.TrimSpace(string(bb)), ",")
	fishes := input.ParseInts(ss)

	// log.Printf("%v", fishes)

	for i := 0; i < steps; i++ {
		stepStartLen := len(fishes)

		// log.Printf("i - %v  %v", stepStartLen, fishes)

		for j := 0; j < stepStartLen; j++ {
			fish := fishes[j]

			if fish == 0 {
				fishes[j] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[j]--
			}
		}
	}

	// log.Printf("%v", fishes)

	return len(fishes)
}

func two(in io.Reader, steps int) int {
	bb, _ := io.ReadAll(in)
	ss := strings.Split(strings.TrimSpace(string(bb)), ",")
	fishes := input.ParseInts(ss)

	ages := make([]int, 9)
	for _, f := range fishes {
		ages[f]++
	}
	// log.Printf("%v, %v", fishes, ages)

	for step := 0; step < steps; step++ {
		// log.Printf("step - %v", ages)

		age0 := ages[0]
		for i := 0; i < len(ages)-1; i++ {
			ages[i] = ages[i+1]
		}
		ages[6] += age0
		ages[8] = age0
	}

	// log.Printf("%v", ages)
	sum := 0
	for _, a := range ages {
		sum += a
	}

	return sum
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	// One
	log.Printf("one: %v", one(input, 80))

	// Reset file reader
	if _, err = input.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	// Two
	log.Printf("two: %v", two(input, 256))
}
