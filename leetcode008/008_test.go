package leetcode008

import "testing"

func TestMyAtoi(t *testing.T) {
	datas := []struct {
		s      string
		result int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"+1", 1},
		{"9223372036854775808", 2147483647},
	}

	for _, data := range datas {
		if actual := MyAtoi(data.s); actual != data.result {
			t.Errorf("actually got %d, but expect %d\n", actual, data.result)
		}
	}
}

func TestMyAtoi2(t *testing.T) {
	datas := []struct {
		s      string
		result int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"+1", 1},
		{"9223372036854775808", 2147483647},
		{"-+12", 0},
		{"  -0012a42", -12},
	}

	for _, data := range datas {
		if actual := MyAtoi2(data.s); actual != data.result {
			t.Errorf("actually got %d, but expect %d\n", actual, data.result)
		}
	}
}
