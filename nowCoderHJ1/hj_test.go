package nowCoderHJ1

import "testing"

func TestLastWordLen3(t *testing.T) {
	datas := []struct {
		s      string
		result int
	}{
		{"hello nowcoder", 8},
		{"XSUWHQ", 6},
		{"ABSIB T", 1},
		{"T ", 1},
		{" ", 0},
		{"     ", 0},
	}

	for i, data := range datas {
		if actual := LastWordLen3(data.s); data.result != actual {
			t.Errorf("Example %d actually got %d, but expect %d", i+1, actual, data.result)
		}
	}
}
