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

	for i := range data {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == 2020 {
				fmt.Printf("found %v %v: %v\n", data[i], data[j], data[i]*data[j])
			}
		}
	}

	for i := range data {
		for j := i + 1; j < len(data); j++ {
			for k := j + 1; k < len(data)-1; k++ {
				if data[i]+data[j]+data[k] == 2020 {
					fmt.Printf("found %v %v %v: %v\n", data[i], data[j], data[k],
						data[i]*data[j]*data[k])
				}
			}
		}
	}
}