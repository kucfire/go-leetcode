package leetcode026

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, data := range datas {
		if actual := RemoveDuplicates(data.nums); actual != data.result {
			t.Errorf("actually got %d; but expect %d", actual, data.result)
		}
	}
}
