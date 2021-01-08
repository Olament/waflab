package transformer

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"unicode"

	"github.com/waflab/waflab/autogen/utils"
)

// probability list for reversing transformation
const (
	randomStringLength        = 10
	reverseLowerCaseProb      = 0.5
	reverseCompressProb       = 0.5
	reverseCSSDecodeProb      = 0.50
	reverseHTMLEntityProb     = 0.50
	reverseCommentProb        = 0.10
	reverseCommentCharProb    = 0.10
	reverseNullProb           = 0.10
	reverseReplaceCommentProb = 0.5
	reverseReplaceNullProb    = 0.5
)

var whiteSpaceCharacters = []string{"\f", "\t", "\n", "\r", "\v"}

// randomStringsInsertion randomly insert string from reserve to str.
// At each rune of str, randomstringsInsertion will randomly pick a string from reserve and
// insert it between the rune with given probability
func randomStringsInsertion(str string, reserve []string, probability float64) string {
	var builder strings.Builder
	for _, r := range str {
		builder.WriteRune(r)
		if utils.RandomBiasedBool(probability) {
			builder.WriteString(reserve[utils.RandomIntWithRange(0, len(reserve))])
		}
	}
	return builder.String()
}

func rune2HexString(r rune) string {
	return hex.EncodeToString([]byte{byte(r)})
}

func reverseBase64Decode(variable string) string {
	return base64.StdEncoding.EncodeToString([]byte(variable))
}

// ModSecurity encode characters using CSS 2.x escape rules where each unicode character is
// represented by a blackslash folloed by up to six hexadecimal characters.
func reverseCSSDecode(variable string) string {
	var builder strings.Builder
	for _, r := range variable {
		if utils.RandomBiasedBool(reverseCSSDecodeProb) {
			builder.WriteString(fmt.Sprintf("\\%06s", rune2HexString(r)))
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// reverseCompressWhiteSpace assume that the only kinds of whitespace character
// variable contains is space and there is not any consecutive space.
func reverseCompressWhiteSpace(variable string) string {
	var builder strings.Builder

	for _, r := range variable {
		if unicode.IsSpace(r) {
			for p := true; p; p = utils.RandomBiasedBool(reverseCompressProb) {
				builder.WriteString(utils.PickRandomString(whiteSpaceCharacters))
			}
		}
		builder.WriteRune(r)
	}

	return builder.String()
}

func reverseHexDecode(variable string) string {
	return hex.EncodeToString([]byte(variable))
}

// reverseHTMLEntityDecode encode the variable
func reverseHTMLEntityDecode(variable string) string {
	// from https://golang.org/src/html/escape.go
	htmlEscaper := map[string]string{
		`&`:    "&amp;",
		`'`:    "&#39;",
		`<`:    "&lt;",
		`>`:    "&gt;",
		`"`:    "&#34;",
		"\xa0": "&nbsp;",
	}

	var builder strings.Builder
	for _, r := range variable {
		if utils.RandomBiasedBool(reverseHTMLEntityProb) {
			if s, okay := htmlEscaper[string(r)]; okay { // special html character
				builder.WriteString(s)
			} else {
				if utils.RandomBiasedBool(0.50) {
					builder.WriteString(fmt.Sprintf("&#x%2s", rune2HexString(r))) // &#xHH, hexadecimal
				} else {
					builder.WriteString(fmt.Sprintf("&#%03d", r)) // &#DDD decimal number
				}
			}
		} else { // not encode
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func reverseLength(variable string) string {
	length, err := strconv.Atoi(variable)
	if err != nil {
		panic("Length must be an integer")
	}
	return utils.RandomString(length)
}

// reverseNormalizePath assume that the variable string is a normalized path,
// otherwise the function may return illegal path
func reverseNormalizePath(variable string) string {
	redundantPath := []string{"", ".", "foo/.."}
	parts := strings.Split(variable, "/")
	res := []string{}

	for index, part := range parts {
		res = append(res, part)
		if index < len(parts)-1 { // add redundant path in between the path
			res = append(res, utils.PickRandomString(redundantPath))
		}
	}

	return strings.Join(res, "/")
}

func reverseNormalizePathWin(variable string) string {
	return strings.ReplaceAll(reverseNormalizePath(variable), "/", "\\")
}

func reverseLowercase(variable string) string {
	var builder strings.Builder
	for _, char := range variable {
		if utils.RandomBiasedBool(reverseLowerCaseProb) {
			builder.WriteRune(unicode.ToUpper(char))
		} else {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

func reverseRemoveComments(variable string) string {
	res := randomStringsInsertion(variable,
		[]string{fmt.Sprintf("/*%s*/", utils.RandomString(randomStringLength))},
		reverseCommentProb)
	res += fmt.Sprintf("#%s", utils.RandomString(randomStringLength))
	return res
}

func reverseRemoveCommentsChar(variable string) string {
	return randomStringsInsertion(variable, []string{"/**/", "--", "#"}, reverseCommentCharProb)
}

func reverseReplaceComments(variable string) string {
	var builder strings.Builder
	for _, r := range variable {
		if unicode.IsSpace(r) && utils.RandomBiasedBool(reverseReplaceCommentProb) {
			builder.WriteString(fmt.Sprintf("/*%s*/", utils.RandomString(randomStringLength)))
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func reverseRemoveNulls(variable string) string {
	return randomStringsInsertion(variable, []string{"\x00"}, reverseNullProb)
}

func reverseReplaceNulls(variable string) string {
	var builder strings.Builder
	for _, r := range variable {
		if unicode.IsSpace(r) && utils.RandomBiasedBool(reverseReplaceNullProb) {
			builder.WriteString("\000")
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func reverseTrim(variable string) string {
	return reverseTrimLeft(reverseTrimRight(variable))
}

func reverseTrimLeft(variable string) string {
	return fmt.Sprintf("%s%s", utils.RandomStringFromSet(randomStringLength, whiteSpaceCharacters), variable)
}

func reverseTrimRight(variable string) string {
	return fmt.Sprintf("%s%s", variable, utils.RandomStringFromSet(randomStringLength, whiteSpaceCharacters))
}

func reverseUrlDecode(variable string) string {
	return url.QueryEscape(variable)
}

// t:utf8toUnicode converts all UTF-8 characters to Unicode using %uHHHH syntax
// Ex: ćat¯’/etç/passwd’ -> %u0107at%u00af%u2019/et%u00e7/passwd%u2019
// To reverse the transformation, we need to convert unicode to utf-8 character
// since Golang does not support u percent format, we need to replace all %u with \u
func reverseUtf8ToUnicode(variable string) string {
	return fmt.Sprintf(strings.ReplaceAll(variable, "%u", "\\u"))
}
