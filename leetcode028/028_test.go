package leetcode028

import "testing"

func TestStrStr(t *testing.T) {
	datas := []struct {
		haystack, needle string
		result           int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"a", "a", 0},
		{"aaa", "a", 0},
	}

	for i, data := range datas {
		if actual := StrStr(data.haystack, data.needle); actual != data.result {
			t.Errorf("example %d actuall got %d, but expect %d\n", i, actual, data.result)
		}
	}
}

func TestStrStr2(t *testing.T) {
	datas := []struct {
		haystack, needle string
		result           int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"a", "a", 0},
		{"aaa", "a", 0},
	}

	for i, data := range datas {
		if actual := StrStr2(data.haystack, data.needle); actual != data.result {
			t.Errorf("example %d actuall got %d, but expect %d\n", i, actual, data.result)
		}
	}
}
