package main

import (
	"fmt"
	"github.com/kanbara/aoc2020/input"
	"regexp"
	"strconv"
)

type bag struct {
	name string
	count int
}

const shinyGold = "shiny gold"

func addBags(bags []string) map[string][]bag {
	contains := map[string][]bag{}

	for _, line := range bags {

		// two capturing groups: 1 is the num of bags inside, 2 is the bag name,
		// and 3 is the name without `bag`
		re := regexp.MustCompile(`(^|\d+) ?((.*?) bag[s]?)+?`)

		matches := re.FindAllStringSubmatch(line, -1)

		key := matches[0][3]
		fmt.Printf("For bag `%v`\n", key)

		// no other bags inside
		if len(matches) == 1 {
			fmt.Printf("-> no bags inside\n\n")
			contains[key] = nil
			continue
		}

		contains[key] = make([]bag, len(matches[1:]))
		for i, m := range matches[1:] { // skip first match
			fmt.Printf("-> adding %v bags of `%v`\n", m[1], m[3])
			count, err := strconv.Atoi(m[1])
			if err != nil {
				panic(err)
			}

			contains[key][i] = bag{name: m[3], count:count}
		}

		fmt.Printf("\n")
	}

	return contains
}

func getNumberOfBags(contains map[string][]bag, bags []bag) int {
	c := 0
	if len(bags) != 0 { // just to not overly print garbage
		fmt.Printf("iteration with %v bags: %v\n", len(bags), bags)

		for k := range bags {
			c += getNumberOfBags(contains, contains[bags[k].name]) // recurse through each bag
			// looking up the entries to iterate over as hashmap keyed on the name
			// which reveals the sub-bags for each bag
		}

		return c + countBags(bags, shinyGold) // return the results upwards
	}

	return 0
}

func countBags(bags []bag, bagToFind string) int {
	for i := range bags {
		if bags[i].name == bagToFind {
			return 1
		}
	}

	return 0
}

func main() {
	data, err := input.ReadFileAsStrings(input.DefaultFilename)
	if err != nil {
		panic(err)
	}

	bags := addBags(data)

	num := 0
	for k := range bags {
		if k == shinyGold {
			continue // we don't consider ourself here
		}

		fmt.Printf("%v\n", k)
		if getNumberOfBags(bags, bags[k]) > 0 {
			fmt.Printf("->could contain %v\n", shinyGold)
			num++
		}

		fmt.Println()
	}

	fmt.Printf("%v bag colours can contain `%v`\n", num, shinyGold)
}