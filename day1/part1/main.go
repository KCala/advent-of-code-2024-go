package main

import (
	"advent-of-code-2024-go/file"
	"fmt"
	"slices"
)

func main() {
	lines := file.MustReadInput()
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
	slices.Sort(left)
	slices.Sort(right)

	distances := make([]int, len(left))
	for i := 0; i < len(left); i++ {
		distances[i] = abs(left[i] - right[i])
	}

	totalDistance := 0
	for _, d := range distances {
		totalDistance += d
	}
	return totalDistance
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
