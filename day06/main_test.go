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
	return strings.NewReader(`3,4,3,1,2`)
}

func Test_one(t *testing.T) {
	tests := []struct {
		name  string
		in    io.Reader
		steps int
		want  int
	}{
		{
			name:  "example1",
			in:    exampleFile(),
			steps: 18,
			want:  26,
		},
		{
			name:  "example2",
			in:    exampleFile(),
			steps: 80,
			want:  5934,
		},
		{
			name:  "real",
			in:    inputFile(),
			steps: 80,
			want:  356190,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := one(tt.in, tt.steps)
			if got != tt.want {
				t.Fatalf("want: %v != got: %v", tt.want, got)
			}
		})
	}
}

func Test_two(t *testing.T) {
	tests := []struct {
		name  string
		in    io.Reader
		steps int
		want  int
	}{
		{
			name:  "example1",
			in:    exampleFile(),
			steps: 18,
			want:  26,
		},
		{
			name:  "example2",
			in:    exampleFile(),
			steps: 80,
			want:  5934,
		},
		{
			name:  "real",
			in:    inputFile(),
			steps: 80,
			want:  356190,
		},
		{
			name:  "part2",
			in:    exampleFile(),
			steps: 256,
			want:  26984457539,
		},

		{
			name:  "real",
			in:    inputFile(),
			steps: 256,
			want:  1617359101538,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := two(tt.in, tt.steps)
			if got != tt.want {
				t.Fatalf("want: %v != got: %v", tt.want, got)
			}
		})
	}
}
