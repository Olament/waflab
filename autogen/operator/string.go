// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package operator

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/waflab/waflab/autogen/utils"
)

// WorkingDirectory for pmFromFile operator
var WorkingDirectory string

func reverseRx(argument string, not bool, flag int) (string, error) {
	str, err := GenerateStringFromRegex(argument, not, flag)
	if err != nil {
		return "", err
	}
	return str, nil
}

func reverseBeginsWith(argument string, not bool) (string, error) {
	return reverseRx(fmt.Sprintf("^%s.*", regexp.QuoteMeta(argument)), not, NoFlag)
}

func reverseContains(argument string, not bool) (string, error) {
	return reverseRx(fmt.Sprintf(".*%s.*", regexp.QuoteMeta(argument)), not, NoFlag)
}

func reverseEndsWith(argument string, not bool) (string, error) {
	return reverseRx(fmt.Sprintf(".*%s$", regexp.QuoteMeta(argument)), not, NoFlag)
}

/*
	Since snort content style use | to note the entrance and exit of embedded binary data, we can split phrase
	by separator | and the non-binary and binary will appear in alternating pattern
	Ex:
	"A|41|A" -> ["A", "41", "A"]
	"|41|A" -> ["", "41", "A"]
*/
func convertSnortStyle(phrase string) (string, error) {
	var builder strings.Builder
	isBinary := false
	for _, part := range strings.Split(phrase, "|") {
		if isBinary {
			decoded, err := hex.DecodeString(part)
			if err != nil {
				return "", err
			}
			builder.Write(decoded)
		} else {
			builder.WriteString(part)
		}
		isBinary = !isBinary
	}
	return builder.String(), nil
}

func reversePm(argument string, not bool) (string, error) {
	phrases := strings.Split(argument, " ")
	phrase := phrases[utils.RandomIntWithRange(0, len(phrases))] // pick phrase from pm's parameters randomly

	return convertSnortStyle(phrase)
}

func reversePmFromFile(argument string, not bool) (string, error) {
	pth := argument
	if !filepath.IsAbs(argument) {
		pth = path.Join(WorkingDirectory, pth)
	}
	file, err := os.Open(pth)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	phrases := []string{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// remove comment marked by #
		parts := strings.Split(line, "#")
		if len(parts[0]) > 0 { // contains non-comment string
			phrases = append(phrases, parts[0])
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return convertSnortStyle(phrases[utils.RandomIntWithRange(0, len(phrases))])
}

func reverseStrEq(argument string, not bool) (string, error) {
	return argument, nil
}

func reverseWithin(argument string, not bool) (string, error) {
	return reverseRx(fmt.Sprintf(".*%s.*", regexp.QuoteMeta(argument)), not, NoFlag)
}
