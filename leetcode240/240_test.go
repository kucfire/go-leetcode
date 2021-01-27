package leetcode240

import "testing"

func TestSearchMatrix(t *testing.T) {
	datas := []struct {
		tasks  [][]int
		target int
		result bool
	}{
		{[][]int{{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}, 5, true},
		{[][]int{{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}, 20, false},
		{[][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		}, 19, true},
	}

	for i, data := range datas {
		if actual := SearchMatrix2(data.tasks, data.target); actual != data.result {
			t.Errorf("%d, got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestSearchMatrixReview(t *testing.T) {
	datas := []struct {
		tasks  [][]int
		target int
		result bool
	}{
		{[][]int{{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}, 5, true},
		{[][]int{{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}, 20, false},
		{[][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		}, 19, true},
	}

	for i, data := range datas {
		if actual := SearchMatrixReview(data.tasks, data.target); actual != data.result {
			t.Errorf("%d, got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestSearchMatrix3(t *testing.T) {
	datas := []struct {
		tasks  [][]int
		target int
		result bool
	}{
		{[][]int{{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}, 5, true},
		{[][]int{{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}, 20, false},
		{[][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		}, 19, true},
	}

	for i, data := range datas {
		if actual := SearchMatrix3(data.tasks, data.target); actual != data.result {
			t.Errorf("%d, got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
