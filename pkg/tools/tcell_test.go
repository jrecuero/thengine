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
