/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int)
	other := 0
	for i := 0; i < len(nums); i++ {
		other = target - nums[i]
		if val, ok := tmpMap[other]; ok {
			return []int{val, i}
		}
		tmpMap[nums[i]] = i
	}
	return nil
}

func main() {
	fmt.Printf("result: %v", twoSum([]int{2, 7, 11, 15}, 9))
}
