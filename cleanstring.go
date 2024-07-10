package cleanstring

import (
	"strings"
)

// isLower returns true if the rune is in the ASCII lowercase range.
func isLower(r rune) bool {
    return 'a' <= r && r <= 'z'
}

// isUpper returns true if the rune is in the ASCII uppercase range.
func isUpper(r rune) bool {
    return 'A' <= r && r <= 'Z'
}

// isAlpha returns true if the rune is in the ASCII alphabet range.
func isAlpha(r rune) bool {
    return isLower(r) || isUpper(r)
}

// isNumber returns true if the rune is in the ASCII number range.
func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

// matches returns true if the rune matches any of the test runes.
func matches(r rune, tests []rune) bool {
    for _, test := range tests {
        if r == test {
            return true
        }
    }

    return false
}

// keep returns true if the rune should be kept.
// Kept runes would be replaced but are not.
func keep(r rune) bool {
    return matches(r, []rune{'+'})
}

// skip returns true if the rune should be skipped.
// Skipped runes are not included in the output string and are not replaced.
func skip(r rune) bool {
    return matches(r, []rune{'\''})
}

// toLower lowercases the rune if it is uppercase.
// Any non-uppercase rune will be returned as-is.
func toLower(r rune) rune {
    if !isUpper(r) {
        return r
    }

    delta := r - 'A'
    return 'a' + delta
}

// Format formats a string.
func Format(str string) string {
	var sb strings.Builder

	var R rune

	for _, r := range str {

		if isAlpha(r) {
			// Transform alphabet character to lowercase
			R = toLower(r)
		} else if isNumber(r) || keep(r) {
			// Keep original rune if it is a number
			R = r
		} else if skip(r) || R == '_' {
			// Do not include multiple unbroken underscores
			continue
		} else {
            // Replace character with an underscore
			R = '_'
		}

		sb.WriteRune(R)
	}

	result := sb.String()
	trimmed := strings.Trim(result, "_")

	return trimmed
}

// join joins multiple strings.
func join(strs []string) string {
    return strings.Join(strs, "+")
}

// Join formats and joins strings.
func Join(strs []string) string {
	formatted := make([]string, len(strs))

	for i, str := range strs {
		formatted[i] = Format(str)
	}

	return join(formatted)
}

// associate associates two strings.
func associate(k, v string) string {
    return k + "-" + v
}

// Associate formats, associates, and joins strings.
func Associate(strs map[string]string) string {
	formatted := make([]string, len(strs))

	i := 0

	for k, v := range strs {
        K := Format(k)
        V := Format(v)
		formatted[i] = associate(K, V)
		i++
	}

	return join(formatted)
}
