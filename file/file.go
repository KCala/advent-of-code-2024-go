package file

import (
	"fmt"
	"os"
	"strings"
)

func MustReadInput() []string {
	path := mustInputFile()
	return MustReadFileLines(path)
}

func mustInputFile() string {
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " <input file>")
		os.Exit(1)
	}
	input := os.Args[1]
	return input
}

func MustReadFileLines(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	trimmed := strings.TrimRight(string(content), "\n\t")
	lines := strings.Split(trimmed, "\n")
	return lines
}
