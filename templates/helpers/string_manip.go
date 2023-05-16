package helpers

import (
	"regexp"
	"strings"
	"unicode"
)

// StripNonAlphabeticalRegexp uses a regular expression pattern [^a-zA-Z]+ to match one or more consecutive
// non-alphabetical characters.
// The regexp.MustCompile function compiles the regular expression pattern into a *regexp.Regexp object, and the
// ReplaceAllString method is used to replace all occurrences of the matched pattern with an empty string.
func StripNonAlphabeticalRegexp(input string) string {
	regex := regexp.MustCompile("[^a-zA-Z]+")
	return regex.ReplaceAllString(input, "")
}

// StripNonAlphabeticalUnicode initializes a strings.Builder to build the result string.
// It iterates over each character of the input string using unicode.IsLetter to check if each is an alphabet letter.
// If it is, the character is appended to the strings.Builder.
func StripNonAlphabeticalUnicode(input string) string {
	var result strings.Builder
	for _, char := range input {
		if unicode.IsLetter(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}
