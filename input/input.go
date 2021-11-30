package input

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(fileName string) []string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("failed to read input")
	}
	s := string(f)
	return strings.Split(s, "\n")
}

func ParseInts(ss []string) []int {
	ii := make([]int, 0)
	for _, s := range ss {
		if s == "" {
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ii = append(ii, i)
	}
	return ii
}
