package main

import (
	"fmt"
	"github.com/kanbara/aoc2020/input"
	"strconv"
	"strings"
)

const (
	nop op = "nop"
	acc op = "acc"
	jmp op = "jmp"
)

type op string

type operation struct {
	code op
	val int16
	visited bool
}

func main() {
	data, err := input.ReadFileAsStrings(input.DefaultFilename)
	if err != nil {
		panic(err)
	}

	operations := make([]operation, len(data))

	for i := range data {
		vals := strings.Split(data[i], " ")
		val, err := strconv.ParseInt(vals[1], 10, 16)
		if err != nil {
			panic(err)
		}

		operations[i] = operation{
			code:    op(vals[0]),
			val:     int16(val),
			visited: false,
		}
	}
	
	fmt.Printf("accumulator just before infinite loop: %v\n", run(operations, infloop))
}

//func mutate(operations []operation) int16 {
//
//}

func infloop(operations []operation, i int16) bool {
	// infinite loop case, break
	if operations[i].visited {
		fmt.Printf("program will repeat at %v, (%v %+d)\n",
			i, operations[i].code, operations[i].val)
		return true
	}

	return false
}

func outofbounds(operations[] operation, i int16) bool {
	// will we jump out of bounds of our stack?
	if i > int16(len(operations)) {
		return true
	}

	return false
}

func run(operations []operation, endcondition func(operations[] operation, i int16) bool) int16 {
	var accval int16
	var i int16

	for {
		if endcondition(operations, i) {
			return accval
		}

		var t int16 // tmp var here to reduce visited logic triplication
		switch operations[i].code {
		case nop:
			t = i+1 // nop, just move the pc
			fmt.Printf("noop, PC->%v\n", t)
		case acc:
			accval += operations[i].val
			t = i+1 // move pc
			fmt.Printf("added %v to acc, PC->%v\n", operations[i].val, t)
		case jmp:
			t = i + operations[i].val // jump by val (+ or -)
			fmt.Printf("jumping PC->%v\n", t)
		default:
			panic(fmt.Sprintf("got unknown operation: %v\n", operations[i].code))
		}

		operations[i].visited = true
		i = t
	}

	return accval
}