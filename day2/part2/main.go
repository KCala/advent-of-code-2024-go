package main

import (
	"advent-of-code-2024-go/file"
	"fmt"
	"strconv"
	"strings"
)

type Report struct {
	Levels []int
}

func main() {
	lines := file.MustReadInputAsLines()
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
		if reportValidAsIs(r) || reportValidWithDeletion(r) {
			valid += 1
		}
	}
	return valid
}

func reportValidAsIs(report Report) bool {
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

func reportValidWithDeletion(report Report) bool {
	for i := 0; i < len(report.Levels); i++ {
		if reportValidAsIs(report.without(i)) {
			return true
		}
	}
	return false
}

func reportValid(report Report, triesLeft int) bool {

	if triesLeft <= 0 {
		return false
	}

	withoutFirst := reportValid(report.without(0), triesLeft-1)
	if withoutFirst == true {
		return true
	}

	ls := report.Levels
	increasing := ls[0] < ls[1]

	for i := 1; i < len(ls); i++ {
		if (ls[i-1] > ls[i]) == increasing {
			return reportValid(report.without(i), triesLeft-1)
		}
		diff := abs(ls[i] - ls[i-1])
		if diff < 1 || diff > 3 {
			return reportValid(report.without(i), triesLeft-1)
		}
	}
	return true
}

func (report Report) without(n int) Report {
	return Report{Levels: without(report.Levels, n)}
}

func without(s []int, n int) []int {
	a := make([]int, len(s))
	copy(a, s)
	return append(a[:n], a[n+1:]...)
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
