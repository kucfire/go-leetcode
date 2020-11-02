/*
	leetcode tag:038 title:外观数列
*/

package leetcode038

import (
	"strconv"
	"strings"
)

// CountAndSay : normal method
func CountAndSay(n int) string {
	result := "1"
	// count := 1
	var tmp strings.Builder
	for i := 2; i <= n; i++ {
		count := 1
		preByte := result[0]
		for j := 1; j < len(result); j++ { // 边界处理，count为1所以从result的第二位开始
			if result[j] == preByte {
				count++
			} else {
				tmp.WriteString(strconv.Itoa(count))
				tmp.WriteByte(preByte)
				preByte = result[j]
				count = 1 // 因为当前在j位上，而进行操作的是j-1上的，所以当前j会被跳过，count应该从1开始
			}
		}
		// result尾部元素处理
		tmp.WriteString(strconv.Itoa(count))
		tmp.WriteByte(preByte)
		result = tmp.String()
		tmp.Reset()
	}
	return result
}
