package main

import (
	"fmt"
	"github.com/kanbara/aoc2020/input"
	"strings"
)

func main() {
	data, err := input.ReadFileSplitByEmptyLine(input.DefaultFilename)
	if err != nil {
		panic(err)
	}

	for _, line := range data {
		fmt.Println(strings.ReplaceAll(line, " ", ""))
	}
}
