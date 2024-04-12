package api_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/api"
)

func TestSizeNew(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0},
			exp:   []int{0, 0},
		},
		{
			input: []int{10, 0},
			exp:   []int{10, 0},
		},
		{
			input: []int{0, 10},
			exp:   []int{0, 10},
		},
		{
			input: []int{10, 10},
			exp:   []int{10, 10},
		},
	}
	for i, c := range cases {
		got := api.NewSize(c.input[0], c.input[1])
		if got == nil {
			t.Errorf("[%d] NewSize Error exp:*Size got:nil", i)
			continue
		}
		if c.exp[0] != got.W {
			t.Errorf("[%d] NewSize X Error exp:%d got:%d", i, c.exp[0], got.W)
		}
		if c.exp[1] != got.H {
			t.Errorf("[%d] NewSize Y Error exp:%d got %d", i, c.exp[1], got.H)
		}
	}
}

func TestSizeGet(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0},
			exp:   []int{0, 0},
		},
		{
			input: []int{10, 0},
			exp:   []int{10, 0},
		},
		{
			input: []int{0, 10},
			exp:   []int{0, 10},
		},
		{
			input: []int{10, 10},
			exp:   []int{10, 10},
		},
	}
	for i, c := range cases {
		p := api.NewSize(c.input[0], c.input[1])
		gotX, gotY := p.Get()
		if c.exp[0] != gotX {
			t.Errorf("[%d] Get X Error exp:%d got:%d", i, c.exp[0], gotX)
		}
		if c.exp[1] != gotY {
			t.Errorf("[%d] Get Y Error exp:%d got %d", i, c.exp[1], gotY)
		}
	}
}

func TestSizeSet(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0},
			exp:   []int{0, 0},
		},
		{
			input: []int{10, 0},
			exp:   []int{10, 0},
		},
		{
			input: []int{0, 10},
			exp:   []int{0, 10},
		},
		{
			input: []int{10, 10},
			exp:   []int{10, 10},
		},
	}
	for i, c := range cases {
		got := api.NewSize(0, 0)
		got.Set(c.input[0], c.input[1])
		if c.exp[0] != got.W {
			t.Errorf("[%d] Set X Error exp:%d got:%d", i, c.exp[0], got.W)
		}
		if c.exp[1] != got.H {
			t.Errorf("[%d] Set Y Error exp:%d got %d", i, c.exp[1], got.H)
		}
	}
}

func TestSizeClone(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0},
			exp:   []int{0, 0},
		},
		{
			input: []int{10, 0},
			exp:   []int{10, 0},
		},
		{
			input: []int{0, 10},
			exp:   []int{0, 10},
		},
		{
			input: []int{10, 10},
			exp:   []int{10, 10},
		},
	}
	for i, c := range cases {
		toClone := api.NewSize(c.input[0], c.input[1])
		got := api.CloneSize(toClone)
		if got == nil {
			t.Errorf("[%d] CloneSize Error exp:*Size got:nil", i)
			continue
		}
		if c.exp[0] != got.W {
			t.Errorf("[%d] CloneSize X Error exp:%d got:%d", i, c.exp[0], got.W)
		}
		if c.exp[1] != got.H {
			t.Errorf("[%d] CloneSize Y Error exp:%d got %d", i, c.exp[1], got.H)
		}
	}
	for i, c := range cases {
		toClone := api.NewSize(c.input[0], c.input[1])
		got := api.NewSize(0, 0)
		got.Clone(toClone)
		if c.exp[0] != got.W {
			t.Errorf("[%d] Clone X Error exp:%d got:%d", i, c.exp[0], got.W)
		}
		if c.exp[1] != got.H {
			t.Errorf("[%d] Clone Y Error exp:%d got %d", i, c.exp[1], got.H)
		}
	}
}

func TestSizeIsEqual(t *testing.T) {
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
		p1 := api.NewSize(c.input[0], c.input[1])
		p2 := api.NewSize(c.input[2], c.input[3])
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

func TestSizeIsZeroSize(t *testing.T) {
	cases := []struct {
		input []int
		exp   bool
	}{
		{
			input: []int{0, 0},
			exp:   true,
		},
		{
			input: []int{1, 0},
			exp:   false,
		},
		{
			input: []int{0, 1},
			exp:   false,
		},
		{
			input: []int{7, 7},
			exp:   false,
		},
	}
	for i, c := range cases {
		point := api.NewSize(c.input[0], c.input[1])
		got := point.IsZeroSize()
		if c.exp != got {
			t.Errorf("[%d] IsZeroSize Error exp:%t got:%t", i, c.exp, got)
		}
	}
}

func TestSizeToString(t *testing.T) {
	cases := []struct {
		input []int
		exp   string
	}{
		{
			input: []int{0, 0},
			exp:   "(0-0)",
		},
		{
			input: []int{1, 0},
			exp:   "(1-0)",
		},
		{
			input: []int{0, 1},
			exp:   "(0-1)",
		},
		{
			input: []int{1, 1},
			exp:   "(1-1)",
		},
	}
	for i, c := range cases {
		p := api.NewSize(c.input[0], c.input[1])
		got := p.ToString()
		if c.exp != got {
			t.Errorf("[%d] ToString Error exp:%s got:%s", i, c.exp, got)
		}
	}
}

func TestSizeSaveToDict(t *testing.T) {
	cases := []struct {
		input []int
		exp   map[string]any
	}{
		{
			input: []int{0, 0},
			exp:   map[string]any{"w": 0, "h": 0},
		},
		{
			input: []int{10, 10},
			exp:   map[string]any{"w": 10, "h": 10},
		},
	}
	for i, c := range cases {
		size := api.NewSize(c.input[0], c.input[1])
		got := size.SaveToDict()
		if c.exp["w"] != got["w"] {
			t.Errorf("[%d] SaveToDict W exp:%d got:%d", i, c.exp["w"], got["w"])
		}
		if c.exp["h"] != got["h"] {
			t.Errorf("[%d] SaveToDict H exp:%d got:%d", i, c.exp["h"], got["h"])
		}
	}
}
