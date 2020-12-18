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

	copied := make([]operation, len(operations))
	copy(copied, operations)
	// need to copy here so we don't overwrite all the visited values
	fmt.Printf("accumulator just before infinite loop: %v\n\n", run(copied, infloop))
	fmt.Printf("accumulator after successful execution: %v\n", mutateAndRun(operations))
}

func mutateAndRun(operations []operation) int16 {
	var mutations int
	// the idea here is that we go through each instruction
	// and for each time we see a jmp or nop we swap them and then add to the list, keeping track
	// of the previous position, and once we've swapped all the operations (as only one needs to be swapped)
	// we run them in goroutines, and collect the accumulated values-- programs which infinite loop
	// will return 0, so we can just ignore it

	var accumulated int16

	for i := 0; i < len(operations); i++ {
		switch operations[i].code {
		case jmp:
			fmt.Printf("Swapping jump at %v\n ", i)
			mutations++
			accumulated += run(swap(operations, i, nop), infloopOrEndNormally)
		case nop:
			fmt.Printf("Swapping nop at %v\n", i)
			mutations++
			accumulated += run(swap(operations, i, jmp), infloopOrEndNormally)
		}
	}

	fmt.Printf("made %v mutations\n", mutations)
	return accumulated
}

func swap(operations []operation, i int, newOp op) []operation {
	c := make([]operation, len(operations))
	copy(c, operations)

	op := c[i]
	op.code = newOp
	c[i] = op

	return c
}

func infloop(operations []operation, i int16, acc int16) (bool, int16) {
	// infinite loop case, break
	if operations[i].visited {
		fmt.Printf("program will repeat at %v, (%v %+d)\n\n",
			i, operations[i].code, operations[i].val)
		return true, acc
	}

	return false, acc
}

func infloopOrEndNormally(operations[] operation, i int16, acc int16) (bool, int16) {
	// will we jump out of bounds of our stack?
	if i >= int16(len(operations)) {
		fmt.Printf("program: %v\n", operations)
		fmt.Printf("program exited normally\n\n")
		return true, acc
	}

	// do we infinite loop here?
	// when we loop forever, just return 0 because we need to discard the value
	if end, _ := infloop(operations, i, acc); end {
		return true, 0
	}

	return false, acc
}

func run(operations []operation,
	endcondition func(operations[] operation, i int16, acc int16) (bool, int16)) int16 {
	var accval int16
	var i int16

	for {
		if end, acc := endcondition(operations, i, accval); end {
			return acc
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