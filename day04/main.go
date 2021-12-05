package main

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/pirfalt/adventofcode2021/input"
)

func one(in io.Reader) int {

	data, err := io.ReadAll(in)
	if err != nil {
		log.Fatal(err)
	}

	// Raw input blocks
	dd := strings.Split(string(data), "\n\n")

	// [nrsInput, boards...]
	nrsInput := strings.Split(dd[0], ",")
	boardInputs := dd[1:]

	nrs := input.ParseInts(nrsInput)

	// boardInputs ([][]strings) -> boards ([][]int)
	boards := make([][]int, 0, len(boardInputs))
	for _, bi := range boardInputs {
		boardOneline := strings.ReplaceAll(bi, "\n", " ")
		boardStrings := strings.Split(boardOneline, " ")
		board := input.ParseInts(boardStrings)

		if len(board) == 0 {
			break
		}
		boards = append(boards, board)
	}

	// log.Printf("nrs: %v\tbs: %v", nrs, boards)

	for i, latestNr := range nrs {
		takenNrs := nrs[0 : i+1]
		for _, b := range boards {
			if checkBoard(b, takenNrs) {
				// log.Println(i, b)
				unmarked := sumOfUnmarked(b, takenNrs)
				// log.Println(unmarked, latestNr)
				return unmarked * latestNr
			}
		}
	}

	return 0
}

func checkBoard(board []int, in []int) bool {
	inSet := map[int]struct{}{}
	for _, nr := range in {
		inSet[nr] = struct{}{}
	}

	// log.Printf("%v\t%v\t%v", board, in, inSet)

	for row := 0; row < 5; row++ {
		rowWin := false

		for col := 0; col < 5; col++ {
			boardNr := board[row*5+col]
			_, found := inSet[boardNr]
			// log.Printf("%v\t%v\t%v\t%v\t", row, col, boardNr, found)
			if !found {
				rowWin = false
				break
			}
			rowWin = true
		}

		if rowWin {
			return true
		}
	}

	for col := 0; col < 5; col++ {
		colWin := false

		for row := 0; row < 5; row++ {
			boardNr := board[row*5+col]
			_, found := inSet[boardNr]
			// log.Printf("%v\t%v\t%v\t%v\t", row, col, boardNr, found)
			if !found {
				colWin = false
				break
			}
			colWin = true
		}

		if colWin {
			return true
		}
	}

	return false
}
func sumOfUnmarked(board []int, in []int) int {
	inSet := map[int]struct{}{}
	for _, nr := range in {
		inSet[nr] = struct{}{}
	}

	sum := 0
	for _, nr := range board {
		_, found := inSet[nr]
		if !found {
			sum += nr
		}
	}

	return sum
}

func two(in io.Reader) int {

	data, err := io.ReadAll(in)
	if err != nil {
		log.Fatal(err)
	}

	// Raw input blocks
	dd := strings.Split(string(data), "\n\n")

	// [nrsInput, boards...]
	nrsInput := strings.Split(dd[0], ",")
	boardInputs := dd[1:]

	nrs := input.ParseInts(nrsInput)

	// boardInputs ([][]strings) -> boards ([][]int)
	boards := make([][]int, 0, len(boardInputs))
	for _, bi := range boardInputs {
		boardOneline := strings.ReplaceAll(bi, "\n", " ")
		boardStrings := strings.Split(boardOneline, " ")
		board := input.ParseInts(boardStrings)

		if len(board) == 0 {
			break
		}
		boards = append(boards, board)
	}

	// log.Printf("nrs: %v\tbs: %v", nrs, boards)

	for i, latestNr := range nrs {
		takenNrs := nrs[0 : i+1]

		if len(boards) == 1 {
			b := boards[0]
			unmarked := sumOfUnmarked(b, takenNrs)
			// log.Println(unmarked, latestNr)
			return unmarked * latestNr
		}

		for bi, b := range boards {
			if checkBoard(b, takenNrs) {
				boards = append(boards[:bi], boards[bi+1:]...) // remove board from list
				// log.Println("dropping board", b)
			}
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
