package first

import (
	"strings"
	"testing"
)

func FindShortestWords(s []string) string {
	var res []string

	if len(s) == 0 {
		return ""
	}

	minLen := len(s[0])
	for _, v := range s {
		if len(v) < minLen {
			minLen = len(v)
		}
	}

	for _, v := range s {
		if len(v) == minLen {
			res = append(res, v)
		}
	}
	return strings.Join(res, " ")
}

func TestFindShortestWords(t *testing.T) {
	tests := []struct {
		words    []string
		expected string
	}{
		{[]string{"The", "quick", "brown", "fox", "jumps"}, "The fox"},
		{[]string{"Hello", "world"}, "Hello world"},
		{[]string{}, ""},
		{[]string{"a", "ab", "abc"}, "a"},
		{[]string{"test", "shortest", "words"}, "test"},
		{[]string{"word", "word", "word"}, "word word word"},
	}

	for _, test := range tests {
		result := FindShortestWords(test.words)
		if result != test.expected {
			t.Errorf("For words %v, expected \"%s\", but got \"%s\"", test.words, test.expected, result)
		}
	}
}
