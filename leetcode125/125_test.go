package leetcode125

import "testing"

func TestIsPalindrome(t *testing.T) {
	datas := []struct {
		s      string
		result bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}

	for _, data := range datas {
		if actual := IsPalindrome(data.s); actual != data.result {
			t.Errorf("got %v; but expect %v", actual, data.result)
		}
	}
}
