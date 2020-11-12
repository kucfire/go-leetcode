/*
	leetcode tag:190 title:颠倒二进制位
*/

package leetcode190

// ReverseBits : method of bit operation
func ReverseBits(num uint32) uint32 {
	var res uint32 = 0
	point := 31
	for num != 0 {
		res += (num & 1) << point
		point--
		num >>= 1
	}
	return res
}
