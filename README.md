# cleanstring

Library for cleanly-formatted permutations of strings.

## rules

Alphanumeric characters are largely unaltered.
 - Alphabet characters are transformed to be lowercase only
 - Numerals remain as their original value

Punctuation is replaced by underscores.
 - Punctuation at either the head or last is removed
 - All other non-alphanumeric characters become underscores
 - Multiple unbroken underscores become one underscore

## usage

```go
str := "The quick brown fox jumps over the lazy dog."
// Prints "the_quick_brown_fox_jumps_over_the_lazy_dog"
fmt.Println(cleanstring.Clean(str))
```

```go
strs := []string{"Two", "Three", "Five", "Seven", "Eleven"}
// Prints "two+three+five+seven+eleven"
fmt.Println(cleanstring.CleanSlice(strs))
```

```go
strMap := map[string]string{"1": "one", "2": "two", "3": "three"}
// Prints "1-one+2-two+3-three"
fmt.Println(cleanstring.CleanMap(strMap))
```
