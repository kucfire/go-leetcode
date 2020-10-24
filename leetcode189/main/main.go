package main

import (
	"fmt"
	"go-leetcode/leetcode189"
)

func main() {
	example1_nums := []int{1, 2, 3, 4, 5, 6, 7}
	example1_k := 3
	example2_nums := []int{-1, -100, 3, 99}
	example2_k := 2

	// leetcode189.Rotate(example1_nums, example1_k)
	// fmt.Println(example1_nums)

	// leetcode189.Rotate(example2_nums, example2_k)
	// fmt.Println(example2_nums)

	leetcode189.Rotate2(example1_nums, example1_k)
	fmt.Println(example1_nums)

	leetcode189.Rotate2(example2_nums, example2_k)
	fmt.Println(example2_nums)
}
