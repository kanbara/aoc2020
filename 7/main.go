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


// value is the current value of the bag containing `bags`
// we are, in effect, computing the total value of the outer bag inside its own recursive subcall
func getTotalBagsInside(contains map[string][]bag, bags []bag, value int) int {
	c := 0
	if len(bags) != 0 { // we have a proper base case now, bags == 0 and not
		for k := range bags {
			inner :=  getTotalBagsInside(contains, contains[bags[k].name], bags[k].count)
			c += inner // recurse through each bag and add the contents
		}

		return c * value + value // return the bag * current bag value, and make sure to add the bags themselves
		// e.g. shiny gold has 2 dark orange and 1 vibrant plum
		// dark orange has 3 bright red and 4 drab purple, and vibrant plum has 1 banana yellow
		// we have then 2 * (3+4) + 2  +   1 * 1 + 1
 	}

	return value
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
	fmt.Printf("%v has %v bags inside\n", shinyGold,
		getTotalBagsInside(bags, bags[shinyGold], 1)-1) // -1 so we don't count our own shiny gold bag
}