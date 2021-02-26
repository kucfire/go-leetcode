package leetcode162

import "testing"

func TestFindPeakElement(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 1},
	}

	for i, data := range datas {
		if actual := FindPeakElement(data.nums); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
