package tools_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

func TestIsSameStyle(t *testing.T) {
	cases := []struct {
		input1 *tcell.Style
		input2 *tcell.Style
		exp    bool
	}{
		{
			input1: engine.NewStyle(tcell.ColorBlack, tcell.ColorWhite, 0),
			input2: engine.NewStyle(tcell.ColorBlack, tcell.ColorWhite, 0),
			exp:    true,
		},
		{
			input1: engine.NewStyle(tcell.ColorBlack, tcell.ColorWhite, 0),
			input2: engine.NewStyle(tcell.ColorBlack, tcell.ColorRed, 0),
			exp:    false,
		},
		{
			input1: engine.NewStyle(tcell.ColorBlack, tcell.ColorWhite, 0),
			input2: engine.NewStyle(tcell.ColorBlue, tcell.ColorWhite, 0),
			exp:    false,
		},
		{
			input1: engine.NewStyle(tcell.ColorBlack, tcell.ColorWhite, 0),
			input2: engine.NewStyle(tcell.ColorYellow, tcell.ColorMaroon, 0),
			exp:    false,
		},
	}
	for i, c := range cases {
		got := tools.IsEqualStyle(c.input1, c.input2)
		if c.exp != got {
			t.Errorf("[%d] IsEqualStyle Error exp:%t got:%t", i, c.exp, got)
		}
	}
}

func TestReverseStyle(t *testing.T) {
	cases := []struct {
		input *tcell.Style
		exp   *tcell.Style
	}{
		{
			input: engine.NewStyle(tcell.ColorBlack, tcell.ColorWhite, 0),
			exp:   engine.NewStyle(tcell.ColorWhite, tcell.ColorBlack, 0),
		},
		{
			input: engine.NewStyle(tcell.ColorRed, tcell.ColorDefault, 0),
			exp:   engine.NewStyle(tcell.ColorDefault, tcell.ColorRed, 0),
		},
		{
			input: nil,
			exp:   nil,
		},
	}
	for i, c := range cases {
		got := tools.ReverseStyle(c.input)
		if (c.exp == nil) && (c.exp != got) {
			t.Errorf("[%d] ReverseStyle Error exp:nil got:%+v", i, got)
		}
		if (c.exp != nil) && !tools.IsEqualStyle(c.exp, got) {
			t.Errorf("[%d] ReverseStyle Error exp:%+v got:%+v", i, c.exp, got)
		}
	}
}
