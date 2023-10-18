package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

var (
	passLen              uint
	alpha                bool
	alphaUpper           bool
	alphaLower           bool
	numeric              bool
	special              bool
	mixAll               bool
	specialCharsASCIINum = []uint{33, 35, 36, 37, 38, 40, 41, 42, 43, 45, 46, 58, 59, 61, 63, 64, 91, 93, 94, 123, 125}
)

const (
	ASCII_UPPER_START = 65  // A
	ASCII_UPPER_END   = 90  // Z
	ASCII_LOWER_START = 97  // a
	ASCII_LOWER_END   = 122 // z
	ASCII_NUM_MIN     = 48  // 0
	ASCII_NUM_MAX     = 57  // 9
)

func init() {
	flag.UintVar(&passLen, "l", 0, "Password Length")
	flag.BoolVar(&alpha, "a", false, "Include alphanumeric random chars (e.g.: [A-Za-z]")
	flag.BoolVar(&alphaUpper, "au", false, "Include alphanumeric upper case random chars (e.g.: [A~Z])")
	flag.BoolVar(&alphaLower, "al", false, "Include alphanumeric lower case random chars (e.g.: [a~z])")
	flag.BoolVar(&numeric, "n", false, "Include numeric random values (e.g.: [0-9])")
	flag.BoolVar(&special, "s", false, "Include special chars random chars")
	flag.BoolVar(&mixAll, "x", false, "Include a mix of all types possible, alpha uppercase, alpha lowercase, special and numbers")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s -l [len] [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func appendNums(numbers *[]uint, min uint, max uint) {
	for n := min; n <= max; n++ {
		*numbers = append(*numbers, n)
	}
}

func getASCIINumbers() []uint {
	var numbers []uint

	if mixAll {
		alpha = true
		numeric = true
		special = true
	}

	if alpha {
		appendNums(&numbers, ASCII_UPPER_START, ASCII_UPPER_END)
		appendNums(&numbers, ASCII_LOWER_START, ASCII_LOWER_END)
	} else {
		if alphaUpper {
			appendNums(&numbers, ASCII_UPPER_START, ASCII_UPPER_END)
		} else if alphaLower {
			appendNums(&numbers, ASCII_LOWER_START, ASCII_LOWER_END)
		}
	}

	if numeric {
		appendNums(&numbers, ASCII_NUM_MIN, ASCII_NUM_MAX)
	}

	if special {
		numbers = append(numbers, specialCharsASCIINum...)
	}

	return numbers
}

func getRandomNum(numbers []uint) uint {
	randNum := rand.Intn(len(numbers))

	return numbers[randNum]
}

func generatePassword() string {
	possibleNumbers := getASCIINumbers()

	var password string

	for i := uint(0); i < passLen; i++ {
		c := rune(getRandomNum(possibleNumbers))
		password += string(c)
	}

	return password
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if passLen == 0 || (passLen > 0 && !alpha && !alphaUpper && !alphaLower && !numeric && !special && !mixAll) {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(generatePassword())
}
