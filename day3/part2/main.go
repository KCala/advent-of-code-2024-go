package main

import (
	"advent-of-code-2024-go/file"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Mul struct {
	a int
	b int
}

func (m Mul) mul() int {
	return m.a * m.b
}

func main() {
	input := file.MustReadInput()
	muls := parse(input)
	solution := solve(muls)
	fmt.Println(solution)
}

func parse(input string) []Mul {

	wrappedInput := "do()" + input + "don't()"
	activePartsRegex := regexp.MustCompile(`(?s)do\(\)(.*?)don't\(\)`)
	activeMatches := activePartsRegex.FindAllStringSubmatch(wrappedInput, -1)
	activeParts := make([]string, 0)
	for _, m := range activeMatches {
		activeParts = append(activeParts, m[1])
	}

	activeString := strings.Join(activeParts, "")

	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	ss := r.FindAllStringSubmatch(activeString, -1)
	muls := make([]Mul, 0)
	for _, s := range ss {
		a, _ := strconv.Atoi(s[1])
		b, _ := strconv.Atoi(s[2])
		mul := Mul{
			a: a,
			b: b,
		}
		muls = append(muls, mul)
	}
	return muls
}

// solve returns a number of invalid reports
func solve(muls []Mul) int {
	total := 0
	for _, m := range muls {
		total += m.mul()
	}
	return total
}
