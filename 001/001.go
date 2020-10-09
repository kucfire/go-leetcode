/*
	leetcode tag:001 title:两数之和
*/

package main

import (
	"fmt"
)

//method of violence
func twoSum(nums []int, target int) []int {
	// var res []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				// res = append(res, i, j)
				return []int{i, j}
			}
		}
	}
	return nil
}

//method of double points : 2020-10-05 isn't work
func twoSum2(nums []int, target int) []int {
	//对nums进行排序
	// sort.Ints(nums)
	head, tail := 0, len(nums)-1
	for head < tail {
		if nums[head]+nums[tail] > target {
			tail--
		} else if nums[head]+nums[tail] < target {
			head++
		} else {
			return []int{head, tail}
		}
	}
	return nil
}

//method of hashMap : the best method
func twoSum3(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		if v, ok := hashMap[num]; ok {
			// nums = nums[:2]
			nums[0] = v
			nums[1] = i
			return nums[:2]

		}
		hashMap[target-num] = i
	}
	return nil
}

func main() {
	//example
	example1 := []int{2, 7, 11, 26}
	example1Target := 9

	//method
	fmt.Println(twoSum(example1, example1Target))
	fmt.Println(twoSum2(example1, example1Target))
	fmt.Println(twoSum3(example1, example1Target))
}
