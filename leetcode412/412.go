/*
	leetcode tag:412 title:Fizz Buzz
*/

package leetcode412

import "strconv"

// normal method
func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		tmp := ""
		if i%3 == 0 {
			tmp += "Fizz"
		}
		if i%5 == 0 {
			tmp += "Buzz"
		}
		if len(tmp) == 0 {
			tmp = strconv.Itoa(i)
		}
		result = append(result, tmp)
	}
	return result
}

// 内存消耗更小
func fizzBuzz2(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		// tmp := ""
		if i%3 == 0 && i%5 == 0 {
			result[i-1] = "FizzBuzz"
		} else if i%5 == 0 {
			result[i-1] = "Buzz"
		} else if i%3 == 0 {
			result[i-1] = "Fizz"
		} else {
			result[i-1] = strconv.Itoa(i)
		}
		// result = append(result, tmp)
	}
	return result
}
