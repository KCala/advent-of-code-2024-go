package file

import (
	"fmt"
	"os"
	"strings"
)

func MustReadInputAsLines() []string {
	path := mustInputFile()
	return MustReadFileLines(path)
}

func MustReadInput() string {
	path := mustInputFile()
	return MustReadFile(path)
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
	content := MustReadFile(path)
	trimmed := strings.TrimRight(content, "\n\t")
	lines := strings.Split(trimmed, "\n")
	return lines
}

func MustReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(content)
}
