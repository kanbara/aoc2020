package input

import (
	"bufio"
	"os"
	"strconv"
)

const DefaultFilename = "input.txt"

func ReadFileAsStrings(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}

	b := bufio.NewScanner(f)

	var out []string
	for b.Scan() {
		out = append(out, b.Text())
	}

	return out, nil
}

func ReadFileAsInts(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []int{}, err
	}

	b := bufio.NewScanner(f)

	var out []int
	for b.Scan() {
		i, err := strconv.Atoi(b.Text())
		if err != nil {
			return []int{}, err
		}

		out = append(out, i)
	}

	return out, nil
}