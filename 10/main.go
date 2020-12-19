package main

import (
	"fmt"
	"github.com/kanbara/aoc2020/input"
	"sort"
)

func main() {
	data, err := input.ReadFileAsInts(input.DefaultFilename)
	if err != nil {
		panic(err)
	}

	data = append(data, 0) // start joltage
	sort.Ints(data)
	deviceJoltage := data[len(data)-1]+3
	fmt.Printf("highest joltage: %v\n", deviceJoltage)
	data = append(data, deviceJoltage)

	ones, threes := computeJoltageDiffs(data)
	fmt.Printf("1-jolt difference %v, 3-jolt difference %v: %v\n", ones, threes, ones*threes)

}

func computeJoltageDiffs(data []int) (int, int) {
	var ones int
	var threes int

	for i := 1; i < len(data); i++ {
		switch data[i] - data[i-1] {
		case 1:
			ones++
		case 3:
			threes++
		}
	}

	return ones, threes
}
