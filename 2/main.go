package main

import (
	"fmt"
	"github.com/kanbara/aoc2020/input"
	"regexp"
	"strconv"
)

type Validator struct {
	lower uint8
	higher uint8
	char rune
	pass string
}

func NewValidator(lower, higher, char, pass string) *Validator {
	li, err := strconv.Atoi(lower)
	if err != nil {
		panic(err)
	}

	hi, err := strconv.Atoi(higher)
	if err != nil {
		panic(err)
	}

	return &Validator{
		lower:  uint8(li),
		higher: uint8(hi),
		char:   rune(char[0]),
		pass:   pass,
	}
}

func (v *Validator) ValidA() bool {
	var count uint8
	for i := range v.pass {
		if rune(v.pass[i]) == v.char {
			count++
		}
	}

	return count >= v.lower && count <= v.higher
}

func (v* Validator) ValidB() bool {
	var pos1 bool
	var pos2 bool

	// remember to be one-indexed as per the policy description!
	if rune(v.pass[v.lower-1]) == v.char {
		pos1 = true
	}

	if rune(v.pass[v.higher-1]) == v.char {
		pos2 = true
	}

	//fmt.Printf("%v -> %v: %v is %v(%v), %v is %v(%v)\n",
	//	v.pass, string(v.char), v.lower-1, string(v.pass[v.lower-1]), pos1,
	//	v.higher-1, string(v.pass[v.higher-1]), pos2)

	return pos1 != pos2
}

func main() {
	data, err := input.ReadFileAsStrings(input.DefaultFilename)
	if err != nil {
		panic(err)
	}

	// format is num-num letter: password
	regex := regexp.MustCompile(`(\d+)-(\d+) (.): (.*)$`)
	var validA uint64
	var validB uint64
	for _, line := range data {
		matches := regex.FindStringSubmatch(line)
		v := NewValidator(matches[1], matches[2], matches[3], matches[4])
		if v.ValidA() {
			validA++
		}

		if v.ValidB() {
			validB++
		}
	}

	fmt.Printf("found %v valid part A passwords\n", validA)
	fmt.Printf("found %v valid part B passwords\n", validB)

}
