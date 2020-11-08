/*
	leetcode tag:070 title:爬楼梯
*/

package leetcode070

// ClimbStairs : method of fiboracci(leetcode--time out)
func ClimbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return ClimbStairs(n-1) + ClimbStairs(n-2)
}

// ClimbStairs : method of array
func ClimbStairs2(n int) int {
	res := []int{1, 1}
	for i := 2; i <= n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res[n]
}
