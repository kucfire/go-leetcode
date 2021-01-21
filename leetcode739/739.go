/*
	leetcode tag:739 title:每日温度
*/

package leetcode739

// DailyTemperatures : method of double points
func DailyTemperatures(T []int) []int {
	var result []int

loop2:
	for i, tem := range T {
		for j := i + 1; j < len(T); j++ {
			if T[j]-tem > 0 {
				result = append(result, j-i)
				continue loop2
			}
		}
		result = append(result, 0)
	}
	return result
}

// DailyTemperatures2 : monotonic stack (reverse)
func DailyTemperatures2(T []int) []int {
	// var result []int
	length := len(T)
	result := make([]int, length)
	stack := []int{}

	// 从数组末尾开始遍历，则单调栈需要以单调递增的方式去设计
	for i := length - 1; i >= 0; i-- {
		// 如果栈里存在元素，并且栈顶的元素大于新元素
		for len(stack) > 0 {
			if T[i] >= T[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				result[i] = stack[len(stack)-1] - i
				break
			}
		}
		stack = append(stack, i)
	}

	return result
}

// DailyTemperatures3 : monotonic stack
func DailyTemperatures3(T []int) []int {
	// var result []int
	length := len(T)
	result := make([]int, length)
	stack := []int{}

	// 从数组末尾开始遍历，则单调栈需要以单调递增的方式去设计
	for i := 0; i < length; i++ {
		// 如果栈里存在元素，并且栈顶的元素大于新元素
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}
