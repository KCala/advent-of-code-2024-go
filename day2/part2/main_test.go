package main

import (
	"advent-of-code-2024-go/file"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	want := []Report{
		{Levels: []int{7, 6, 4, 2, 1}},
		{Levels: []int{1, 2, 7, 8, 9}},
		{Levels: []int{9, 7, 6, 2, 1}},
		{Levels: []int{1, 3, 2, 4, 5}},
		{Levels: []int{8, 6, 4, 4, 1}},
		{Levels: []int{1, 3, 6, 7, 9}},
	}

	lines := file.MustReadFileLines("example.txt")
	got := parse(lines)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("parse() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolve(t *testing.T) {

	reports := []Report{
		{Levels: []int{7, 6, 4, 2, 1}},
		{Levels: []int{1, 2, 7, 8, 9}},
		{Levels: []int{9, 7, 6, 2, 1}},
		{Levels: []int{1, 3, 2, 4, 5}},
		{Levels: []int{8, 6, 4, 4, 1}},
		{Levels: []int{1, 3, 6, 7, 9}},
		{Levels: []int{1, 7, 6, 4, 2, 1}}, //safe if 0th element removed
	}
	want := 5

	got := solve(reports)

	if got != want {
		t.Errorf("solve() got = %v, want %v", got, want)
	}
}
