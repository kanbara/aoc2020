package main

import (
	"fmt"
	"github.com/kanbara/aoc2020/input"
	"math"
	"sort"
	"strconv"
)

const (
	F = '0'
	FChar = 'F'
	B = '1'
	BChar = 'B'
	L = '0'
	LChar = 'L'
	R = '1'
	RChar = 'R'
)

func convert(raw string, isRow bool) uint8 {
	str := "0b0"// 7 bytes binary
	for _, char := range raw {
		if isRow {
			if char == FChar {
				str += string(F)
			} else {
				str += string(B)
			}
		} else {
			if char == LChar {
				str += string(L)
			} else {
				str += string(R)
			}
		}
	}

	i, err := strconv.ParseInt(str, 0, 8)
	if err != nil {
		panic(err)
	}

	return uint8(i)
}

func rowSeatAndID(raw string) (row, seat uint8, ID uint16) {
	row = convert(raw[:7], true)
	seat = convert(raw[7:], false)
	ID = (uint16(row) * 8) + uint16(seat)
	fmt.Printf("%v -> row: %03d, seat: %03d | ID: %v\n",
		raw, row, seat, ID)
	return
}

func main() {
	data, err := input.ReadFileAsStrings(input.DefaultFilename)
	if err != nil {
		panic(err)
	}

	var IDs []int

	var highestID uint16
	highestIDStr := ""
	for _, entry := range data {
		_, _, ID := rowSeatAndID(entry)

		if ID > highestID {
			highestID = ID
			highestIDStr = entry
		}

		IDs = append(IDs, int(ID))
	}

	fmt.Printf("\nHighest ID Found\n")
	_, _, _  = rowSeatAndID(highestIDStr)

	sort.Ints(IDs)

	sum := 0
	for i := range IDs {
		sum += IDs[i]
	}

	min := IDs[0]
	max := IDs[len(IDs)-1]

	// min + min+1 + min+2 ... + max-1 + max
	fmt.Printf("total sum with our missing seat from %v->%v: %v\n", min, max, sum)

	// sum of series: n(n+1) / 2
	// 0 + 1 + 2 + ... + n
	fullSum := (max * (max + 1)) / 2
	fmt.Printf("total sum from 0->%v: %v\n", max, fullSum)

	// subtract values we don't have, e.g. 0 + 1 + 2 + ... m
	subsetSum := (min * (min + 1)) / 2

	// if we're odd we'll be off by min
	if min % 2 != 0 {
		subsetSum -= min
	}

	fmt.Printf("removing subset 0->%v: %v\n", min, subsetSum)
	fmt.Printf("\nFound my seat, ID %v\n", math.Abs(float64(sum - fullSum + subsetSum)))
}
