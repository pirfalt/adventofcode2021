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
	return strings.NewReader(`2199943210
3987894921
9856789892
8767896789
9899965678`)
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
			want: 15,
		},
		{
			name: "real",
			in:   inputFile(),
			want: 462,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := one(parseInput(tt.in))
			if got != tt.want {
				t.Fatalf("want: %v != got: %v", tt.want, got)
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
			want: 1134,
		},
		{
			name: "real",
			in:   inputFile(),
			want: 1397760,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := two(parseInput(tt.in))
			if got != tt.want {
				t.Fatalf("want: %v != got: %v", tt.want, got)
			}
		})
	}
}
