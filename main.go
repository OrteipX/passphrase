package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	alphaUpperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaLowerChars = "abcdefghijklmnopqrstuvwxyz"
	numericChars    = "0123456789"
	specialChars    = "!#$%&()*+-=?@[]^{}"
)

var (
	passLen     uint
	alpha       bool
	alphaUpper  bool
	alphaLower  bool
	numeric     bool
	special     bool
	mixAll      bool
	exclude     string
	includeOnly string
)

func init() {
	flag.UintVar(&passLen, "l", 0, "Password Length")
	flag.BoolVar(&alpha, "a", false, "Include alphanumeric random chars (e.g.: [A-Za-z])")
	flag.BoolVar(&alphaUpper, "au", false, "Include alphanumeric uppercase chars (e.g.: [A-Z])")
	flag.BoolVar(&alphaLower, "al", false, "Include alphanumeric lowercase chars (e.g.: [a-z])")
	flag.BoolVar(&numeric, "n", false, "Include numeric chars (e.g.: [0-9])")
	flag.BoolVar(&special, "s", false, "Include special characters")
	flag.BoolVar(&mixAll, "x", false, "Include a mix of all types (alpha, numeric, special)")
	flag.StringVar(&exclude, "e", "", "Exclude characters from the password gen")
	flag.StringVar(&includeOnly, "i", "", "Include only the characters to the password gen")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s -l [len] [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
}

type Password struct {
	rand *rand.Rand
}

func (p Password) removeIgnoredChars(base, ignore string) string {
	if ignore == "" {
		return base
	}

	ignoreSet := make(map[rune]struct{}, len(ignore))

	for _, c := range ignore {
		ignoreSet[c] = struct{}{}
	}

	var builder strings.Builder

	builder.Grow(len(base))

	for _, c := range base {
		if _, exists := ignoreSet[c]; !exists {
			builder.WriteRune(c)
		}
	}

	return builder.String()
}

func (p Password) generate() string {
	var possibleChars string

	if includeOnly != "" {
		possibleChars = includeOnly
	} else {
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

		if exclude != "" {
			possibleChars = p.removeIgnoredChars(possibleChars, exclude)
		}
	}

	if len(possibleChars) == 0 {
		fmt.Println("No valid character sets chosen for password generation.")

		os.Exit(1)
	}

	var password bytes.Buffer

	for i := uint(0); i < passLen; i++ {
		c := possibleChars[p.rand.Intn(len(possibleChars))]

		password.WriteByte(c)
	}

	return password.String()
}

func main() {
	flag.Usage = usage

	flag.Parse()

	if passLen == 0 || (includeOnly == "" && !alpha && !alphaUpper && !alphaLower && !numeric && !special && !mixAll) {
		flag.Usage()

		os.Exit(1)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	password := Password{r}

	fmt.Println(password.generate())
}
