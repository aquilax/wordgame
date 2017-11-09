package wordgame

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	testCases := []struct {
		words    []string
		l        int
		chars    string
		expected []string
	}{
		{[]string{"abc", "cdb", "rrr"}, 0, "cb", []string{"abc", "cdb"}},
		{[]string{}, 0, "", []string{}},
		{[]string{"abcdefgh"}, 0, "r", []string{}},
		{[]string{"aabbcc", "acb"}, 0, "aacb", []string{"aabbcc"}},

		{[]string{"abc", "cdb", "rrr"}, 3, "cb", []string{"abc", "cdb"}},
		{[]string{}, 2, "", []string{}},
		{[]string{"abcdefgh"}, 2, "r", []string{}},
		{[]string{"aabbcc", "acb"}, 6, "aacb", []string{"aabbcc"}},
	}
	for _, tc := range testCases {
		wl := NewFromStrings(tc.words)
		result := wl.Filter(GivenWithExtra(tc.chars, tc.l))
		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %+v, got %+v", tc.expected, result)
		}
	}
}

func BenchmarkFilter(b *testing.B) {
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
		wl.Filter(GivenWithExtra(words[n%len], 0))
	}
}
