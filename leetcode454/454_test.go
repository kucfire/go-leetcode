package leetcode454

import "testing"

func TestFourSumCount(t *testing.T) {
	datas := []struct {
		A      []int
		B      []int
		C      []int
		D      []int
		result int
	}{
		{
			[]int{1, 2},
			[]int{-2, -1},
			[]int{-1, 2},
			[]int{0, 2},
			2,
		},
	}

	for i, data := range datas {
		if actual := FourSumCount(data.A, data.B, data.C, data.D); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
