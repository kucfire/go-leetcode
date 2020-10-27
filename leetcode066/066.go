package leetcode066

func PlusOne(digits []int) []int {
	tmp := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if tmp != 0 {
			digits[i] += tmp
			tmp = 0
		}
		if digits[i] >= 10 {
			tmp = digits[i] / 10
			digits[i] = digits[i] % 10
		} else {
			break
		}
		if i == 0 && tmp != 0 {
			digits = append([]int{tmp}, digits...)
		}
	}
	return digits
}
