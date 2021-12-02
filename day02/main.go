package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func one(in io.Reader) int {
	// Scan format and output vars
	lineFmt := "%s %d\n"
	var (
		cmd   string
		units int
	)

	// Movements
	m := map[string]int{}

	for {
		// Parse input using `fmt.Fscanf` 😍
		_, err := fmt.Fscanf(in, lineFmt, &cmd, &units)
		if err != nil {
			break
		}
		m[cmd] += units
	}
	// fmt.Println(m) // Check the map

	pos := m["forward"]
	dep := m["down"] - m["up"]
	return pos * dep
}

func two(in io.Reader) int {
	var (
		sfmt  = "%s %d\n"
		cmd   string
		units int
	)

	pos := 0
	aim := 0
	dep := 0

	for {
		_, err := fmt.Fscanf(in, sfmt, &cmd, &units)
		if err != nil {
			break
		}

		switch cmd {
		case "up":
			aim -= units
		case "down":
			aim += units
		case "forward":
			pos += units
			dep += aim * units
		}
		// fmt.Println(pos, aim, dep) // Check the map
	}

	return pos * dep
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	// One
	log.Printf("one: %v", one(input))

	// Reset file reader
	if _, err = input.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	// Two
	log.Printf("two: %v", two(input))
}
