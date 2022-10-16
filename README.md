# cleanstring

Library for cleanly-formatted permutations of strings.

## rules

Alphanumeric characters are largely unaltered.
 - Alphabet characters are transformed to be lowercase only
 - Numerals remain as their original value

Punctuation is replaced by underscores.
 - Punctuation at either the head or tail is removed
 - All other non-alphanumeric characters become underscores
 - Multiple unbroken underscores become one underscore

## examples

```go
str := "The quick brown fox jumps over the lazy dog."
// Prints "the_quick_brown_fox_jumps_over_the_lazy_dog"
fmt.Println(cleanstring.Clean(str))
```
