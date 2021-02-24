/*
	leetcode tag:454 title:四数相加II
*/

package leetcode454

func FourSumCount(A []int, B []int, C []int, D []int) int {
	count := 0

	Lists := [][]int{A, B, C, D}
	// Target := 0

	var dp func(lists [][]int, target int)
	dp = func(lists [][]int, target int) {
		// 基准条件
		if len(lists) == 0 {
			if target == 0 {
				count++
			}
			return
		}

		// 循环条件
		list := lists[0]
		for _, num := range list {
			tmp := target
			tmp += num
			dp(lists[1:], tmp)
		}
	}

	dp(Lists, 0)

	return count
}

func FourSumCount2(A []int, B []int, C []int, D []int) int {
	count := 0
	Map1 := map[int]int{}
	for _, v := range A {
		for _, w := range B {
			Map1[v+w]++
		}
	}

	for _, v := range C {
		for _, w := range D {
			count += Map1[-(v + w)]
		}
	}

	return count
}
