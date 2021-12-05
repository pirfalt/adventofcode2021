package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func one(in io.Reader) int {
	scanner := bufio.NewScanner(in)

	var (
		sfmt = "%b\n"
		line uint
	)

	var indexBitSums []uint
	lineCount := 0

	// Scan lines
	for scanner.Scan() {
		l := scanner.Text()
		_, err := fmt.Sscanf(l, sfmt, &line)
		if err != nil {
			break
		}
		lineCount++

		lineLen := len(l)
		if indexBitSums == nil {
			indexBitSums = make([]uint, lineLen)
		}

		// Mask bits and add to slice accumilators
		//  binary `1010101` -> []{+1,+0,+1,+0,+1,+0,+1}
		// log.Println("--")
		for i := 0; i < lineLen; i++ {
			mask := uint(1) << i
			bit := line & mask
			bit = bit >> i // matched bit, shifted to 0|1
			indexBitSums[lineLen-1-i] += bit
			// log.Println(rr, i, bit, mask)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Slice of bit sums -> output numbers
	gamma := uint(0)
	espilon := uint(0)
	for i := 0; i < len(indexBitSums); i++ {
		bitSum := indexBitSums[i]
		isMostCommon := bitSum > uint(lineCount)/2

		shiftLen := uint(len(indexBitSums) - 1 - i)
		bit := uint(1) << shiftLen
		if isMostCommon {
			gamma = gamma | bit
		} else {
			espilon = espilon | bit
		}

		// log.Printf("%d %t   %d  %d", bitSum, bitSum > uint(lineCount)/2, gamma, espilon)
	}

	return int(gamma) * int(espilon)
}

func two(in io.Reader) int {
	scanner := bufio.NewScanner(in)

	var (
		sfmt = "%b\n"
		line uint
	)

	var (
		lines   []uint
		lineLen int
	)

	// Scan lines
	for scanner.Scan() {
		l := scanner.Text()

		_, err := fmt.Sscanf(l, sfmt, &line)
		if err != nil {
			break
		}
		if lines == nil {
			lineLen = len(l)
			lines = make([]uint, 0)
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// log.Println(lineLen, lines)

	// Run for oxygen and scrubber
	// Could probably do it in one loop, but functions are easy.
	oxygen := run(lines, lineLen, true)
	// log.Printf("=============")
	scrubber := run(lines, lineLen, false)

	// log.Printf("oxygen: %v, scrubber: %v", oxygen, scrubber)

	return int(oxygen) * int(scrubber)
}

func run(lines []uint, lineLen int, negate bool) uint {
	for pos := lineLen - 1; pos >= 0; pos-- {

		// Count `1`:s at position
		posSum := uint(0)
		for _, l := range lines {
			bit := (l >> pos) & uint(1)
			if bit > 0 {
				posSum += 1
			}
		}

		// `oneCommon` = `1` is the most "common" at pos
		// log.Println("maskHits", maskHits)
		equalCount := int(posSum)*2 == len(lines)
		oneCommon := equalCount || posSum > uint(len(lines))/2

		// Kept values
		keep := make([]uint, 0, len(lines))
		for _, l := range lines {
			// Mask bit value for pos
			bit := (l >> pos) & uint(1)
			// log.Println(oneCommon, bit)

			if !negate {
				// `1` is common and bit > 1
				if oneCommon && (bit > 0) {
					keep = append(keep, l)
				}

				// `0` is common and bit == 0
				if !oneCommon && !(bit > 0) {
					keep = append(keep, l)
				}
			} else {
				// Same as above, but negated.
				// Using more branches instead of smarter conditions.
				if oneCommon && !(bit > 0) {
					keep = append(keep, l)
				}

				if !oneCommon && (bit > 0) {
					keep = append(keep, l)
				}
			}
		}

		// Return value if it's the last
		// log.Printf("lines:%v, keep: %v", lines, keep)
		if len(keep) == 1 {
			return keep[0]
		}

		// Keep only the relevant values
		lines = keep
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
	// log.Printf("one: %v", one(input))

	// Reset file reader
	if _, err = input.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	// Two
	// log.Printf("two: %v", two(input))
}
