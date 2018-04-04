package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func checkPalindrome(s string) (result bool) {
	s1 := strings.ToLower(s)
	result = strings.Compare(s1, reverseString(s1)) == 0
	return
}

func main() {
	var inputVal string
	fmt.Println("Enter string:")
	fmt.Scanf("%s\n", &inputVal)
	if checkPalindrome(inputVal) {
		fmt.Println("Entered value is not palindrome")
	} else {
		fmt.Println("Its a palindrome")
	}
}
