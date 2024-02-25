package engine_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

type cell struct {
	color *api.Color
	ch    rune
}

func TestCell(t *testing.T) {
	cases := []struct {
		input *cell
		exp   *cell
	}{
		{
			input: &cell{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
			exp:   &cell{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
		},
		{
			input: &cell{color: api.NewColor(api.ColorRed, api.ColorDefault), ch: 0},
			exp:   &cell{color: api.NewColor(api.ColorRed, api.ColorDefault), ch: 0},
		},
	}
	for i, c := range cases {
		got := engine.NewCell(c.input.color, c.input.ch)
		if !c.exp.color.IsEqual(got.Color) {
			t.Errorf("[%d] NewCell Color Error exp:%+v got:%+v", i, c.exp.color, got.Color)
		}
		if c.exp.ch != got.Rune {
			t.Errorf("[%d] NewCell Rune Error exp:%c got:%c", i, c.exp.ch, got.Rune)
		}
	}
}

func TestCellEmptyCell(t *testing.T) {
	cases := []struct {
		input *cell
		exp   *cell
	}{
		{
			input: nil,
			exp:   &cell{color: nil, ch: 0},
		},
	}
	for i, c := range cases {
		got := engine.NewEmptyCell()
		if c.exp.color != got.Color {
			t.Errorf("[%d] NewEmptyCell Color Error exp:%+v got:%+v", i, c.exp.color, got.Color)
		}
		if c.exp.ch != got.Rune {
			t.Errorf("[%d] NewEmptyCell Rune Error exp:%c got:%c", i, c.exp.ch, got.Rune)
		}
	}
}

func TestCellCloneCell(t *testing.T) {
	cases := []struct {
		input *cell
		exp   *cell
	}{
		{
			input: &cell{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
			exp:   &cell{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
		},
		{
			input: &cell{color: api.NewColor(api.ColorRed, api.ColorDefault), ch: 0},
			exp:   &cell{color: api.NewColor(api.ColorRed, api.ColorDefault), ch: 0},
		},
	}
	for i, c := range cases {
		toClone := engine.NewCell(c.input.color, c.input.ch)
		got := engine.CloneCell(toClone)
		if !c.exp.color.IsEqual(got.Color) {
			t.Errorf("[%d] CloneCell Color Error exp:%+v got:%+v", i, c.exp.color, got.Color)
		}
		if c.exp.ch != got.Rune {
			t.Errorf("[%d] CloneCell Rune Error exp:%c got:%c", i, c.exp.ch, got.Rune)
		}
	}
}

func TestCellClone(t *testing.T) {
	cases := []struct {
		input *cell
		exp   *cell
	}{
		{
			input: &cell{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
			exp:   &cell{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
		},
		{
			input: &cell{color: api.NewColor(api.ColorRed, api.ColorDefault), ch: 0},
			exp:   &cell{color: api.NewColor(api.ColorRed, api.ColorDefault), ch: 0},
		},
	}
	for i, c := range cases {
		toClone := engine.NewCell(c.input.color, c.input.ch)
		got := engine.NewEmptyCell()
		got.Clone(toClone)
		if !c.exp.color.IsEqual(got.Color) {
			t.Errorf("[%d] CloneCell Color Error exp:%+v got:%+v", i, c.exp.color, got.Color)
		}
		if c.exp.ch != got.Rune {
			t.Errorf("[%d] CloneCell Rune Error exp:%c got:%c", i, c.exp.ch, got.Rune)
		}
	}
}

func TestCellIsEqual(t *testing.T) {
	cases := []struct {
		input []*cell
		exp   bool
	}{
		{
			input: []*cell{
				{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
				{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
			},
			exp: true,
		},
		{
			input: []*cell{
				{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
				{color: api.NewColor(api.ColorRed, api.ColorWhite), ch: 'x'},
			},
			exp: false,
		},
		{
			input: []*cell{
				{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
				{color: api.NewColor(api.ColorBlue, api.ColorDefault), ch: 'x'},
			},
			exp: false,
		},
		{
			input: []*cell{
				{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
				{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'y'},
			},
			exp: false,
		},
	}
	for i, c := range cases {
		cell1 := engine.NewCell(c.input[0].color, c.input[0].ch)
		cell2 := engine.NewCell(c.input[1].color, c.input[1].ch)
		got := cell1.IsEqual(cell2)
		if c.exp != got {
			t.Errorf("[%d] IsEqual Error exp:%t got:%t", i, c.exp, got)
		}
		got = cell2.IsEqual(cell1)
		if c.exp != got {
			t.Errorf("[%d] IsEqual (reverse) Error exp:%t got:%t", i, c.exp, got)
		}
	}
}

func TestCellToString(t *testing.T) {
	cases := []struct {
		input *cell
		exp   string
	}{
		{
			input: &cell{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
			exp:   "[x][blue:white]",
		},
	}
	for i, c := range cases {
		newCell := engine.NewCell(c.input.color, c.input.ch)
		got := newCell.ToString()
		if c.exp != got {
			t.Errorf("[%d] ToString Error exp:%s got:%s", i, c.exp, got)
		}
	}
}

func TestCellSaveToDict(t *testing.T) {
	cases := []struct {
		input *cell
		exp   map[string]any
	}{
		{
			input: &cell{color: api.NewColor(api.ColorBlue, api.ColorWhite), ch: 'x'},
			exp:   map[string]any{"rune": 'x', "color": map[string]any{"fg": api.ColorBlue, "bg": api.ColorWhite}},
		},
	}
	for i, c := range cases {
		newCell := engine.NewCell(c.input.color, c.input.ch)
		got := newCell.SaveToDict()
		if c.exp["rune"] != got["rune"] {
			t.Errorf("[%d] SaveToDict Rune exp:%c got:%c", i, c.exp["rune"], got["rune"])
		}
		if c.exp["color"].(map[string]any)["fg"] != got["color"].(map[string]any)["fg"] {
			t.Errorf("[%d] SaveToDict Fg exp:%d got:%d",
				i, c.exp["color"].(map[string]any)["fg"], got["color"].(map[string]any)["fg"])
		}
		if c.exp["color"].(map[string]any)["bg"] != got["color"].(map[string]any)["bg"] {
			t.Errorf("[%d] SaveToDict Bg exp:%d got:%d",
				i, c.exp["color"].(map[string]any)["bg"], got["color"].(map[string]any)["bg"])
		}
	}
}
