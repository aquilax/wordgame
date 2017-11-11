# wordgame [![Build Status](https://travis-ci.org/aquilax/wordgame.svg?branch=master)](https://travis-ci.org/aquilax/wordgame) [![GoDoc](https://godoc.org/github.com/aquilax/wordgame?status.svg)](https://godoc.org/github.com/aquilax/wordgame)

Package wordgame provides a dictionary search for word games.

Given a dictionary and list of required characteds, the search returns
list of matching words which include all the required characters.

## Usage

```go
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
```
