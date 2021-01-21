package leetcode009

import "testing"

func TestIsPalindrome(t *testing.T) {
	datas := []struct {
		x      int
		result bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}

	for i, data := range datas {
		if actual := IsPalindrome(data.x); actual != data.result {
			t.Errorf("Example %d actually got %v,but expect %v", i+1, actual, data.result)
		}
	}
}
