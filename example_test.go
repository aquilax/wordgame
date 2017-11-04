package wordgame_test

import (
	"fmt"

	"github.com/aquilax/wordgame"
)

func ExampleSearchString() {
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
