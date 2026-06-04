package main

import "fmt"

/**
DIFFICULTY: medium
Given a string s, find the length of the longest substring without duplicate characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	filterString := make(map[rune]int)
	best, start := 0, 0
	for i, char := range s {
		if prev, ok := filterString[char]; ok && prev > start {
			start = prev
		}

		if length := i - start + 1; length > best {
			best = length
		}

		filterString[char] = i + 1
	}

	return best
}

func main() {
	s := "abcabcbb"
	l := lengthOfLongestSubstring(s)

	fmt.Println(l)

	s = "bbbbb"
	l = lengthOfLongestSubstring(s)

	fmt.Println(l)

	s = "pwwkew"
	l = lengthOfLongestSubstring(s)

	fmt.Println(l)
}
