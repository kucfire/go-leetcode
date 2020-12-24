package leetcode448

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	datas := []struct {
		nums   []int
		result []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
		// {[]int{5}, []int{1, 2, 3, 4}},
		{[]int{1, 1}, []int{2}},
	}

	for i, data := range datas {
		if actual := FindDisappearedNumbers(data.nums); !reflect.DeepEqual(data.result, actual) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestFindDisappearedNumbers2(t *testing.T) {
	datas := []struct {
		nums   []int
		result []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
		// {[]int{5}, []int{1, 2, 3, 4}},
		{[]int{1, 1}, []int{2}},
	}

	for i, data := range datas {
		if actual := FindDisappearedNumbers2(data.nums); !reflect.DeepEqual(data.result, actual) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
