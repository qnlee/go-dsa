package string_manip_test

import (
	"go-dsa/template/string_manip"
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
		actual := string_manip.StripNonAlphabeticalRegexp(c.in)
		if actual != c.expected {
			t.Errorf("StripNonAlphabeticalRegexp(%s): Expected [%s], but got [%s]", c.in, c.expected, actual)
		}
		actual = string_manip.StripNonAlphabeticalUnicode(c.in)
		if actual != c.expected {
			t.Errorf("StripNonAlphabeticalUnicode(%s): Expected [%s], but got [%s]", c.in, c.expected, actual)
		}
	}
}

func BenchmarkStripNonAlphabeticalRegexp(b *testing.B) {
	input := "A man a plan, a canal: Panama"
	for n := 0; n < b.N; n++ {
		string_manip.StripNonAlphabeticalRegexp(input)
	}
}

func BenchmarkStripNonAlphabeticalUnicode(b *testing.B) {
	input := "A man a plan, a canal: Panama"
	for n := 0; n < b.N; n++ {
		string_manip.StripNonAlphabeticalUnicode(input)
	}
}
