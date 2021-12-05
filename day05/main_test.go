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
	return strings.NewReader(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
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
			want: 5,
		},
		{
			name: "real",
			in:   inputFile(),
			want: 5608,
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

func Test_two(t *testing.T) {
	tests := []struct {
		name string
		in   io.Reader
		want int
	}{
		{
			name: "example",
			in:   exampleFile(),
			want: 12,
		},
		{
			name: "real",
			in:   inputFile(),
			want: 20299,
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
