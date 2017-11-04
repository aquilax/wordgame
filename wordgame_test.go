package wordgame

import (
	"reflect"
	"testing"
)

func TestSearchString(t *testing.T) {
	testCases := []struct {
		words    []string
		chars    string
		expected Matches
	}{
		{[]string{"abc", "cdb", "rrr"}, "cb", Matches{"abc", "cdb"}},
		{[]string{}, "", Matches{}},
		{[]string{"abcdefgh"}, "r", Matches{}},
		{[]string{"aabbcc", "acb"}, "aacb", Matches{"aabbcc"}},
	}
	for _, tc := range testCases {
		wl := NewFromStrings(tc.words)
		result := wl.SearchString(tc.chars)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %+v, got %+v", tc.expected, result)
		}
	}
}
