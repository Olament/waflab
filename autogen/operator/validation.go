package operator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/waflab/waflab/autogen/utils"
)

const (
	byteRangeStringLength = 10
)

type byteRange struct {
	lower byte
	upper byte
}

/*
This function assume that argument follows the format: <number>, <range>, <number>.
In addition to that, the function assume that the range are non-overlapping. Having an overlapping range
in the argument will cause the the function yield a byte string where overlapped string byte have an increasing chance
to show up.
*/
func reverseValidateByteRange(argument string, not bool) (string, error) {
	var byteRanges []byteRange
	size := 0 // number of byte in byteRanges

	// construct byteRanges
	byteRangesString := strings.Split(argument, ",")
	for _, byteRangeString := range byteRangesString {
		byteRangeString = strings.TrimSpace(byteRangeString)
		if strings.Contains(byteRangeString, "-") {
			// a range, Ex: 10-20
			parts := strings.Split(byteRangeString, "-")
			lower, err := strconv.Atoi(parts[0])
			upper, err := strconv.Atoi(parts[1])
			if err != nil {
				return "", err
			}
			byteRanges = append(byteRanges, byteRange{byte(lower), byte(upper)})
			size = size + (upper - lower + 1)
		} else {
			// a number: Ex: 1
			num, err := strconv.Atoi(byteRangeString)
			if err != nil {
				return "", err
			}
			byteRanges = append(byteRanges, byteRange{byte(num), byte(num)})
			size = size + 1
		}
	}

	// build string from byte range
	var build strings.Builder
	for i := 0; i < byteRangeStringLength; i++ {
		num := utils.RandomIntWithRange(0, size)
		for _, r := range byteRanges {
			if (r.upper - r.lower) < byte(num) {
				continue
			}
			build.WriteByte(r.lower + byte(num))
		}
	}

	return build.String(), nil
}

// @validateUtf8Encoding return true if the input string is not a validate ut8-encoded string.
func reverseValidateUtf8Encoding(argument string, not bool) (string, error) {
	// \xFF\xFE is an invalid utf8 header
	return fmt.Sprintf("\xFF\xFE%s", utils.RandomString(10)), nil
}

// @validateUrlEncoding return true if the input string is not a validate url-encoded string.
// A string is not a validate URL-encoding string if
// 1. Not enough byte. Ex. "%", "%1"
// 2. Non-hexadecimal character used. Ex. "%1Z"
func reverseValidateURLEncoding(argument string, not bool) (string, error) {
	res := "%"
	for i := 0; i < 2; i++ {
		if utils.RandomBiasedBool(0.5) { // generate string w/ not enough byte
			return res, nil
		}
		// concat [G-Z], a non-hexadecimal characters, to res
		// there is no difference between [G-Z] and [g-z] since url encoding is case-insensitive
		res += string(int32(utils.RandomIntWithRange(71, 91)))
	}
	return res, nil
}
