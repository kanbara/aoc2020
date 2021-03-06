package main

import (
	"encoding/hex"
	"fmt"
	"github.com/kanbara/aoc2020/input"
	"strconv"
	"strings"
)

const (
	birthYear key = "byr"
	issueYear key = "iyr"
	expiryYear key = "eyr"
	height key = "hgt"
	hairColour key = "hcl"
	eyeColour key = "ecl"
	passportID key = "pid"
	countryID key = "cid"
)

type key string
type Pass map[key]string

func (p *Pass) Read(raw string) {
	kvs := strings.Split(raw, " ")
	for _, kv := range kvs {
		s := strings.Split(kv, ":")
		(*p)[key(s[0])] = s[1]
	}
}

func (p *Pass) Valid() bool {
	// can we be stupid and just check for 8 fields-- or 7 and only cid missing?

	if len(*p) == 8 {
		fmt.Println("all keys present")
		return true
	}

	if len(*p) == 7 {
		fmt.Printf("7 keys present...")
		if _, ok := (*p)[countryID]; !ok {
			fmt.Println("but cid is missing")
			return true
		}

		fmt.Println("but cid is not the missing key")
	}

	fmt.Printf("only got %v keys\n", len(*p))
	return false
}

//byr (Birth Year) - four digits; at least 1920 and at most 2002.
//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
//hgt (Height) - a number followed by either cm or in:
//If cm, the number must be at least 150 and at most 193.
//If in, the number must be at least 59 and at most 76.
//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
//pid (Passport ID) - a nine-digit number, including leading zeroes.
//cid (Country ID) - ignored, missing or not.
func (p *Pass) FieldsValid() bool {
	fmt.Printf("validating: ")
	valid := true
	for k, v := range *p {
		fmt.Printf("%v...", k)
		v := isFieldValid(k, v)
		valid = valid && v
		if v {
			fmt.Printf("VALID ")
		} else {
			fmt.Printf("INVALID ")
		}
	}

	fmt.Println("")
	return valid
}

func isFieldValid(field key, val string) bool {
	switch field {
	case birthYear:
		i, err := strconv.Atoi(val)
		if err != nil {
			return false
		}

		return len(val) == 4 && i >= 1920 && i <= 2002
 	case issueYear:
		i, err := strconv.Atoi(val)
		if err != nil {
			return false
		}

		return len(val) == 4 && i >= 2010 && i <= 2020
	case expiryYear:
		i, err := strconv.Atoi(val)
		if err != nil {
			return false
		}

		return len(val) == 4 && i >= 2020 && i <= 2030
	case height:
		suf := val[len(val)-2:]
		i, err := strconv.Atoi(val[0:len(val)-2])
		if err != nil {
			return false
		}
		return (suf == "cm" && i >= 150 && i <= 193) ||
			(suf == "in" && i >= 59 && i <= 76)
	case hairColour:
		bang := val[0]
		if bang != '#' {
			return false
		}

		rest := val[1:]
		_, err := hex.DecodeString(rest)
		if err != nil {
			return false
		}

		return true
	case eyeColour:
		return val == "amb" ||
			val == "blu" ||
			val == "brn" ||
			val == "gry" ||
			val == "grn" ||
			val == "hzl" ||
			val == "oth"
	case passportID:
		_, err := strconv.Atoi(val)
		if err != nil {
			return false
		}

		return len(val) == 9
	case countryID:
		return true
	}

	return false
}

func main() {
	data, err := input.ReadFileSplitByEmptyLine(input.DefaultFilename)
	if err != nil {
		panic(err)
	}

	var count uint8
	var fieldsValid uint8
	for _, line := range data {
		p := Pass{}
		p.Read(line)

		fmt.Println(line)

		if p.Valid() {
			count++

			if p.FieldsValid() {
				fieldsValid++
			}
		}

		fmt.Printf("\n")

	}

	fmt.Printf("got %v/%v valid passports\n", count, len(data))
	fmt.Printf("of those %v/%v had valid fields\n", fieldsValid, count)

}
