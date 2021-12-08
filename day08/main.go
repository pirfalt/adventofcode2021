package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func one(inp []string) int {
	sum := 0
	for _, s := range inp {
		parts := strings.Split(s, "|")
		signals, outputs := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		_ = signals
		// log.Println("s:", signals)
		// log.Println("o:", outputs)
		oo := strings.Split(outputs, " ")
		// log.Println(oo)

		for _, o := range oo {
			switch len(o) {
			case 2, 3, 4, 7:
				sum++
			}
		}
	}

	return sum
}

func two(inp []string) int {
	sum := 0
	for _, s := range inp {
		lineParts := strings.Split(s, "|")
		signals, outputs := strings.TrimSpace(lineParts[0]), strings.TrimSpace(lineParts[1])
		// log.Println("s:", signals)
		// log.Println("o:", outputs)
		sigs := strings.Split(signals, " ")
		outs := strings.Split(outputs, " ")
		// log.Println(oo)

		m := map[string]int{}
		find := func(nr int) string {
			for s, n := range m {
				if n == nr {
					return s
				}
			}
			return ""
		}

		// Three passes

		// First, from only inputs
		for _, sig := range sigs {
			o := sortString(sig)
			switch len(o) {
			case 2:
				m[o] = 1
			case 3:
				m[o] = 7
			case 4:
				m[o] = 4
			case 7:
				m[o] = 8
			}
		}

		// Second, depends on values from first
		for _, sig := range sigs {
			o := sortString(sig)
			switch len(o) {
			case 6: // 0 | 6 | 9
				// contins 7 -> 0
				// contins 4 -> 9
				// -> 6
				switch {
				case contains(o, find(4)):
					m[o] = 9
				case contains(o, find(7)):
					m[o] = 0
				default:
					m[o] = 6
				}
			}
		}

		// Third, depends on values from second
		for _, sig := range sigs {
			o := sortString(sig)
			switch len(o) {
			case 5: // 2 | 3 | 5
				// contains 1 -> 3
				// iscontainedby 9 -> 5
				// -> 2
				switch {
				case contains(o, find(1)):
					m[o] = 3
				case contains(find(9), o):
					m[o] = 5
				default:
					m[o] = 2
				}
			}
		}

		// log.Println(m)

		// Passes done, build display number
		display := 0
		for i := 0; i < 4; i++ {
			o := outs[i]
			s := sortString(o)
			nr := m[s]
			// log.Println(display, s, nr)
			display = display*10 + nr
		}
		// log.Println(display)

		// Add to sum
		sum += display
	}

	return sum
}

func sortString(in string) string {
	cc := strings.Split(in, "")
	sort.Strings(cc)
	return strings.Join(cc, "")
}

func contains(in, target string) bool {
	runes := map[rune]struct{}{}
	for _, r := range in {
		runes[r] = struct{}{}
	}

	for _, r := range target {
		_, found := runes[r]
		if !found {
			return false
		}
	}
	return true
}

func parseInput(in io.Reader) []string {
	buf, _ := io.ReadAll(in)
	ss := strings.Split(strings.TrimSpace(string(buf)), "\n")
	return ss
}

func main() {
	in, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	// One
	log.Printf("one: %v", one(parseInput(in)))

	// Two
	log.Printf("two: %v", two(parseInput(in)))
}
