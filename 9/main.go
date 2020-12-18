package main

import (
	"fmt"
	"github.com/kanbara/aoc2020/input"
	"math"
)

func main() {
	data, err := input.ReadFileAsInts(input.DefaultFilename)
	if err != nil {
		panic(err)
	}


	window := 25

	invalid := getInvalidNumber(window, data)
	if invalid != nil {
		fmt.Printf("%v has no sum in the previous %v numbers!\n", *invalid , window)

		minmax := findContiguousSum(*invalid, data)
		if minmax != nil {
			fmt.Printf("contiguous minmax sum: %v\n", *minmax)
		}
	}
}

func findContiguousSum(num int, data []int) *int {
	sum := 0
	var tried []int

	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			sum += data[j]
			tried = append(tried, data[j])

			if sum > num {
				break
			}

			if sum == num {
				fmt.Printf("tried: %v\n", tried)
				min := math.MaxInt64
				max := 0

				for _, n := range tried {
					if n < min {
						min = n
					}

					if n > max {
						max = n
					}
				}

				minmax := min + max
				fmt.Printf("got min %v, max %v\n", min, max)
				return &minmax
			}
		}

		sum = 0
		tried = []int{}
	}

	return nil
}

func getInvalidNumber(window int, data []int) *int {
	for i := window; i < len(data); i++ {
		r := combine(data[i-window : i])
		found := false
		for _, v := range r {
			if v == data[i] {
				found = true
			}
		}

		if !found {
			return &data[i]
		}
	}

	return nil
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


