package leetcode191

import (
	"strconv"
	"testing"
)

func TestHammingWeight2(t *testing.T) {
	example1, _ := strconv.ParseUint("00000000000000000000000000001011", 2, 32)
	example2, _ := strconv.ParseUint("00000000000000000000000010000000", 2, 32)
	example3, _ := strconv.ParseUint("11111111111111111111111111111101", 2, 32)
	datas := []struct {
		num    uint32
		result int
	}{
		{uint32(example1), 3},
		{uint32(example2), 1},
		{uint32(example3), 31},
	}

	for i, data := range datas {
		if actual := HammingWeight2(data.num); actual != data.result {
			t.Errorf("example %d actually got %d, but expect %d", i+1, actual, data.result)
		}
	}
}

func TestHammingWeight(t *testing.T) {
	example1, _ := strconv.ParseUint("00000000000000000000000000001011", 2, 32)
	example2, _ := strconv.ParseUint("00000000000000000000000010000000", 2, 32)
	example3, _ := strconv.ParseUint("11111111111111111111111111111101", 2, 32)
	datas := []struct {
		num    uint32
		result int
	}{
		{uint32(example1), 3},
		{uint32(example2), 1},
		{uint32(example3), 31},
	}

	for i, data := range datas {
		if actual := HammingWeight(data.num); actual != data.result {
			t.Errorf("example %d actually got %d, but expect %d", i+1, actual, data.result)
		}
	}
}
