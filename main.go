package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	passLen    uint
	alpha      bool
	alphaUpper bool
	alphaLower bool
	numeric    bool
	special    bool
	mixAll     bool

	alphaUpperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaLowerChars = "abcdefghijklmnopqrstuvwxyz"
	numericChars    = "0123456789"
	specialChars    = "!#$%&()*+-.:;=?@[]^{}"
)

func init() {
	flag.UintVar(&passLen, "l", 0, "Password Length")
	flag.BoolVar(&alpha, "a", false, "Include alphanumeric random chars (e.g.: [A-Za-z])")
	flag.BoolVar(&alphaUpper, "au", false, "Include alphanumeric uppercase chars (e.g.: [A-Z])")
	flag.BoolVar(&alphaLower, "al", false, "Include alphanumeric lowercase chars (e.g.: [a-z])")
	flag.BoolVar(&numeric, "n", false, "Include numeric chars (e.g.: [0-9])")
	flag.BoolVar(&special, "s", false, "Include special characters")
	flag.BoolVar(&mixAll, "x", false, "Include a mix of all types (alpha, numeric, special)")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s -l [len] [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func generatePassword() string {
	var possibleChars string

	if mixAll {
		possibleChars = alphaUpperChars + alphaLowerChars + numericChars + specialChars
	} else {
		if alpha {
			possibleChars += alphaUpperChars + alphaLowerChars
		} else {
			if alphaUpper {
				possibleChars += alphaUpperChars
			}
			if alphaLower {
				possibleChars += alphaLowerChars
			}
		}
		if numeric {
			possibleChars += numericChars
		}
		if special {
			possibleChars += specialChars
		}
	}

	if len(possibleChars) == 0 {
		fmt.Println("No valid character sets chosen for password generation.")
		os.Exit(1)
	}

	var password bytes.Buffer

	for i := uint(0); i < passLen; i++ {
		c := possibleChars[rand.Intn(len(possibleChars))]
		password.WriteByte(c)
	}

	return password.String()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if passLen == 0 || (!alpha && !alphaUpper && !alphaLower && !numeric && !special && !mixAll) {
		flag.Usage()
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	fmt.Println(generatePassword())
}
