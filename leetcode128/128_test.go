package leetcode128

import "testing"

func TestLongestConsecutive(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 0, -1}, 3},
		{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
	}

	for i, data := range datas {
		if actual := LongestConsecutive(data.nums); actual != data.result {
			t.Errorf("Example %v actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
