package engine_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/engine"
)

type cell struct {
	style *tcell.Style
	ch    rune
}

func TestCell(t *testing.T) {
	cases := []struct {
		input *cell
		exp   *cell
	}{
		{
			input: &cell{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
			exp:   &cell{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
		},
		{
			input: &cell{style: engine.NewStyle(tcell.ColorRed, tcell.ColorDefault, 0), ch: 0},
			exp:   &cell{style: engine.NewStyle(tcell.ColorRed, tcell.ColorDefault, 0), ch: 0},
		},
	}
	for i, c := range cases {
		got := engine.NewCell(c.input.style, c.input.ch)
		if !engine.CompareStyleColor(c.exp.style, got.Style) {
			t.Errorf("[%d] NewCell Color Error exp:%+v got:%+v", i, c.exp.style, got.Style)
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
			exp:   &cell{style: nil, ch: 0},
		},
	}
	for i, c := range cases {
		got := engine.NewEmptyCell()
		if c.exp.style != got.Style {
			t.Errorf("[%d] NewEmptyCell Color Error exp:%+v got:%+v", i, c.exp.style, got.Style)
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
			input: &cell{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
			exp:   &cell{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
		},
		{
			input: &cell{style: engine.NewStyle(tcell.ColorRed, tcell.ColorDefault, 0), ch: 0},
			exp:   &cell{style: engine.NewStyle(tcell.ColorRed, tcell.ColorDefault, 0), ch: 0},
		},
	}
	for i, c := range cases {
		toClone := engine.NewCell(c.input.style, c.input.ch)
		got := engine.CloneCell(toClone)
		if !engine.CompareStyleColor(c.exp.style, got.Style) {
			t.Errorf("[%d] CloneCell Color Error exp:%+v got:%+v", i, c.exp.style, got.Style)
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
			input: &cell{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
			exp:   &cell{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
		},
		{
			input: &cell{style: engine.NewStyle(tcell.ColorRed, tcell.ColorDefault, 0), ch: 0},
			exp:   &cell{style: engine.NewStyle(tcell.ColorRed, tcell.ColorDefault, 0), ch: 0},
		},
	}
	for i, c := range cases {
		toClone := engine.NewCell(c.input.style, c.input.ch)
		got := engine.NewEmptyCell()
		got.Clone(toClone)
		if !engine.CompareStyleColor(c.exp.style, got.Style) {
			t.Errorf("[%d] CloneCell Color Error exp:%+v got:%+v", i, c.exp.style, got.Style)
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
				{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
				{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
			},
			exp: true,
		},
		{
			input: []*cell{
				{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
				{style: engine.NewStyle(tcell.ColorRed, tcell.ColorWhite, 0), ch: 'x'},
			},
			exp: false,
		},
		{
			input: []*cell{
				{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
				{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorDefault, 0), ch: 'x'},
			},
			exp: false,
		},
		{
			input: []*cell{
				{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
				{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'y'},
			},
			exp: false,
		},
	}
	for i, c := range cases {
		cell1 := engine.NewCell(c.input[0].style, c.input[0].ch)
		cell2 := engine.NewCell(c.input[1].style, c.input[1].ch)
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
			input: &cell{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
			exp:   "[x][blue:white:0]",
		},
	}
	for i, c := range cases {
		newCell := engine.NewCell(c.input.style, c.input.ch)
		got := newCell.ToString()
		if c.exp != got {
			t.Errorf("[%d] ToString Error exp:%s got:%s", i, c.exp, got)
		}
	}
}

// func TestCellSaveToDict(t *testing.T) {
// 	cases := []struct {
// 		input *cell
// 		exp   map[string]any
// 	}{
// 		{
// 			input: &cell{style: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0), ch: 'x'},
// 			exp:   map[string]any{"rune": 'x', "color": map[string]any{"fg": tcell.ColorBlue, "bg": tcell.ColorWhite}},
// 		},
// 	}
// 	for i, c := range cases {
// 		newCell := engine.NewCell(c.input.style, c.input.ch)
// 		got := newCell.SaveToDict()
// 		if c.exp["rune"] != got["rune"] {
// 			t.Errorf("[%d] SaveToDict Rune exp:%c got:%c", i, c.exp["rune"], got["rune"])
// 		}
// 		if c.exp["color"].(map[string]any)["fg"] != got["color"].(map[string]any)["fg"] {
// 			t.Errorf("[%d] SaveToDict Fg exp:%d got:%d",
// 				i, c.exp["color"].(map[string]any)["fg"], got["color"].(map[string]any)["fg"])
// 		}
// 		if c.exp["color"].(map[string]any)["bg"] != got["color"].(map[string]any)["bg"] {
// 			t.Errorf("[%d] SaveToDict Bg exp:%d got:%d",
// 				i, c.exp["color"].(map[string]any)["bg"], got["color"].(map[string]any)["bg"])
// 		}
// 	}
// }
