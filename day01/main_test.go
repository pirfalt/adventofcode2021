package main

import (
	"testing"

	"github.com/pirfalt/adventofcode2021/input"
)

func Test_one(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want int
	}{
		{
			name: "example",
			in: []string{
				"199",
				"200",
				"208",
				"210",
				"200",
				"207",
				"240",
				"269",
				"260",
				"263",
			},
			want: 7,
		},
		{
			name: "real",
			in:   input.ReadLines("./input.txt"),
			want: 1709,
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
		in   []string
		want int
	}{
		{
			name: "example",
			in: []string{
				"199",
				"200",
				"208",
				"210",
				"200",
				"207",
				"240",
				"269",
				"260",
				"263",
			},
			want: 5,
		},
		{
			name: "real",
			in:   input.ReadLines("./input.txt"),
			want: 1761,
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
