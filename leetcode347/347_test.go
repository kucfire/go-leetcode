package leetcode347

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	datas := []struct {
		nums   []int
		k      int
		result []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
	}

	for i, data := range datas {
		if actual := TopKFrequent(data.nums, data.k); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("example %d actually got %v, but expect %v", i, actual, data.result)
		}
	}
}
