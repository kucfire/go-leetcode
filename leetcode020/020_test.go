package leetcode020

import "testing"

func TestIsValid(t *testing.T) {
	datas := []struct {
		s      string
		result bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for i, data := range datas {
		if actual := IsValid(data.s); actual != data.result {
			t.Errorf("example %d actual got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
