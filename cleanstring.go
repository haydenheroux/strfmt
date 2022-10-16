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

func toLower(r rune) rune {
  lower := strings.ToLower(string(r))[0]
  return rune(lower)
}

func Clean(str string) string {
  var sb strings.Builder  

  for i, r := range str {

    var newR rune

    if isAlpha(r) {
      newR = toLower(r)
    } else if isNumber(r) {
      newR = r
    } else if i == 0 || i == len(str)-1 {
      continue
    } else {
      newR = '_'
    }

    sb.WriteRune(newR)
  }

  return sb.String()

}
