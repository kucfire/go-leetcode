package leetcode055

import "testing"

func TestCanJump(t *testing.T) {
	datas := []struct {
		nums   []int
		result bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{2, 0}, true},
		{[]int{2, 0, 0}, true},
	}

	for i, data := range datas {
		if actual := CanJump(data.nums); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
