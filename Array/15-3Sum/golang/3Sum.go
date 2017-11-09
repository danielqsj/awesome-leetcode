/*
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.

For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	tmp := nums[0]
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == tmp {
			continue
		} else {
			tmp = nums[i]
		}
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {

			}
		}
	}
}

func main() {
	fmt.Printf("result: %v", threeSum([]int{1, 2}))
}
