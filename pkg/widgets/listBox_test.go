package widgets_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/widgets"
)

func TestListBox(t *testing.T) {
	var gotSelection string
	var exp string
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	got := widgets.NewListBox("test/1", api.NewPoint(0, 0), api.NewSize(20, 5), &styleOne, []string{"one", "two", "three"}, 0)
	if got == nil {
		t.Errorf("[1] NewListBox Error exp:*ListBox got:nil")
		return
	}
	exp = "one"
	gotSelection = got.GetSelection()
	if gotSelection != exp {
		t.Errorf("[1] GetSelection Error exp:%s got:%s", exp, gotSelection)
	}

	event := tcell.NewEventKey(tcell.KeyDown, 0, 0)
	got.Update(event)
	gotSelection = got.GetSelection()
	if gotSelection != exp {
		t.Errorf("[1] GetSelection Error exp:%s got:%s", exp, gotSelection)
	}

	got.AcquireFocus()
	got.Update(event)
	exp = "two"
	gotSelection = got.GetSelection()
	if gotSelection != exp {
		t.Errorf("[1] GetSelection Error exp:%s got:%s", exp, gotSelection)
	}

	got.Update(event)
	exp = "three"
	gotSelection = got.GetSelection()
	if gotSelection != exp {
		t.Errorf("[1] GetSelection Error exp:%s got:%s", exp, gotSelection)
	}

	got.Update(event)
	exp = "three"
	gotSelection = got.GetSelection()
	if gotSelection != exp {
		t.Errorf("[1] GetSelection Error exp:%s got:%s", exp, gotSelection)
	}

	event = tcell.NewEventKey(tcell.KeyUp, 0, 0)
	got.Update(event)
	exp = "two"
	gotSelection = got.GetSelection()
	if gotSelection != exp {
		t.Errorf("[1] GetSelection Error exp:%s got:%s", exp, gotSelection)
	}

	event = tcell.NewEventKey(tcell.KeyUp, 0, 0)
	got.Update(event)
	exp = "one"
	gotSelection = got.GetSelection()
	if gotSelection != exp {
		t.Errorf("[1] GetSelection Error exp:%s got:%s", exp, gotSelection)
	}

	event = tcell.NewEventKey(tcell.KeyUp, 0, 0)
	got.Update(event)
	exp = "one"
	gotSelection = got.GetSelection()
	if gotSelection != exp {
		t.Errorf("[1] GetSelection Error exp:%s got:%s", exp, gotSelection)
	}
}
