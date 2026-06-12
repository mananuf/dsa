package main

import (
	"fmt"
	"math"
)

/*
 * 
 * DIFFICULTY: HARD
 * Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

 The overall run time complexity should be O(log (m+n)).

 Example 1:

 Input: nums1 = [1,3], nums2 = [2]
 Output: 2.00000
 Explanation: merged array = [1,2,3] and median is 2.
 Example 2:

 Input: nums1 = [1,2], nums2 = [3,4]
 Output: 2.50000
 Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := ((x + y + 1) / 2) - partitionX

		maxLeftX := math.MinInt64
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}

		minRightX := math.MaxInt64
		if partitionX < x {
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt64
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}

		minRightY := math.MaxInt64
		if partitionY < y {
			minRightY = nums2[partitionY]
		}

		// Check if we found the correct partition
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			// If the total number of elements is odd
			if (x+y)%2 == 1 {
				return float64(max(maxLeftX, maxLeftY))
			}
			// If the total number of elements is even
			return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
		} else if maxLeftX > minRightY {
			// We are too far to the right in nums1, move left
			high = partitionX - 1
		} else {
			// We are too far to the left in nums1, move right
			low = partitionX + 1
		}
	}

	return 0
}

func main() {
	nums1 := []int{1, 3, 4}
	nums2 := []int{2, 5}
	val := findMedianSortedArrays(nums1, nums2)
	fmt.Println(val)

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	val = findMedianSortedArrays(nums1, nums2)
	fmt.Println(val)

	nums1 = []int{1, 2, 7, 8}
	nums2 = []int{3, 4, 9}
	val = findMedianSortedArrays(nums1, nums2)
	fmt.Println(val)
}
