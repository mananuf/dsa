package main

import (
	"fmt"
)

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single
// digit. Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the
// number 0 itself.
//
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
// Example 2:

// Input: l1 = [0], l2 = [0]
// Output: [0]
// Example 3:

// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
	}
	return dummy.Next
}

// func reverseSlice(s []int) []int {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// 	return s
// }
// func sliceToInt(s []int) (r int) {
// 	for _, digit := range s {
// 		r = r*10 + digit
// 	}
// 	return
// }

// func intToSlice(n int) (s []int) {
// 	if n == 0 {
// 		s = []int{0}
// 		return
// 	}

// 	for n > 0 {
// 		s = append([]int{n % 10}, s...)
// 		n /= 10
// 	}

// 	return s
// }

// func sliceToListNode(s []int) *ListNode {
// 	if len(s) == 0 {
// 		return nil
// 	}
	
// 	var emptyListNode ListNode
// 	current := &emptyListNode
// 	for _, val := range s {
// 		current.Next = &ListNode{Val: val}
// 		current = current.Next
// 	}

// 	return emptyListNode.Next
// }

func main() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	v := addTwoNumbers(l1, l2)
	for v != nil {
		fmt.Println(v.Val, v.Next)
		v = v.Next
	}
}
