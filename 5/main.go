package main

import (
	"fmt"
	"github.com/kanbara/aoc2020/input"
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
	var mine int

	for i := range IDs {

		if i <= 0 {
			continue
		}

		if i >= len(IDs)-1 {
			break
		}
		current := IDs[i]
		diff := current - IDs[i+1]

		switch diff {
		case -2:
			mine = current+1
			break
		case -1:
			continue
		case +1:
			continue
		case +2:
			mine = current-1
			break
		}
	}

	if mine == len(IDs) {
		panic("couldn't find seat\n")
	}
	fmt.Printf("\nFound my seat, ID %v\n", mine)
}
