package main

import (
	"fmt"
)

/*
*
* Given a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:

Input: s = "cbbd"
Output: "bb"
*/
func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)

		maxLen := len1
		if len2 > len1 {
			maxLen = len2
		}

		if maxLen > (end - start + 1) {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}

func main() {
	s := longestPalindrome("babad")

	fmt.Println(s)

	s = longestPalindrome("cbbd")

	fmt.Println(s)
}
