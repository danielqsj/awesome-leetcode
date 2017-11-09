/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/
package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	tmp1, tmp2 := 0, 0
	if len1 == 0 && len2 == 0 {
		return -1
	} else if len1 == 0 {
		if len2%2 == 0 {
			return (float64(nums2[len2/2]) + float64(nums2[len2/2-1])) / 2
		} else {
			return float64(nums2[len2/2])
		}
	} else if len2 == 0 {
		if len1%2 == 0 {
			return (float64(nums1[len1/2]) + float64(nums1[len1/2-1])) / 2
		} else {
			return float64(nums1[len1/2])
		}
	}
	for i, j, k := 0, 0, 0; ; k++ {
		if k == (len1+len2)/2-1 {
			if i == len1 {
				tmp1 = nums2[j]
			} else if j == len2 {
				tmp1 = nums1[i]
			} else {
				tmp1 = min(nums1[i], nums2[j])
			}
		} else if k == (len1+len2)/2 {
			if i == len1 {
				tmp2 = nums2[j]
			} else if j == len2 {
				tmp2 = nums1[i]
			} else {
				tmp2 = min(nums1[i], nums2[j])
			}
			break
		}
		if i == len1 {
			j++
			continue
		}
		if j == len2 {
			i++
			continue
		}
		if nums1[i] >= nums2[j] {
			j++
		} else {
			i++
		}
	}
	if (len1+len2)%2 == 0 {
		return (float64(tmp1) + float64(tmp2)) / 2.0
	} else {
		return float64(tmp2)
	}
}

func min(num1, num2 int) int {
	if num1 >= num2 {
		return num2
	} else {
		return num1
	}
}

func main() {
	fmt.Printf("result: %v", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
