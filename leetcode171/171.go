/*
	leetcode tag:171 title:Excel表列序号
*/

package leetcode171

func TitleToNumber(s string) int {
	// List := map[byte]int{
	// 	'A': 1,
	// 	'B': 2,
	// 	'C': 3,
	// 	'D': 4,
	// 	'E': 5,
	// 	'F': 6,
	// 	'G': 7,
	// 	'H': 8,
	// 	'I': 9,
	// 	'J': 10,
	// 	'K': 11,
	// 	'L': 12,
	// 	'M': 13,
	// 	'N': 14,
	// 	'O': 15,
	// 	'P': 16,
	// 	'Q': 17,
	// 	'R': 18,
	// 	'S': 19,
	// 	'T': 20,
	// 	'U': 21,
	// 	'V': 22,
	// 	'W': 23,
	// 	'X': 24,
	// 	'Y': 25,
	// 	'Z': 26,
	// }

	list := []rune(s)

	result := 0

	// for i := 0; i < len(s); i++ {
	// 	result = result*26 + List[s[i]]
	// }

	for _, r := range list {
		result = result*26 + (int(r-'A') + 1)
	}

	return result
}
