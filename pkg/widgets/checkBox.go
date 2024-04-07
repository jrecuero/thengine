// checkBox.go module contains all attributes and methods required to implement
// a basic and generic checkbox widget.
package widgets

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// CheckBox
//
// -----------------------------------------------------------------------------

// Checkbox structure defines a baseline for any checkbox widget.
type CheckBox struct {
	*Widget
	selections     []string
	selected       []bool
	selectionIndex int
	scroller       *Scroller
}

// NewCheckBox function creates a new CheckBox instance.
func NewCheckBox(name string, position *api.Point, size *api.Size, style *tcell.Style, selections []string, selectionIndex int) *CheckBox {
	selectionsLength := len(selections)
	// Add padding for every menu item to fill the whole horizontal length.
	paddingSelections := make([]string, selectionsLength)
	for i, s := range selections {
		paddingSelections[i] = fmt.Sprintf("%-*s", size.W-4, s)
	}
	tools.Logger.WithField("module", "check-box").WithField("function", "NewCheckBox").Debugf("%s %+v", name, paddingSelections)
	checkBox := &CheckBox{
		Widget:         NewWidget(name, position, size, style),
		selections:     paddingSelections,
		selected:       make([]bool, selectionsLength),
		selectionIndex: selectionIndex,
	}
	checkBox.scroller = NewVerticalScroller(selectionsLength, size.H-2)
	checkBox.SetFocusType(engine.SingleFocus)
	checkBox.SetFocusEnable(true)
	checkBox.updateCanvas()
	return checkBox
}

// -----------------------------------------------------------------------------
// CheckBox private methods
// -----------------------------------------------------------------------------

func (c *CheckBox) execute(args ...any) {
	tools.Logger.WithField("module", "check-box").WithField("function", "execute").Debugf("%s %+v", c.GetName(), args)
	switch args[0].(string) {
	case "up":
		if c.selectionIndex > 0 {
			c.selectionIndex--
			c.updateCanvas()
		}
	case "down":
		if c.selectionIndex < (len(c.selections) - 1) {
			c.selectionIndex++
			c.updateCanvas()
		}
	case "run":
		c.selected[c.selectionIndex] = !c.selected[c.selectionIndex]
		c.updateCanvas()
	}
}

// updateCanvas method updates the check box canvas with proper selections to be
// displayed and the proper selected option.
func (c *CheckBox) updateCanvas() {
	// update the scroller with the selection index.
	c.scroller.Update(c.selectionIndex)
	canvas := c.GetCanvas()
	c.scroller.CreateIter()
	for x := 1; c.scroller.IterHasNext(); {
		index, y := c.scroller.IterGetNext()
		selection := c.selections[index]
		if c.selected[index] {
			selection = "x " + selection
		} else {
			selection = "- " + selection
		}
		if index == c.selectionIndex {
			canvas.WriteStringInCanvasAt(selection, c.GetStyle(), api.NewPoint(x, y))
		} else {
			reverseStyle := tools.ReverseStyle(c.GetStyle())
			canvas.WriteStringInCanvasAt(selection, reverseStyle, api.NewPoint(x, y))
		}
	}
}

// -----------------------------------------------------------------------------
// CheckBox public methods
// -----------------------------------------------------------------------------

// GetSelection method returns a list of strings with all options being
// selected.
func (c *CheckBox) GetSelection() []string {
	var result []string
	for index, selection := range c.selections {
		if c.selected[index] {
			result = append(result, selection)
		}
	}
	return result
}

// SetSelection method update the list of selected selections in the check box
// widget.
func (c *CheckBox) SetSelection(indexes ...int) {
	for _, index := range indexes {
		c.selected[index] = true
	}
	c.updateCanvas()
}

// Update method executes all check box functionality every tick time. Keyboard
// inut is scanned in order to move the selection index and proceed to select
// any option.
func (c *CheckBox) Update(event tcell.Event) {
	if !c.HasFocus() {
		return
	}
	actions := []*KeyboardAction{
		{
			Key:      tcell.KeyDown,
			Callback: c.execute,
			Args:     []any{"down"},
		},
		{
			Key:      tcell.KeyUp,
			Callback: c.execute,
			Args:     []any{"up"},
		},
		{
			Key:      tcell.KeyEnter,
			Callback: c.execute,
			Args:     []any{"run"},
		},
	}
	c.HandleKeyboardForActions(event, actions)
}

var _ engine.IEntity = (*CheckBox)(nil)
