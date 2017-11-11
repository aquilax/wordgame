// Package wordgame provides a dictionary search for word games.
// Given a dictionary and list of required characteds, the search returns
// list of matching words which include all the required characters.
package wordgame

import "sync"

// FilterFunc filtering function
type FilterFunc func([]rune) bool

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

func (sm searchMap) copy() searchMap {
	newMap := make(searchMap)
	for k, v := range sm {
		newMap[k] = v
	}
	return newMap
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

// Filter searches for matches by applying the filter function
func (wl WordList) Filter(ff FilterFunc) []string {
	result := make([]string, 0)
	for _, w := range wl {
		if ff(w) {
			result = append(result, string(w))
		}
	}
	return result
}

// FilterConcurrent searches for matches by applying the filter function using
// pool of workers
func (wl WordList) FilterConcurrent(ff FilterFunc, numWorkers int) []string {
	result := make([]string, 0)
	in := make(chan []rune)
	out := make(chan []rune)

	var w sync.WaitGroup
	var w2 sync.WaitGroup

	for n := 0; n < numWorkers; n++ {
		w.Add(1)
		go func() {
			for wrd := range in {
				if ff(wrd) {
					out <- wrd
				}
			}
			w.Done()
		}()
	}

	for _, wrd := range wl {
		in <- wrd
	}

	w2.Add(1)
	go func() {
		for wrd := range out {
			result = append(result, string(wrd))
		}
		w2.Done()
	}()

	close(in)
	w.Wait()
	close(out)
	w2.Wait()
	return result
}

// GivenWithExtra returns function matching all given characters plus any
// other character
func GivenWithExtra(letters string, l int) FilterFunc {
	sm := newSearchMap([]rune(letters))
	return func(w []rune) bool {
		if l > 0 && l != len(w) {
			return false
		}
		for sr, c := range sm {
			found := 0
			for _, r := range w {
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
}

// OnlyGiven returns function matching only the given characters
func OnlyGiven(letters string, l int) FilterFunc {
	sm := newSearchMap([]rune(letters))
	return func(w []rune) bool {
		if l > 0 && l != len(w) {
			return false
		}
		smap := sm.copy()
		for _, r := range w {
			if _, ok := smap[r]; !ok {
				return false
			}
			smap[r] = smap[r] - 1
			if smap[r] < 0 {
				return false
			}
		}
		return true
	}
}
