# wordgame

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

```
