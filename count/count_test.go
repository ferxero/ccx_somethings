package count

import (
	"testing"
)

func TestMax(t *testing.T) {
	cases := []struct {
		inx, iny, want int
	}{
		{1, 2, 2},
		{3, 2, 3},
		{0, 0, 0},
	}

	for _, c := range cases {
		got := Max(c.inx, c.iny)
		if got != c.want {
			t.Errorf("Max(%v, %v) == %v want %v", c.inx, c.iny, got, c.want)
		}
	}
}

func TestMaxOfN(t *testing.T) {
	cases := []struct {
		list []int
		want int
	}{
		{[]int{1, 2, 2, 4}, 4},
		{[]int{1, 2}, 2},
		{[]int{2, 2, 2}, 2},
	}
	for _, c := range cases {
		got := MaxOfN(c.list)
		if got != c.want {
			t.Errorf("MaxOfN(%v) == %v want %v", c.list, got, c.want)
		}
	}

}
