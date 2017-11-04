// Package wordgame provides a dictionary search for word games.
// Given a dictionary and list of required characteds, the search returns
// list of matching words which include all the required characters.
package wordgame

type searchMap map[rune]int

func newSearchMap(ar []rune) searchMap {
	sm := make(searchMap)
	for _, r := range ar {
		if _, ok := sm[r]; !ok {
			sm[r] = 0
		}
		sm[r] = sm[r] + 1
	}
	return sm
}

// WordList holds the dictionary
type WordList [][]rune

// New creates new empty dictionary
func New() *WordList {
	return &WordList{}
}

// NewFromStrings creates new dictionary and populates it with words
func NewFromStrings(sl []string) *WordList {
	var wl WordList
	for _, w := range sl {
		wl = append(wl, []rune(w))
	}
	return &wl
}

// SearchString searches for matches by stirng of characters
func (wl WordList) SearchString(s string) []string {
	return wl.Search([]rune(s))
}

// Search searches for matches by array of runes
func (wl WordList) Search(ar []rune) []string {
	result := make([]string, 0)
	sm := newSearchMap(ar)
	for _, w := range wl {
		if IsValid(sm, w) {
			result = append(result, string(w))
		}
	}
	return result
}

// IsValid returns true if the characters are contained in the word
func IsValid(sm searchMap, word []rune) bool {
	for sr, c := range sm {
		found := 0
		for _, r := range word {
			if r == sr {
				found++
			}
		}
		if found < c {
			return false
		}
	}
	return true
}
