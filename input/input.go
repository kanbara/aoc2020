package input

import (
	"bufio"
	"bytes"
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

func ReadFileSplitByEmptyLine(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	// This split function passes double-newline-delimited tokens, which allows
	// us to get the entire string for each line automatically


	// Set the split function for the scanning operation.
	scanner.Split(splitByEmptyLine)
	// Validate the input

	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func splitByEmptyLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte{'\n', '\n'}); i >= 0 {
		// we don't want any newlines in our output, so the token will be
		// the entire keyval

		// trim the end and replace the inners with spaces
		// f.e.
		// a
		// b\n
		// c\n
		// \n
		//
		// -> a b c, not a bc
		trimmed := bytes.TrimSpace(data[0:i+1])
		fixed := bytes.ReplaceAll(trimmed, []byte{'\n'}, []byte{' '})
		// we still return i+1 bytes here as we need to consume that
		// much of the original buffer
		return i+1, fixed, nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		trimmed := bytes.TrimSpace(data)
		fixed := bytes.ReplaceAll(trimmed, []byte{'\n'}, []byte{' '})
		return len(data), fixed, nil
	}

	// Request more data.
	return 0, nil, nil
}
