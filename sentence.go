/* 
Takes a sentence and changes each letter to the next letter in the alphabet
*/

package main

import (
    "fmt"
)

func main() {
	var sentence string = "I love coding"
	var newSentence string

	for _, elem := range sentence {		

		if string(elem) == " "{
			newSentence = newSentence + string(elem)
			continue
		}

		newSentence = newSentence + string(elem + 1)
	}	

	fmt.Println(newSentence)
}
