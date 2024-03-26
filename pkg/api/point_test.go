package api_test

import (
	"math"
	"testing"

	"github.com/jrecuero/thengine/pkg/api"
)

func TestPointNew(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0},
			exp:   []int{0, 0},
		},
		{
			input: []int{1, 0},
			exp:   []int{1, 0},
		},
		{
			input: []int{0, 1},
			exp:   []int{0, 1},
		},
		{
			input: []int{1, 1},
			exp:   []int{1, 1},
		},
	}
	for i, c := range cases {
		got := api.NewPoint(c.input[0], c.input[1])
		if got == nil {
			t.Errorf("[%d] NewPoint Error exp:*Point got:nil", i)
			continue
		}
		if c.exp[0] != got.X {
			t.Errorf("[%d] NewPoint X Error exp:%d got:%d", i, c.exp[0], got.X)
		}
		if c.exp[1] != got.Y {
			t.Errorf("[%d] NewPoint Y Error exp:%d got %d", i, c.exp[1], got.Y)
		}
	}
}

func TestPointGet(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0},
			exp:   []int{0, 0},
		},
		{
			input: []int{1, 0},
			exp:   []int{1, 0},
		},
		{
			input: []int{0, 1},
			exp:   []int{0, 1},
		},
		{
			input: []int{1, 1},
			exp:   []int{1, 1},
		},
	}
	for i, c := range cases {
		p := api.NewPoint(c.input[0], c.input[1])
		gotX, gotY := p.Get()
		if c.exp[0] != gotX {
			t.Errorf("[%d] Get X Error exp:%d got:%d", i, c.exp[0], gotX)
		}
		if c.exp[1] != gotY {
			t.Errorf("[%d] Get Y Error exp:%d got %d", i, c.exp[1], gotY)
		}
	}
}

func TestPointSet(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0},
			exp:   []int{0, 0},
		},
		{
			input: []int{1, 0},
			exp:   []int{1, 0},
		},
		{
			input: []int{0, 1},
			exp:   []int{0, 1},
		},
		{
			input: []int{1, 1},
			exp:   []int{1, 1},
		},
	}
	for i, c := range cases {
		got := api.NewPoint(0, 0)
		got.Set(c.input[0], c.input[1])
		if c.exp[0] != got.X {
			t.Errorf("[%d] Set X Error exp:%d got:%d", i, c.exp[0], got.X)
		}
		if c.exp[1] != got.Y {
			t.Errorf("[%d] Set Y Error exp:%d got %d", i, c.exp[1], got.Y)
		}
	}
}

func TestPointClone(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0},
			exp:   []int{0, 0},
		},
		{
			input: []int{1, 0},
			exp:   []int{1, 0},
		},
		{
			input: []int{0, 1},
			exp:   []int{0, 1},
		},
		{
			input: []int{1, 1},
			exp:   []int{1, 1},
		},
	}
	for i, c := range cases {
		toClone := api.NewPoint(c.input[0], c.input[1])
		got := api.ClonePoint(toClone)
		if got == nil {
			t.Errorf("[%d] ClonePoint Error exp:*Point got:nil", i)
			continue
		}
		if c.exp[0] != got.X {
			t.Errorf("[%d] ClonePoint X Error exp:%d got:%d", i, c.exp[0], got.X)
		}
		if c.exp[1] != got.Y {
			t.Errorf("[%d] ClonePoint Y Error exp:%d got %d", i, c.exp[1], got.Y)
		}
	}
	for _, c := range cases {
		toClone := api.NewPoint(c.input[0], c.input[1])
		got := api.NewPoint(0, 0)
		got.Clone(toClone)
		if c.exp[0] != got.X {
			t.Errorf("Clone X Error exp:%d got:%d", c.exp[0], got.X)
		}
		if c.exp[1] != got.Y {
			t.Errorf("Clone Y Error exp:%d got %d", c.exp[1], got.Y)
		}
	}
}

func TestPointIsEqual(t *testing.T) {
	cases := []struct {
		input []int
		exp   bool
	}{
		{
			input: []int{0, 0, 0, 0},
			exp:   true,
		},
		{
			input: []int{1, 0, 1, 0},
			exp:   true,
		},
		{
			input: []int{0, 1, 0, 1},
			exp:   true,
		},
		{
			input: []int{7, 7, 7, 7},
			exp:   true,
		},
		{
			input: []int{0, 1, 0, 0},
			exp:   false,
		},
		{
			input: []int{1, 0, 0, 0},
			exp:   false,
		},
		{
			input: []int{0, 0, 0, 1},
			exp:   false,
		},
		{
			input: []int{7, 7, 5, 7},
			exp:   false,
		},
	}
	for i, c := range cases {
		p1 := api.NewPoint(c.input[0], c.input[1])
		p2 := api.NewPoint(c.input[2], c.input[3])
		got := p1.IsEqual(p2)
		if c.exp != got {
			t.Errorf("[%d] IsEqual Error exp:%t got:%t", i, c.exp, got)
		}
		got = p2.IsEqual(p1)
		if c.exp != got {
			t.Errorf("[%d] IsEqual (reverse) Error exp:%t got:%t", i, c.exp, got)
		}
	}
}

func TestPointAdd(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0, 0, 0},
			exp:   []int{0, 0},
		},
		{
			input: []int{1, 0, 1, 0},
			exp:   []int{2, 0},
		},
		{
			input: []int{0, 1, 0, 1},
			exp:   []int{0, 2},
		},
		{
			input: []int{1, 2, 3, 4},
			exp:   []int{4, 6},
		},
	}
	for i, c := range cases {
		got := api.NewPoint(c.input[0], c.input[1])
		p2 := api.NewPoint(c.input[2], c.input[3])
		got.Add(p2)
		if c.exp[0] != got.X {
			t.Errorf("[%d] Add-X Error exp:%d got:%d", i, c.exp[0], got.X)
		}
		if c.exp[1] != got.Y {
			t.Errorf("[%d] Add-Y Error exp:%d got:%d", i, c.exp[1], got.Y)
		}
	}
}

func TestPointDistance(t *testing.T) {
	cases := []struct {
		input []int
		exp   []any
	}{
		{
			input: []int{0, 0, 0, 0},
			exp:   []any{0.0, 0, 0},
		},
		{
			input: []int{0, 0, 2, 2},
			exp:   []any{math.Sqrt(8), 2, 2},
		},
		{
			input: []int{5, 4, 2, 1},
			exp:   []any{math.Sqrt(18), -3, -3},
		},
		{
			input: []int{5, 1, 2, 5},
			exp:   []any{5.0, -3, 4},
		},
	}
	for i, c := range cases {
		p1 := api.NewPoint(c.input[0], c.input[1])
		p2 := api.NewPoint(c.input[2], c.input[3])
		gotDistance, gotX, gotY := p1.Distance(p2)
		if c.exp[0] != gotDistance {
			t.Errorf("[%d] Distance Error exp:%f got:%f", i, c.exp[0], gotDistance)
		}
		if c.exp[1] != gotX {
			t.Errorf("[%d] Distance-X Error exp:%d got:%d", i, c.exp[1], gotX)
		}
		if c.exp[2] != gotY {
			t.Errorf("[%d] Distance-Y Error exp:%d got:%d", i, c.exp[2], gotY)
		}
	}
}

func TestPointToString(t *testing.T) {
	cases := []struct {
		input []int
		exp   string
	}{
		{
			input: []int{0, 0},
			exp:   "(0,0)",
		},
		{
			input: []int{1, 0},
			exp:   "(1,0)",
		},
		{
			input: []int{0, 1},
			exp:   "(0,1)",
		},
		{
			input: []int{1, 1},
			exp:   "(1,1)",
		},
	}
	for i, c := range cases {
		p := api.NewPoint(c.input[0], c.input[1])
		got := p.ToString()
		if c.exp != got {
			t.Errorf("[%d] ToString Error exp:%s got:%s", i, c.exp, got)
		}
	}
}

func TestPointSaveToDict(t *testing.T) {
	cases := []struct {
		input []int
		exp   map[string]any
	}{
		{
			input: []int{0, 0},
			exp:   map[string]any{"x": 0, "y": 0},
		},
		{
			input: []int{10, 10},
			exp:   map[string]any{"x": 10, "y": 10},
		},
	}
	for i, c := range cases {
		point := api.NewPoint(c.input[0], c.input[1])
		got := point.SaveToDict()
		if c.exp["x"] != got["x"] {
			t.Errorf("[%d] SaveToDict X exp:%d got:%d", i, c.exp["x"], got["x]"])
		}
		if c.exp["y"] != got["y"] {
			t.Errorf("[%d] SaveToDict Y exp:%d got:%d", i, c.exp["y"], got["y"])
		}
	}
}
