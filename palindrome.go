//Determine if a given string is a palindrome

package main

import (
    "fmt"
    "strings"
)

func main() {
	var palindrome string = "tacocat"

	if isPalindrome(palindrome) {
		fmt.Printf("%v is a palindrome \n", palindrome)		
	}else{
		fmt.Printf("%v is NOT a palindrome \n", palindrome) 
	}		
}

func isPalindrome(palindrome string) bool {
	var isPalindrome bool = true
	
	pList := strings.Split(palindrome, "");
	
	for i := 0; i < len(pList)/2; i++ {
		if(pList[i] != pList[(len(pList) - i) - 1]){
			isPalindrome = false
		}
	}	

    return isPalindrome
}