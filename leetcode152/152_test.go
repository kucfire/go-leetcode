package leetcode152

import (
	"reflect"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
	}

	for i, data := range datas {
		if actual := MaxProduct(data.nums); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestMaxProductDP(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2}, -2},
	}

	for i, data := range datas {
		if actual := MaxProductDP(data.nums); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
