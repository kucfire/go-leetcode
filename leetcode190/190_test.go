package leetcode190

import "testing"

func TestReverseBits(t *testing.T) {
	datas := []struct {
		num    uint32
		result uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
	}

	for i, data := range datas {
		if actual := ReverseBits(data.num); actual != data.result {
			t.Errorf("example %v actual got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
