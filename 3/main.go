package main

import (
	"fmt"
	"github.com/kanbara/aoc2020/input"
)

const (
	TREE rune = '#'
)

func traverse(right, down int, data []string) int {
	var trees int
	for i := 0; i < len(data); i+=down {
		// we can determine the position we're in with simple mod, e.g. i * RIGHT % len(line)
		// we need to take care to normalise the position to account for how many times we jump
		// rows, e.g. if we have 0 1 2 3 4 but skip 2 we want data from 0 2 4 but we've only
		// actually moved 0 1 2 times
		if rune(data[i][(i/down * right) % len(data[i])]) == TREE {
			trees++
		}
	}

	return trees
}

func main() {
	data, err := input.ReadFileAsStrings(input.DefaultFilename)
	if err != nil {
		panic(err)
	}

	// part 1
	fmt.Printf("hit %d trees\n\n", traverse(3,1, data))

	cases := [][]int{
		{1,1},
		{3,1},
		{5,1},
		{7,1},
		{1,2},
	}

	res := 1
	for _, c := range cases {
		trees := traverse(c[0], c[1], data)
		fmt.Printf("hit %d trees with slope (%v, %v)\n", trees, c[0], c[1])
		res *= trees
	}

	fmt.Printf("trees total: %v\n", res)
}
