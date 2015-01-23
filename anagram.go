/* 
- Checks is if list of words are anagrams for variable 'anagram'
- if words aren't equal in length, skip it homes.
- Check sum using ASCII values, if they're the same, you're in anagram town.
*/

package main

import (
    "fmt"
)

func main() {
	var anagram string = "alerting"
	var alist = []string{ "altering", "integral", "hamburger", "relating", "triangle", "pickles" }
	
	score := getSum(anagram)	

	for _, elem := range alist {

		if len(anagram) != len(elem){
			continue
		}
		
		if score == getSum(elem){
			fmt.Printf("%v is an anagram of %v \n", elem, anagram)
		}
	}	
}

func getSum(word string) int32{
	var sum int32	

	for _, elem := range word {
		sum += elem
	}

	return sum
}