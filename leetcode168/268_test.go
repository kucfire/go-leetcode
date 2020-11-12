package leetcode268

import "testing"

func TestMissingNumber(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{0}, 1},
		{[]int{1}, 0},
	}

	for i, data := range datas {
		if actual := MissingNumber(data.nums); actual != data.result {
			t.Errorf("example %d actually got %d, but expect %d", i+1, actual, data.result)
		}
	}
}
