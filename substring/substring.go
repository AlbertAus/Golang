package main

import (
	"fmt"
	"strings"
)

func main() {
	textToSearch := "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
	subtext := "Peter"

	fmt.Print(SubString(textToSearch, subtext))
}

// SubString function receiving two strings as input. One called 'textToSerach' and another one called 'subtext',
// it will return all the positions of the beginning of each match for the subtext within the textToSearch. (See readme.md file)
func SubString(textToSearch string, subtext string) []int {
	var p []int

	// Change the input texts and subtext to lower cases.
	textToSearch = strings.ToLower(textToSearch)
	subtext = strings.ToLower(subtext)

	// Get the lengths of both textToSearch and subtext
	textToSearchLength := len(textToSearch)
	subtextLength := len(subtext)
	// tmpText :=

	for i := 0; i < (textToSearchLength - subtextLength); i++ {

		// Compare two strings and append the results to postion p []int array
		if textToSearch[i:i+subtextLength] == subtext {
			p = append(p, i+1)
		}
	}

	// Return the result of p
	if len(p) > 0 {
		return p
	} else {
		return nil
	}
}
