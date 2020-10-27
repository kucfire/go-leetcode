package leetcode066

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	datas := []struct {
		digits []int
		result []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}

	for _, data := range datas {
		if actual := PlusOne(data.digits); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("actually got %v, but expect %v", actual, data.result)
		}
	}
}
