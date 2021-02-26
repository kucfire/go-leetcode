/*
	leetcode tag:179 title:最大数
*/

package leetcode179

import (
	"sort"
	"strconv"
	"strings"
)

func LargestNumber(nums []int) string {
	listStr := make([]string, 0)
	for _, num := range nums {
		listStr = append(listStr, strconv.Itoa(num))
	}

	sort.Slice(listStr, func(x, y int) bool {
		return listStr[x]+listStr[y] >= listStr[y]+listStr[x]
	})

	result := strings.Join(listStr, "")
	if result[0] == '0' {
		return "0"
	}

	return result
}
