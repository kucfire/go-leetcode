package leetcode283

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	datas := []struct {
		nums   []int
		result []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}

	for _, data := range datas {
		if MoveZeroes(data.nums); !reflect.DeepEqual(data.nums, data.result) {
			t.Errorf("actually got %v, but exoect %v\n", data.nums, data.result)
		}
	}
}

func TestMoveZeroesReview(t *testing.T) {
	datas := []struct {
		nums   []int
		result []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}

	for _, data := range datas {
		if MoveZeroesReview(data.nums); !reflect.DeepEqual(data.nums, data.result) {
			t.Errorf("actually got %v, but exoect %v\n", data.nums, data.result)
		}
	}
}
