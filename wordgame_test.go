package wordgame

import (
	"reflect"
	"testing"
)

func TestSearchString(t *testing.T) {
	testCases := []struct {
		words    []string
		chars    string
		expected []string
	}{
		{[]string{"abc", "cdb", "rrr"}, "cb", []string{"abc", "cdb"}},
		{[]string{}, "", []string{}},
		{[]string{"abcdefgh"}, "r", []string{}},
		{[]string{"aabbcc", "acb"}, "aacb", []string{"aabbcc"}},
	}
	for _, tc := range testCases {
		wl := NewFromStrings(tc.words)
		result := wl.SearchString(tc.chars)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %+v, got %+v", tc.expected, result)
		}
	}
}

func BenchmarkSearchString(b *testing.B) {
	words := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	len := len(words)
	wl := NewFromStrings(words)
	for n := 0; n < b.N; n++ {
		wl.SearchString(words[n%len])
	}
}
