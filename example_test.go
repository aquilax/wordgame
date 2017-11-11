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

func ExampleWordList_Filter_Len() {
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

func ExampleWordList_FilterConcurrent() {
	wl := wordgame.NewFromStrings([]string{
		"cow",
		"chicken",
		"horse",
		"brocolly",
	})
	result := wl.FilterConcurrent(wordgame.GivenWithExtra("co", 0), 2)
	fmt.Printf("%+v", result)
	// Output: [cow brocolly]
}

func ExampleWordList_FilterConcurrent_OnlyGiven() {
	wl := wordgame.NewFromStrings([]string{
		"cow",
		"chicken",
		"horse",
		"brocolly",
	})
	result := wl.FilterConcurrent(wordgame.OnlyGiven("cowz", 0), 2)
	fmt.Printf("%+v", result)
	// Output: [cow]
}
