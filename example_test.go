package wordgame_test

import (
	"fmt"

	"github.com/aquilax/wordgame"
)

func ExampleWordList_Filter() {
	wl := wordgame.NewFromStrings([]string{
		"cow",
		"chicken",
		"horse",
		"brocolly",
	})
	result := wl.Filter(wordgame.GivenWithExtra("co", 0))
	fmt.Printf("%+v", result)
	// Output: [cow brocolly]
}

func ExampleWordList_FilterLen() {
	wl := wordgame.NewFromStrings([]string{
		"cow",
		"chicken",
		"horse",
		"brocolly",
		"coworker",
		"comb",
	})
	result := wl.Filter(wordgame.GivenWithExtra("co", 4))
	fmt.Printf("%+v", result)
	// Output: [comb]
}
