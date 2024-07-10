# strfmt

A string-formatting library for creating simple, command-line-friendly permutations.

## Rules
 - ASCII characters only
 - Lowercase characters only (numbers are preserved)
 - Other characters are replaced by `_`
 - Muliple `_`s are compressed into a single `_`
 - Strings are joined by `+`s
 - Strings are associated by `-`s

## Usage

```go
str := "The quick brown fox jumps over the lazy dog."
// Prints "the_quick_brown_fox_jumps_over_the_lazy_dog"
fmt.Println(strfmt.Format(str))
```
```go
strs := []string{"Two", "Three", "Five", "Seven", "Eleven"}
// Prints "two+three+five+seven+eleven"
fmt.Println(strfmt.Join(strs))
```

```go
strMap := map[string]string{"1": "one", "2": "two", "3": "three"}
// Prints "1-one+2-two+3-three"
fmt.Println(strfmt.Associate(strMap))
```
