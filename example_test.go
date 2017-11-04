package wordgame_test

import (
	"fmt"

	"github.com/aquilax/wordgame"
)

func ExampleWordList_SearchString() {
	wl := wordgame.NewFromStrings([]string{
		"cow",
		"chicken",
		"horse",
		"brocolly",
	})
	result := wl.SearchString("co")
	fmt.Printf("%+v", result)
	// Output: [cow brocolly]
}
