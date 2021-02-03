package leetcode344

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	datas := []struct {
		s      []byte
		result []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}

	for _, data := range datas {
		if ReverseString(data.s); !reflect.DeepEqual(data.s, data.result) {
			t.Errorf("actually got %v, but expect %v\n", string(data.s), string(data.result))
		}
	}
}

func TestReverseStringReview(t *testing.T) {
	datas := []struct {
		s      []byte
		result []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}

	for _, data := range datas {
		if ReverseStringReview(data.s); !reflect.DeepEqual(data.s, data.result) {
			t.Errorf("actually got %v, but expect %v\n", string(data.s), string(data.result))
		}
	}
}
