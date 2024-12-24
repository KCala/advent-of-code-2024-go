package main

import (
	"advent-of-code-2024-go/file"
	"fmt"
	"strconv"
	"strings"
)

type Report = struct {
	Levels []int
}

func main() {
	lines := file.MustReadInput()
	reports := parse(lines)
	solution := solve(reports)
	fmt.Println(solution)
}

func parse(lines []string) []Report {
	reports := make([]Report, 0)
	for _, l := range lines {
		levelsStr := strings.Fields(l)
		levelsD := make([]int, 0)
		for _, s := range levelsStr {
			lvl, _ := strconv.Atoi(s)
			levelsD = append(levelsD, lvl)
		}
		report := Report{
			Levels: levelsD,
		}
		reports = append(reports, report)
	}
	return reports
}

// solve returns a number of invalid reports
func solve(reports []Report) int {
	valid := 0
	for _, r := range reports {
		if reportValid(r) {
			valid += 1
		}
	}
	return valid
}

func reportValid(report Report) bool {
	ls := report.Levels
	increasing := ls[0] < ls[1]

	for i := 1; i < len(ls); i++ {
		if (ls[i-1] > ls[i]) == increasing {
			return false
		}
		diff := abs(ls[i] - ls[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
