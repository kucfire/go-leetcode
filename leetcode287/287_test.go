package leetcode287

import "testing"

func TestFindDuplicate(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{
			[]int{1, 3, 4, 2, 2},
			2,
		},
		{
			[]int{3, 1, 3, 4, 2},
			3,
		},
		{
			[]int{1, 1},
			1,
		},
		{
			[]int{1, 1, 2},
			1,
		},
	}

	for i, data := range datas {
		if actual := FindDuplicate(data.nums); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
