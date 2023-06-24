package lc290_word_pattern

import "testing"

func Test_WordPattern(t *testing.T) {
	tests := map[string]struct {
		pattern  string
		s        string
		expected bool
	}{
		"Matching Mapping": {
			pattern:  "abba",
			s:        "dog cat cat dog",
			expected: true,
		},
		"Non-Matching Mapping": {
			pattern:  "abba",
			s:        "dog cat cat fish",
			expected: false,
		},
		"Repeated Words, Different Characters": {
			pattern:  "abba",
			s:        "dog cat dog cat",
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := wordPattern(test.pattern, test.s)
			if result != test.expected {
				t.Errorf("Expected %v for pattern: %s and string: %s, but got %v",
					test.expected, test.pattern, test.s, result)
			}
		})
	}
}
