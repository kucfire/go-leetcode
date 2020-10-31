package leetcode387

import "testing"

func TestFirstUniqChar(t *testing.T) {
	datas := []struct {
		s      string
		result int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aadadaad", -1},
	}

	for _, data := range datas {
		if actual := FirstUniqChar(data.s); actual != data.result {
			t.Errorf("actually got %d, but expect %d", actual, data.result)
		}
	}
}
