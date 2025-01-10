package main

import (
	"advent-of-code-2024-go/file"
	"fmt"
)

func main() {
	lines := file.MustReadInputAsLines()
	left, right := parse(lines)
	solution := solve(left, right)
	fmt.Println(solution)
}

func parse(lines []string) (left []int, right []int) {
	for _, line := range lines {
		//extract 2 numbers from each line, separated by n spaces
		var l, r int
		fmt.Sscanf(line, "%d %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func solve(left []int, right []int) int {

	rightCounts := make(map[int]int, 0)
	for _, r := range right {
		rightCounts[r] += 1
	}

	score := 0
	for _, l := range left {
		score += l * rightCounts[l]
	}

	return score
}
