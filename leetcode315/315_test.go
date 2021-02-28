package leetcode315

import (
	"reflect"
	"testing"
)

func TestCountSmaller(t *testing.T) {
	datas := []struct {
		nums   []int
		result []int
	}{
		{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
	}

	for i, data := range datas {
		if actual := CountSmaller(data.nums); !reflect.DeepEqual(data.result, actual) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
