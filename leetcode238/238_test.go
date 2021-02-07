package leetcode238

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	datas := []struct {
		nums   []int
		result []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{0, 0}, []int{0, 0}},
	}

	for i, data := range datas {
		if actual := ProductExceptSelf(data.nums); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
