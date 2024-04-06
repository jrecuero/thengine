// listBox.go module contains all attributes and methods required to implement
// a basic and generic list box.
package widgets

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// ListBox
//
// -----------------------------------------------------------------------------

// ListBox structure defines a baseline for any listbox widget.
type ListBox struct {
	*Widget
	selections     []string
	selectionIndex int
	scroller       *Scroller
}

// NewListBox function creates a new ListBox instance.
func NewListBox(name string, position *api.Point, size *api.Size, style *tcell.Style, selections []string, selectionIndex int) *ListBox {
	selectionsLength := len(selections)
	// Add padding for every menu item to fill the whole horizontal length.
	paddingSelections := make([]string, selectionsLength)
	for i, s := range selections {
		paddingSelections[i] = fmt.Sprintf("%-*s", size.W-2, s)
	}
	tools.Logger.WithField("module", "list-box").WithField("function", "NewListBox").Debugf("%s %+v", name, paddingSelections)
	listBox := &ListBox{
		Widget:         NewWidget(name, position, size, style),
		selections:     paddingSelections,
		selectionIndex: selectionIndex,
	}
	listBox.scroller = NewVerticalScroller(selectionsLength, size.H-2)
	listBox.SetFocusType(engine.SingleFocus)
	listBox.SetFocusEnable(true)
	listBox.updateCanvas()
	return listBox
}

// -----------------------------------------------------------------------------
// ListBox private methods
// -----------------------------------------------------------------------------

func (l *ListBox) execute(args ...any) {
	tools.Logger.WithField("module", "list-box").WithField("function", "execute").Debugf("%s %+v", l.GetName(), args)
	switch args[0].(string) {
	case "up":
		if l.selectionIndex > 0 {
			l.selectionIndex--
			l.updateCanvas()
		}
	case "down":
		if l.selectionIndex < (len(l.selections) - 1) {
			l.selectionIndex++
			l.updateCanvas()
		}
	case "run":
	}
}

// updateCanvas method updates the list box canvas with proper selections to be
// displayed and the proper selected option.
func (l *ListBox) updateCanvas() {
	// update the scroller with the selection index.
	l.scroller.Update(l.selectionIndex)
	tools.Logger.WithField("module", "list-box").WithField("function", "updateCanvas").Debugf("[%d, %d]", l.scroller.StartSelection, l.scroller.EndSelection)
	canvas := l.GetCanvas()
	l.scroller.CreateIter()
	for x := 1; l.scroller.IterHasNext(); {
		index, y := l.scroller.IterGetNext()
		selection := l.selections[index]
		tools.Logger.WithField("module", "list-box").WithField("function", "updateCanvas").Debugf("selection %s", selection)
		if index == l.selectionIndex {
			canvas.WriteStringInCanvasAt(selection, l.GetStyle(), api.NewPoint(x, y))
		} else {
			reverseStyle := tools.ReverseStyle(l.GetStyle())
			canvas.WriteStringInCanvasAt(selection, reverseStyle, api.NewPoint(x, y))
		}
	}
	//for index, x, y := l.scroller.StartSelection, 1, 1; index <= l.scroller.EndSelection; index, y = index+1, y+1 {
	//    selection := l.selections[index]
	//    tools.Logger.WithField("module", "list-box").WithField("function", "updateCanvas").Debugf("selection %s", selection)
	//    if index == l.selectionIndex {
	//        canvas.WriteStringInCanvasAt(selection, l.GetStyle(), api.NewPoint(x, y))
	//    } else {
	//        reverseStyle := tools.ReverseStyle(l.GetStyle())
	//        canvas.WriteStringInCanvasAt(selection, reverseStyle, api.NewPoint(x, y))
	//    }
	//}
}

// -----------------------------------------------------------------------------
// ListBox public methods
// -----------------------------------------------------------------------------

// GetSelection method returns the option for the selected index.
func (l *ListBox) GetSelection() string {
	return strings.TrimSpace(l.selections[l.selectionIndex])
}

// Update method executes all listbox functionality every tick time. Keyboard
// inut is scanned in order to move the selection index and proceed to select
// any option.
func (l *ListBox) Update(event tcell.Event) {
	if !l.HasFocus() {
		return
	}
	actions := []*KeyboardAction{
		{
			Key:      tcell.KeyDown,
			Callback: l.execute,
			Args:     []any{"down"},
		},
		{
			Key:      tcell.KeyUp,
			Callback: l.execute,
			Args:     []any{"up"},
		},
		{
			Key:      tcell.KeyEnter,
			Callback: l.execute,
			Args:     []any{"run"},
		},
	}
	l.HandleKeyboardForActions(event, actions)
}

var _ engine.IEntity = (*ListBox)(nil)
