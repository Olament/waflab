package payload

import (
	"net/url"
	"strings"
)

const htmlCharacters = "!#$'()*+,/:;=?@[]-_.~%"

func httpRequestEscape(content string) string {
	var builder strings.Builder
	for _, r := range content {
		switch r {
		case '\b':
			builder.WriteString(`\b`)
		case '\f':
			builder.WriteString(`\f`)
		case '\n':
			builder.WriteString(`\n`)
		case '\r':
			builder.WriteString(`\r`)
		case '\t':
			builder.WriteString(`\t`)
		case '"':
			builder.WriteString(`\"`)
		case '\'':
			builder.WriteString(`\'`)
		case '\\':
			builder.WriteString(`\\`)
		default:
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// isURLEncoded check if a string is url-encoded by comparing the
// url-decoded value to its original value
func isURLEncoded(value string) bool {
	decoded, err := url.QueryUnescape(value)
	if err != nil {
		return false
	}
	if decoded != value {
		return true
	}
	return false
}

func isValidURL(value string) bool {
	for _, r := range value {
		// check if rune is one of the reserved or unreserved html characters
		if !((65 <= r && r <= 90) || (97 <= r && r <= 122) ||
			(48 <= r && r <= 57) || strings.Contains(htmlCharacters, string(r))) {
			return false
		}
	}
	return true
}
