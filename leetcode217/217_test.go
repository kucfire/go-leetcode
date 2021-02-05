package leetcode217

import "testing"

func TestContainsDuplicate(t *testing.T) {
	datas := []struct {
		nums   []int
		result bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, data := range datas {
		if actual := ContainsDuplicate(data.nums); actual != data.result {
			t.Errorf("actually got %v; but expect %v", actual, data.result)
		}
	}
}

func TestContainsDuplicate2(t *testing.T) {
	datas := []struct {
		nums   []int
		result bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, data := range datas {
		if actual := ContainsDuplicate2(data.nums); actual != data.result {
			t.Errorf("actually got %v; but expect %v", actual, data.result)
		}
	}
}

func TestContainsDuplicateReview(t *testing.T) {
	datas := []struct {
		nums   []int
		result bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{1}, false},
	}

	for _, data := range datas {
		if actual := ContainsDuplicateReview(data.nums); actual != data.result {
			t.Errorf("actually got %v; but expect %v", actual, data.result)
		}
	}
}
