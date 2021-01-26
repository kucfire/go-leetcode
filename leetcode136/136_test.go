package leetcode136

import "testing"

func TestSingleNumber(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, data := range datas {
		if actual := SingleNumber(data.nums); actual != data.result {
			t.Errorf("actually got %d, but expect %d", actual, data.result)
		}
	}
}

func TestSingleNumber2(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, data := range datas {
		if actual := SingleNumber2(data.nums); actual != data.result {
			t.Errorf("actually got %d, but expect %d", actual, data.result)
		}
	}
}

func TestSingleNumberBit(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, data := range datas {
		if actual := SingleNumberBit(data.nums); actual != data.result {
			t.Errorf("actually got %d, but expect %d", actual, data.result)
		}
	}
}

func TestSingleNumberHashMap(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, data := range datas {
		if actual := SingleNumberHashMap(data.nums); actual != data.result {
			t.Errorf("actually got %d, but expect %d", actual, data.result)
		}
	}
}
