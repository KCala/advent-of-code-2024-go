package main

import (
	"advent-of-code-2024-go/file"
	"slices"
	"testing"
)

func TestParse(t *testing.T) {
	wantLeft := []int{3, 4, 2, 1, 3, 3}
	wantRight := []int{4, 3, 5, 3, 9, 3}

	lines := file.MustReadFileLines("example.txt")
	gotLeft, gotRight := parse(lines)

	if !slices.Equal(gotLeft, wantLeft) {
		t.Errorf("parse() gotLeft = %v, want %v", gotLeft, wantLeft)
	}

	if !slices.Equal(gotRight, wantRight) {
		t.Errorf("parse() gotRight = %v, want %v", gotRight, wantRight)
	}
}

func TestSolve(t *testing.T) {

	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}
	want := 31

	got := solve(left, right)

	if got != want {
		t.Errorf("solve() got = %v, want %v", got, want)
	}
}
