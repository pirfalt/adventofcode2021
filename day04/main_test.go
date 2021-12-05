package main

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func inputFile() io.Reader {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return f
}
func exampleFile() io.Reader {
	return strings.NewReader(`7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7

`)
}

func Test_one(t *testing.T) {
	tests := []struct {
		name string
		in   io.Reader
		want int
	}{
		{
			name: "example",
			in:   exampleFile(),
			want: 4512,
		},
		{
			name: "real",
			in:   inputFile(),
			want: 23177,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := one(tt.in)
			if got != tt.want {
				t.Fatalf("want: %v != got: %v", tt.want, got)
			}
		})
	}
}

func Test_checkBoard(t *testing.T) {
	tests := []struct {
		name  string
		board []int
		in    []int
		want  bool
	}{
		{
			name: "win",
			board: []int{
				14, 21, 17, 24, 4,
				10, 16, 15, 9, 19,
				18, 8, 23, 26, 20,
				22, 11, 13, 6, 5,
				2, 0, 12, 3, 7,
			},
			in:   []int{14, 21, 17, 24, 4},
			want: true,
		},

		{
			name: "win row2",
			board: []int{
				10, 16, 15, 9, 19,
				14, 21, 17, 24, 4,
				18, 8, 23, 26, 20,
				22, 11, 13, 6, 5,
				2, 0, 12, 3, 7,
			},
			in:   []int{14, 21, 17, 24, 4},
			want: true,
		},

		{
			name: "win col",
			board: []int{
				10, 14, 16, 15, 9,
				18, 21, 8, 23, 26,
				22, 17, 11, 13, 6,
				2, 24, 0, 12, 3,
				19, 4, 20, 5, 7,
			},
			in:   []int{14, 21, 17, 24, 4},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkBoard(tt.board, tt.in); got != tt.want {
				t.Errorf("checkBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_two(t *testing.T) {
	tests := []struct {
		name string
		in   io.Reader
		want int
	}{
		{
			name: "example",
			in:   exampleFile(),
			want: 1924,
		},
		{
			name: "real",
			in:   inputFile(),
			want: 6804,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := two(tt.in)
			if got != tt.want {
				t.Fatalf("want: %v != got: %v", tt.want, got)
			}
		})
	}
}
