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

	maps := make([]map[rune]interface{}, len(data))
	sum := 0
	for i, line := range data {
		trimmed := strings.ReplaceAll(line, " ", "")
		trimmed = strings.ReplaceAll(trimmed, "\n", "")
		fmt.Printf("%v", trimmed)

		maps[i] = make(map[rune]interface{})
		for _, char := range trimmed {
			maps[i][char] = nil
		}

		fmt.Printf(" has %v unique answers\n", len(maps[i]))
		sum += len(maps[i])
	}
	
	fmt.Printf("total unique answers %v\n", sum)
}
