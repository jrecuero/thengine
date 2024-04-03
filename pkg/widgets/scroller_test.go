package widgets_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/widgets"
)

func TestScrollerVertical(t *testing.T) {
	cases := []struct {
		setup []int
		input []int
		exp   [][]int
	}{
		{
			setup: []int{4, 4},
			input: []int{0, 1, 2, 3, 2, 1, 0},
			exp:   [][]int{{0, 3}, {0, 3}, {0, 3}, {0, 3}, {0, 3}, {0, 3}, {0, 3}},
		},
		{
			setup: []int{5, 3},
			input: []int{0, 1, 2, 3, 4, 3, 2, 1, 0},
			exp:   [][]int{{0, 2}, {0, 2}, {0, 2}, {1, 3}, {2, 4}, {2, 4}, {2, 4}, {1, 3}, {0, 2}},
		},
		{
			setup: []int{2, 3},
			input: []int{0, 1, 1, 0},
			exp:   [][]int{{0, 1}, {0, 1}, {0, 1}, {0, 1}},
		},
	}
	for i, c := range cases {
		got := widgets.NewVerticalScroller(c.setup[0], c.setup[1])
		if got == nil {
			t.Errorf("[%d] NewVerticalScroller Error exp:*Scroller got:nil", i)
			return
		}
		for j, selection := range c.input {
			got.Update(selection)
			if c.exp[j][0] != got.StartSelection {
				t.Errorf("[%d] StartSelection [%d]Error exp:%d got:%d", i, j, c.exp[j][0], got.StartSelection)
			}
			if c.exp[j][1] != got.EndSelection {
				t.Errorf("[%d] EndSelection [%d]Error exp:%d got:%d", i, j, c.exp[j][1], got.EndSelection)
			}
		}
	}
}

func TestScrollerHorizontal(t *testing.T) {
	cases := []struct {
		setup []int
		input []int
		exp   [][]int
	}{
		{
			setup: []int{16, 16, 4},
			input: []int{0, 1, 2, 3, 2, 1, 0},
			exp:   [][]int{{0, 3}, {0, 3}, {0, 3}, {0, 3}, {0, 3}, {0, 3}, {0, 3}},
		},
		{
			setup: []int{20, 12, 4},
			input: []int{0, 1, 2, 3, 4, 3, 2, 1, 0},
			exp:   [][]int{{0, 2}, {0, 2}, {0, 2}, {1, 3}, {2, 4}, {2, 4}, {2, 4}, {1, 3}, {0, 2}},
		},
		{
			setup: []int{8, 12, 4},
			input: []int{0, 1, 1, 0},
			exp:   [][]int{{0, 1}, {0, 1}, {0, 1}, {0, 1}},
		},
	}
	for i, c := range cases {
		got := widgets.NewScroller(c.setup[0], c.setup[1], c.setup[2])
		if got == nil {
			t.Errorf("[%d] NewVerticalScroller Error exp:*Scroller got:nil", i)
			return
		}
		for j, selection := range c.input {
			got.Update(selection)
			if c.exp[j][0] != got.StartSelection {
				t.Errorf("[%d] StartSelection [%d]Error exp:%d got:%d", i, j, c.exp[j][0], got.StartSelection)
			}
			if c.exp[j][1] != got.EndSelection {
				t.Errorf("[%d] EndSelection [%d]Error exp:%d got:%d", i, j, c.exp[j][1], got.EndSelection)
			}
		}
	}
}
