package leetcode334

import "testing"

func TestIncreasingTriplet(t *testing.T) {
	datas := []struct {
		nums   []int
		result bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{2, 1, 5, 0, 4, 6}, true},
		{[]int{2, 4, -2, -3}, false},
		{[]int{20, 100, 10, 12, 5, 13}, true},
		{[]int{1, 5, 0, 4, 1, 3}, true},
	}

	for i, data := range datas {
		if actual := IncreasingTriplet(data.nums); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
