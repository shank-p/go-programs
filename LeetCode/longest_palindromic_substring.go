/*
	5. Longest Palindromic Substring
	- https://leetcode.com/problems/longest-palindromic-substring/description/
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(start int, end int, subStr *string) bool {
	for i:=0; i<(len((*subStr)[start:end])/2) ; i++ {
		if (*subStr)[start+i]!=(*subStr)[end-1-i] {
			return false
		}
	}
	return true
}


func longestPalindrome(s string) (subStr string) {
	for windowLen:=len(s); windowLen>1; windowLen-- {
		start := 0
		end := start + windowLen
		for ;end<=len(s); end++ {
			if isPalindrome(start, end, &s) {
				return s[start:end] 
			}
			start += 1
		} 
	}
		return string(s[0])
}


func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Print("Enter input string `s` : ")
	s, _ = reader.ReadString('\n')
	s = strings.TrimSpace(s)
	fmt.Println("Result : ", longestPalindrome(s))
}
