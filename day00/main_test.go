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
	return strings.NewReader(`exmple input`)
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
			want: 0,
		},
		{
			name: "real",
			in:   inputFile(),
			want: 0,
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
			want: 0,
		},
		{
			name: "real",
			in:   inputFile(),
			want: 0,
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
