// listBox.go module contains all attributes and methods required to implement
// a basic and generic list box widget.
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
	tools.Logger.WithField("module", "listbox").
		WithField("function", "NewListBox").
		Infof("%s %s %s %+v", name, position.ToString(), size.ToString(), selections)
	selectionsLength := len(selections)
	// Add padding for every menu item to fill the whole horizontal length.
	paddingSelections := make([]string, selectionsLength)
	for i, s := range selections {
		paddingSelections[i] = fmt.Sprintf("%-*s", size.W-2, s)
	}
	//tools.Logger.WithField("module", "listbox").
	//    WithField("function", "NewListBox").
	//    Debugf("%s %+v", name, paddingSelections)
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
	tools.Logger.WithField("module", "listbox").
		WithField("method", "execute").
		Debugf("%s %+v", l.GetName(), args)
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
		if l.callback != nil {
			l.RunCallback(l)
			//if l.callbackArgs != nil {
			//    l.callback(l, l.callbackArgs...)
			//} else {
			//    l.callback(l)
			//}
		}
	}
}

// updateCanvas method updates the list box canvas with proper selections to be
// displayed and the proper selected option.
func (l *ListBox) updateCanvas() {
	// update the scroller with the selection index.
	l.scroller.Update(l.selectionIndex)
	canvas := l.GetCanvas()
	canvas.WriteRectangleInCanvasAt(nil, nil, l.GetStyle(), engine.CanvasRectSingleLine)
	l.scroller.CreateIter()
	//tools.Logger.WithField("module", "listbox").
	//    WithField("method", "updateCanvas").
	//    Debugf("iter %s", iter.ToString())
	for x, y := 1, 1; l.scroller.IterHasNext(); y++ {
		index, _ := l.scroller.IterGetNext()
		selection := l.selections[index]
		if index == l.selectionIndex {
			reverseStyle := tools.ReverseStyle(l.GetStyle())
			canvas.WriteStringInCanvasAt(selection, reverseStyle, api.NewPoint(x, y))
		} else {
			canvas.WriteStringInCanvasAt(selection, l.GetStyle(), api.NewPoint(x, y))
		}
	}
}

// -----------------------------------------------------------------------------
// ListBox public methods
// -----------------------------------------------------------------------------

// GetSelection method returns the option for the selected index.
func (l *ListBox) GetSelection() string {
	return strings.TrimSpace(l.selections[l.selectionIndex])
}

// GetSelectionIndex method returns the selected index.
func (l *ListBox) GetSelectionIndex() int {
	return l.selectionIndex
}

// Update method executes all listbox functionality every tick time. Keyboard
// inut is scanned in order to move the selection index and proceed to select
// any option.
func (l *ListBox) Update(event tcell.Event, scene engine.IScene) {
	defer l.Entity.Update(event, scene)
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

var _ engine.IObject = (*ListBox)(nil)
var _ engine.IFocus = (*ListBox)(nil)
var _ engine.IEntity = (*ListBox)(nil)
