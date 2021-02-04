package leetcode169

import "testing"

func TestMajorityElementCount(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{2, 2}, 2},
	}

	for i, data := range datas {
		if actual := MajorityElementCount(data.nums); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestMajorityElementHashmap(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{2, 2}, 2},
	}

	for i, data := range datas {
		if actual := MajorityElementHashmap(data.nums); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestMajorityElementSort(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{2, 2}, 2},
	}

	for i, data := range datas {
		if actual := MajorityElementSort(data.nums); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestMajorityElementBit(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{2, 2}, 2},
		{[]int{-2147483648}, -2147483648},
	}

	for i, data := range datas {
		if actual := MajorityElementBit(data.nums); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestMajorityElementReview(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{2, 2}, 2},
		{[]int{-2147483648}, -2147483648},
	}

	for i, data := range datas {
		if actual := MajorityElementReview(data.nums); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestMajorityElementReviewStack(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
		{[]int{2, 2}, 2},
		{[]int{-2147483648}, -2147483648},
	}

	for i, data := range datas {
		if actual := MajorityElementReviewStack(data.nums); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
