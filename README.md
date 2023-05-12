# Clean String

A string-parsing library for creating permutations of strings without spaces, Unicode, and other bloat. Creates string representations of slices and maps.

## Rules
 - Alphabetical characters become lowercase
 - Numeric characters are preserved
 - All other characters[^1] become underscores
 - Consecutive underscores are removed

[^1]: Since the plus sign character (`+`) and the minus sign character (`-`) are used to represent lists and maps, they are exceptions to this rule.

## Usage

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

## Documentation

```go
package cleanstring // import "github.com/haydenheroux/cleanstring"
```

**FUNCTIONS**

```
func Clean(str string) string
    Clean returns a string which is "cleanly formatted". Alphanumeric characters
    are generally maintained. Punctuation is substituted for underscores.

func CleanMap(strs map[string]string) string
    CleanMap returns the result of cleanly-formatting then joining a map of
    strings. Each key and value in the slice is first cleanly-formatted.
    All key-value pairs are then joined using "-". All pairs are then joined
    using "+".

func CleanSlice(strs []string) string
    CleanSlice returns the result of cleanly-formatting then joining a slice of
    strings. Each string in the slice is first cleanly-formatted. All strings
    are then joined using "+".
```
