package leetcode198

import "testing"

func TestRob(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1, 1, 2}, 4},
	}

	for i, data := range datas {
		if actual := Rob(data.nums); actual != data.result {
			t.Errorf("example %d actually got %d, but expect %d", i, actual, data.result)
		}
	}
}
