package cleanstring

import (
  "strings"
)

func isAlpha(r rune) bool {
  isUppercase := 'A' <= r && r <= 'Z'
  isLowercase := 'a' <= r && r <= 'z' 
  return isUppercase || isLowercase
}

func isNumber(r rune) bool {
  return '0' <= r && r <= '9'
}

func isAtHeadOrTail(i int, str string) bool {
  return i == 0 || i == len(str)-1
}

func isInBody(i int, str string) bool {
  return !isAtHeadOrTail(i, str)
}

func toLower(r rune) rune {
  lower := strings.ToLower(string(r))[0]
  return rune(lower)
}

func Clean(str string) string {
  var sb strings.Builder  

  var newR rune
  var prevR rune

  for i, r := range str {

    if isAlpha(r) {
      // Transform alphabet character to lowercase
      newR = toLower(r)
    } else if isNumber(r) {
      // Keep original rune if it is a number
      newR = r
    // All branches from here are non-alphanum cases
    } else if isInBody(i, str) && prevR != '_' {
      // Include underscore if underscore was not the previous rune
      newR = '_'
    } else {
      // Do not include multiple unbroken underscores, leaving one underscore
      continue
    }

    prevR = newR

    sb.WriteRune(newR)
  }

  newStr := sb.String()
  trimmed := strings.Trim(newStr, "_")

  return trimmed

}
