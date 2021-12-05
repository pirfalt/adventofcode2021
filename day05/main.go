package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type point struct {
	x, y int
}
type line struct {
	p1, p2 point
}

func (l line) isHorizontal() bool {
	return l.p1.y == l.p2.y
}

func (l line) isVertical() bool {
	return l.p1.x == l.p2.x
}

func (l line) isDiagonal() bool {
	xd := l.p1.x - l.p2.x
	yd := l.p1.y - l.p2.y

	abs := func(nr int) int {
		if nr > 0 {
			return nr
		}
		return 0 - nr
	}

	return abs(xd) == abs(yd)
}

func (l line) crossedPoints(countDiagonal bool) []point {
	points := []point{}
	if l.isHorizontal() {
		y := l.p1.y

		xLow, xHigh := l.p1.x, l.p2.x // min(x1,x2), max(x1,x2)
		if xHigh < xLow {
			xLow, xHigh = xHigh, xLow
		}

		for x := xLow; x <= xHigh; x++ {
			points = append(points, point{x, y})
		}
	}
	if l.isVertical() {
		x := l.p1.x

		yLow, yHigh := l.p1.y, l.p2.y // min(y1,y2), may(y1,y2)
		if yHigh < yLow {
			yLow, yHigh = yHigh, yLow
		}

		for y := yLow; y <= yHigh; y++ {
			points = append(points, point{x, y})
		}
	}
	if countDiagonal && l.isDiagonal() {
		// step direction for x is always +1
		xStep := 1
		// Swap so p1.x is always smaller than p2.x
		p1, p2 := l.p1, l.p2
		if p1.x > p2.x {
			p1, p2 = p2, p1
		}

		// Step direction for y may be -1
		yStep := 1
		if p1.y > p2.y {
			yStep = -1
		}

		// log.Println("d", l)
		for x, y := p1.x, p1.y; x <= p2.x; x, y = x+xStep, y+yStep {
			p := point{x, y}
			// log.Println("p", p)
			points = append(points, p)
		}
	}
	return points
}

func printMarks(board []int, boardSize int) string {
	ss := make([][]string, boardSize)
	for i, m := range board {
		s := fmt.Sprintf("%d", m)
		row := i / boardSize
		ss[row] = append(ss[row], s)
	}
	rows := make([]string, 0)
	for _, s := range ss {
		rows = append(rows, strings.Join(s, ""))
	}

	s := strings.Join(rows, "\n")
	return s
}

func one(in io.Reader) int {
	var (
		sfmt = "%d,%d -> %d,%d\n"
		l    line
	)
	lines := []line{}
	for {
		// Scan "%d,%d -> %d,%d\n" into the poins of a line
		_, err := fmt.Fscanf(in, sfmt, &l.p1.x, &l.p1.y, &l.p2.x, &l.p2.y)
		if err != nil {
			break
		}
		// Drop irrelevant lines early
		if !(l.isHorizontal() || l.isVertical()) {
			continue
		}
		lines = append(lines, l)
	}
	// log.Println(lines)

	// Mark all lines crossed points in a list
	// The list is used as a "2d" map, of (point.x + point.y*size) cordinates.
	boardSize := 1000 // 10 is enough for the example, which can be printed
	marks := make([]int, boardSize*boardSize)
	for _, l := range lines {
		for _, p := range l.crossedPoints(false) { // Do not include diagonals
			markIdx := (p.x + p.y*boardSize)
			// log.Println("p", p, markIdx)
			marks[markIdx]++
		}
		// log.Printf("marks\n%s", printMarks(marks))
	}

	if boardSize == 10 {
		// Print the board, for human firendly debugging
		log.Printf("marks\n%s", printMarks(marks, boardSize))
	}

	// Count the results
	ret := 0
	for _, crossing := range marks {
		if crossing >= 2 {
			ret++
		}
	}

	return ret
}

func two(in io.Reader) int {
	var (
		sfmt = "%d,%d -> %d,%d\n"
		l    line
	)
	lines := []line{}
	for {
		_, err := fmt.Fscanf(in, sfmt, &l.p1.x, &l.p1.y, &l.p2.x, &l.p2.y)
		if err != nil {
			break
		}

		// Check for `isDiagonal` as well
		if !(l.isHorizontal() || l.isVertical() || l.isDiagonal()) {
			continue
		}
		lines = append(lines, l)
	}

	boardSize := 1000
	marks := make([]int, boardSize*boardSize)
	for _, l := range lines {
		for _, p := range l.crossedPoints(true) { // Include diagnonal
			markIdx := (p.x + p.y*boardSize)
			marks[markIdx]++
		}
	}

	if boardSize == 10 {
		log.Printf("marks\n%s", printMarks(marks, boardSize))
	}

	ret := 0
	for _, crossing := range marks {
		if crossing >= 2 {
			ret++
		}
	}

	return ret
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
