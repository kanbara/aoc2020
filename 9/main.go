package main

import (
	"fmt"
	"github.com/kanbara/aoc2020/input"
)

func main() {
	data, err := input.ReadFileAsInts(input.DefaultFilename)
	if err != nil {
		panic(err)
	}

	window := 25

	for i := window; i < len(data); i++ {
		r := combine(data[i-window:i])
		found := false
		for _, v := range r {
			if v == data[i] {
				found = true
			}
		}

		if !found {
			fmt.Printf("%v has no sum in the previous %v numbers!\n", data[i], window)
			break
		}
	}
}

func combine(nums []int) []int {
	var combinations func(nums []int, curr int, end int)
	var res []int

	combinations = func(nums []int, curr int, end int) {
		if curr > end {
			return
		}

		for _, n := range nums[curr+1:] {
			res = append(res, n+nums[curr])
		}

		combinations(nums, curr+1, end)
	}

	combinations(nums, 0, len(nums)-2)

	var uniq []int
	m := map[int]interface{}{}
	for _, v := range res {
		if _, ok := m[v]; !ok {
			// entry wasn't found yet, add
			uniq = append(uniq, v)
			m[v] = nil
		}
	}

	return uniq
}


