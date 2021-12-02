package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func one(in io.Reader) int {
	var (
		sfmt  = "%s %d\n"
		cmd   string
		units int
	)

	for {
		_, err := fmt.Fscanf(in, sfmt, &cmd, &units)
		if err != nil {
			break
		}
	}

	return 0
}

func two(in io.Reader) int {
	var (
		sfmt  = "%s %d\n"
		cmd   string
		units int
	)

	for {
		_, err := fmt.Fscanf(in, sfmt, &cmd, &units)
		if err != nil {
			break
		}
	}

	return 0
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
