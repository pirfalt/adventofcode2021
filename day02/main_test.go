package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_one(t *testing.T) {
	tests := []struct {
		name string
		in   io.Reader
		want int
	}{
		{
			name: "example",
			in: strings.NewReader(`forward 5
down 5
forward 8
up 3
down 8
forward 2`),
			want: 150,
		},
		{
			name: "real",
			in:   func() io.Reader { f, _ := os.Open("./input.txt"); return f }(),
			want: 1636725,
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
			in: strings.NewReader(`forward 5
down 5
forward 8
up 3
down 8
forward 2`),
			want: 900,
		},
		{
			name: "real",
			in:   func() io.Reader { f, _ := os.Open("./input.txt"); return f }(),
			want: 1872757425,
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
