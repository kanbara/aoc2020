package main

import (
	"fmt"
	"github.com/kanbara/aoc2020/input"
	"strings"
)

func main() {
	data, err := input.ReadFileAsStrings(input.DefaultFilename)
	if err != nil {
		panic(err)
	}

	maps := make([]map[rune]int, len(data))
	sumAny := 0
	sumEveryone := 0

	for i := 0; i < len(data); i++ {
		str := ""
		total := 0
		for i < len(data) {
			if data[i] == "" {
				break
			} else {
				str += data[i]
				total++

				if i+1 != len(data) {
					i++
				} else {
					break
				}
			}
		}

		trimmed := strings.ReplaceAll(str, " ", "")
		trimmed = strings.ReplaceAll(trimmed, "\n", "")
		fmt.Printf("%v", trimmed)

		maps[i] = make(map[rune]int)
		for _, char := range trimmed {
			maps[i][char]++
		}

		fmt.Printf(" has %v unique answers\n", len(maps[i]))
		sumAny += len(maps[i])

		for _, v := range maps[i] {
			if v == total {
				sumEveryone++
			}
		}
	}

	fmt.Printf("total unique answers %v\n", sumAny)
	fmt.Printf("total which everyone in a group answered %v\n", sumEveryone)

}
