// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// The MIT License (MIT)
// Copyright (C) 2016 Lucas Jones
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package operator

import (
	"errors"
	"log"
	"math"
	"regexp"
	"regexp/syntax"
	"strings"

	"github.com/waflab/waflab/autogen/utils"
)

const (
	maxRetry            = 50
	repeatedstringLimit = 10
	printableChars      = "!\"#$%&'()*+,-./0123456789:<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	runeRangeEnd        = 0x10ffff
)

var ErrFailedGeneration = errors.New("autogen/operator: Unable to generation string from regexp")

func generate(re *syntax.Regexp) []rune {
	stack := []*syntax.Regexp{}
	stack = append(stack, re)
	res := []rune{}

	var prevOP syntax.Op = syntax.OpNoMatch

	for len(stack) > 0 {
		n := len(stack) - 1
		re := stack[n]
		stack = stack[:n]

		op := re.Op
		switch op {
		case syntax.OpNoMatch:
		case syntax.OpEmptyMatch:
		case syntax.OpLiteral:
			res = append(res, re.Rune...)
			prevOP = syntax.OpLiteral
		case syntax.OpCharClass:
			// if previous OP is \b and current OP is \s
			if prevOP == syntax.OpWordBoundary && equal(re.Rune, []rune{9, 10, 12, 13, 32, 32}) {
				continue
			}
			// number of possible chars
			sum := 0
			for i := 0; i < len(re.Rune); i += 2 {
				sum += int(re.Rune[i+1]-re.Rune[i]) + 1
				if re.Rune[i+1] == runeRangeEnd {
					sum = -1
					break
				}
			}
			// pick random char in range (inverse match group)
			if sum == -1 {
				possibleChars := []uint8{}
				for j := 0; j < len(printableChars); j++ {
					c := printableChars[j]
					//fmt.Printf("Char %c %d\n", c, c)
					// Check c in range
					for i := 0; i < len(re.Rune); i += 2 {
						if rune(c) >= re.Rune[i] && rune(c) <= re.Rune[i+1] {
							possibleChars = append(possibleChars, c)
							break
						}
					}
				}
				if len(possibleChars) > 0 {
					c := possibleChars[utils.RandomIntWithRange(0, len(possibleChars))]
					return []rune(string([]byte{c}))
				}
			}
			r := utils.RandomIntWithRange(0, int(sum))
			var ru rune
			sum = 0
			for i := 0; i < len(re.Rune); i += 2 {
				gap := int(re.Rune[i+1]-re.Rune[i]) + 1
				if sum+gap > r {
					ru = re.Rune[i] + rune(r-sum)
					break
				}
				sum += gap
			}
			res = append(res, ru)
			prevOP = syntax.OpCharClass
		case syntax.OpAnyCharNotNL, syntax.OpAnyChar:
			chars := printableChars
			if op == syntax.OpAnyCharNotNL {
				chars = printableChars
			}
			c := chars[utils.RandomIntWithRange(0, len(chars))]
			res = append(res, []rune(string([]byte{c}))...)
			prevOP = syntax.OpAnyCharNotNL
		case syntax.OpBeginLine, syntax.OpBeginText:
			prevOP = syntax.OpBeginLine
		case syntax.OpEndLine, syntax.OpEndText:
			prevOP = syntax.OpEndLine
		case syntax.OpWordBoundary:
			if (len(res) == 0 || res[len(res)-1] != 32) && prevOP != syntax.OpBeginLine {
				res = append(res, 32) // rune codepoint for space character
			}
			prevOP = syntax.OpWordBoundary
		case syntax.OpNoWordBoundary:
		case syntax.OpCapture:
			stack = append(stack, re.Sub0[0])
			prevOP = syntax.OpCapture
		case syntax.OpStar:
			// Repeat zero or more times
			count := utils.RandomIntWithRange(0, repeatedstringLimit+1)
			for i := 0; i < count; i++ {
				stack = append(stack, re.Sub...)
			}
			prevOP = syntax.OpStar
		case syntax.OpPlus:
			// Repeat one or more times
			count := utils.RandomIntWithRange(0, repeatedstringLimit) + 1
			for i := 0; i < count; i++ {
				stack = append(stack, reverse(re.Sub)...)
			}
			prevOP = syntax.OpPlus
		case syntax.OpQuest:
			// Zero or one instances
			count := utils.RandomIntWithRange(0, 2)
			for i := 0; i < count; i++ {
				stack = append(stack, reverse(re.Sub)...)
			}
			prevOP = syntax.OpQuest
		case syntax.OpRepeat:
			// Repeat one or more times
			count := 0
			re.Max = int(math.Min(float64(re.Max), float64(repeatedstringLimit)))
			if re.Max > re.Min {
				count = utils.RandomIntWithRange(0, re.Max-re.Min+1)
			}
			for i := 0; i < re.Min || i < (re.Min+count); i++ {
				stack = append(stack, reverse(re.Sub)...)
			}
			prevOP = syntax.OpRepeat
		case syntax.OpConcat:
			// Concatenate sub-regexes
			stack = append(stack, reverse(re.Sub)...)
			prevOP = syntax.OpConcat
		case syntax.OpAlternate:
			i := utils.RandomIntWithRange(0, len(re.Sub))
			stack = append(stack, re.Sub[i])
			prevOP = syntax.OpAlternate
		default:
			log.Fatalf("[reg-gen] Unhandled op: %s", op.String())
		}
	}

	return res
}

func reverse(input []*syntax.Regexp) []*syntax.Regexp {
	n := len(input)
	new := make([]*syntax.Regexp, n)
	for i := 0; i < n; i++ {
		new[i] = input[n-i-1]
	}
	return new
}

func equal(r1, r2 []rune) bool {
	if len(r1) != len(r2) {
		return false
	}
	for index, r := range r1 {
		if r != r2[index] {
			return false
		}
	}
	return true
}

// Generate a negated string from something
func GenerateStringFromRegex(expression string, not bool, flag int) (res string, err error) {
	re, err := syntax.Parse(expression, syntax.PerlX)
	if err != nil {
		return "", err
	}
	regex, err := regexp.Compile(expression)
	if err != nil {
		return "", err
	}
	if not {
		for i := 0; i < maxRetry; i++ {
			res = utils.RandomString(10)
			if regex.MatchString(res) {
				break
			}
		}
		if !regex.MatchString(res) {
			return "", ErrFailedGeneration
		}
	} else {
		rs := generate(re)
		if flag == NoUTF8 { // no UTF8 encoding
			var bs = make([]byte, len(rs))
			for i := 0; i < len(rs); i++ {
				bs[i] = byte(rs[i])
			}
			res = string(bs)
		} else { // default
			res = string(rs)
		}
		replacer := strings.NewReplacer("K", " ", "ſ", " ") // case-folding
		res = replacer.Replace(res)
	}

	return res, nil
}
