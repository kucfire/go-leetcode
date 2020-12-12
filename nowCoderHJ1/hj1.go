/*
	牛客网 tag:hj1 title:字符串最后一个单词的长度
*/
package nowCoderHJ1

import "strings"

// LastWordLen : normal method , fail
func LastWordLen(s string) int {
	var result []string
	tmpWord := ""
	for i, str := range s {
		// 空字符串会在这里被过滤
		if string(str) == " " && len(tmpWord) > 0 {
			result = append(result, tmpWord)
			tmpWord = ""
			continue
		}

		tmpWord = tmpWord + string(str)

		// 边界判断
		if i == len(s)-1 && len(tmpWord) > 0 {
			result = append(result, tmpWord)
			tmpWord = ""
		}
	}
	return len(result[len(result)-1])
}

// LastWordLen2 : strings.Split , fail
func LastWordLen2(s string) int {
	if len(s) == 0 || s == " " {
		return 0
	}
	arrayS := strings.Split(s, " ")
	return len(arrayS[len(arrayS)-1])

}

// LastWordLen3 : 从后往前遍历
func LastWordLen3(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		}
		if count != 0 && s[i] == ' ' {
			break
		}
	}
	return count
}

// func main() {
// 	fmt.Println(lastWordLen("hello nowcoder"))
// }
