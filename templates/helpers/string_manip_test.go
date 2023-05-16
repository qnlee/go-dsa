package helpers_test

import (
	"go-dsa/templates/helpers"
	"testing"
)

func TestStripNonAlphabeticalRegexp(t *testing.T) {
	cases := []struct {
		in       string
		expected string
	}{
		{in: "A man, a plan, a canal: Panama", expected: "AmanaplanacanalPanama"},
		{in: "race a car", expected: "raceacar"},
		{in: " ", expected: ""},
		{in: "", expected: ""},
		{in: "testnochange", expected: "testnochange"},
		{in: "$E%ig`~h&t8*Z@@er-o0_{E}ig,./h!>?t#", expected: "EightZeroEight"},
	}

	for _, c := range cases {
		actual := helpers.StripNonAlphabeticalRegexp(c.in)
		if actual != c.expected {
			t.Errorf("StripNonAlphabeticalRegexp: Expected [%s], but got [%s]", c.expected, actual)
		}
		actual = helpers.StripNonAlphabeticalUnicode(c.in)
		if actual != c.expected {
			t.Errorf("StripNonAlphabeticalUnicode: Expected [%s], but got [%s]", c.expected, actual)
		}
	}
}

func BenchmarkStripNonAlphabeticalRegexp(b *testing.B) {
	input := "A man a plan, a canal: Panama"
	for n := 0; n < b.N; n++ {
		helpers.StripNonAlphabeticalRegexp(input)
	}
}

func BenchmarkStripNonAlphabeticalUnicode(b *testing.B) {
	input := "A man a plan, a canal: Panama"
	for n := 0; n < b.N; n++ {
		helpers.StripNonAlphabeticalUnicode(input)
	}
}
