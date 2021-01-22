package leetcode027

import "testing"

func TestRemoveElement(t *testing.T) {
	datas := []struct {
		nums   []int
		val    int
		result int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for i, data := range datas {
		if actual := RemoveElement(data.nums, data.val); actual != data.result {
			t.Errorf("Example %d actually got %d, but expect %d", i+1, actual, data.result)
		}
	}
}
