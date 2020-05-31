package shaker

import "testing"

func TestDice(t *testing.T) {
	min, max := 1, 6
	actual := dice()
	if actual < min || actual > max {
		t.Error("Expected an int between [1,6], Got", actual)
	}
}

func TestShake(t *testing.T) {
	type test struct {
		dice int
		min  int
		max  int
	}

	tests := []test{
		{2, 11, 66},
		{4, 1111, 6666},
		{6, 111111, 666666},
	}

	for _, v := range tests {
		actual := Shake(v.dice)
		if actual < v.min || actual > v.max {
			t.Error("Expected an int between [",
				v.min, ",", v.max, "], Got:", actual)
		}
	}
}
