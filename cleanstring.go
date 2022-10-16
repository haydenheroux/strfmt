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

  var R rune

  for _, r := range str {

    if isAlpha(r) {
      // Transform alphabet character to lowercase
      R = toLower(r)
    } else if isNumber(r) {
      // Keep original rune if it is a number
      R = r
    // All branches from here are non-alphanum cases
    } else if R != '_' {
      // Include underscore if underscore was not the previous rune
      R = '_'
    } else {
      // Do not include multiple unbroken underscores, leaving one underscore
      continue
    }

    sb.WriteRune(R)
  }

  newStr := sb.String()
  trimmed := strings.Trim(newStr, "_")

  return trimmed
}

func CleanSlice(strs []string) string {
  newStrs := make([]string, len(strs))
  
  for i, str := range strs {
    newStrs[i] = Clean(str)
  }

  return strings.Join(newStrs, "+")
}

func CleanMap(strs map[string]string) string {
  newStrs := make([]string, len(strs))

  i := 0

  for key, value := range strs {
    cleanKey := Clean(key)
    cleanValue := Clean(value)
    newStrs[i] = cleanKey + "-" + cleanValue
    i++
  }

  return strings.Join(newStrs, "+")
}
