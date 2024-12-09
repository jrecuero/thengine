// choose.go module contains all logic required to implement a choose widget
// where only one option from all possible selection can be selected.
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
// Choose
//
// -----------------------------------------------------------------------------

// Choose structure defines a baseline for any choose widget
type Choose struct {
	*Widget
	selections     []string
	selected       int
	selectionIndex int
	scroller       *Scroller
}

// -----------------------------------------------------------------------------
// New Choose functions
// -----------------------------------------------------------------------------

// NewChoose function creates a new Choose instance.
func NewChoose(name string, position *api.Point, size *api.Size, style *tcell.Style,
	selections []string, selectionIndex int) *Choose {
	selectionsLength := len(selections)
	// Add padding for every menu item to fill the whole horizontal length.
	paddingSelections := make([]string, selectionsLength)
	for i, s := range selections {
		paddingSelections[i] = fmt.Sprintf("%-*s", size.W-4, s)
	}
	tools.Logger.WithField("module", "choose").
		WithField("function", "NewChoose").
		Debugf("%s %+v", name, paddingSelections)
	choose := &Choose{
		Widget:         NewWidget(name, position, size, style),
		selections:     paddingSelections,
		selected:       -1,
		selectionIndex: selectionIndex,
	}
	choose.scroller = NewVerticalScroller(selectionsLength, size.H-2)
	choose.SetFocusType(engine.SingleFocus)
	choose.SetFocusEnable(true)
	choose.updateCanvas()
	return choose
}

// -----------------------------------------------------------------------------
// Choose private methods
// -----------------------------------------------------------------------------

// execute method handles any keyboard input.
func (c *Choose) execute(args ...any) {
	tools.Logger.WithField("struct", "choose").
		WithField("method", "execute").
		Debugf("%s %+v", c.GetName(), args)
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
		c.selected = c.selectionIndex
		c.updateCanvas()
	}
}

// updateCanvas method updates the choose canvas with proper selections to be
// displayed and the proper selected option.
func (c *Choose) updateCanvas() {
	// update the scroller with the selection index.
	c.scroller.Update(c.selectionIndex)
	canvas := c.GetCanvas()
	canvas.WriteRectangleInCanvasAt(nil, nil, c.GetStyle(), engine.CanvasRectSingleLine)
	c.scroller.CreateIter()
	for x, y := 1, 1; c.scroller.IterHasNext(); y++ {
		index, _ := c.scroller.IterGetNext()
		selection := c.selections[index]
		if c.selected == index {
			selection = "x " + selection
		} else {
			selection = "- " + selection
		}
		if index == c.selectionIndex {
			reverseStyle := tools.ReverseStyle(c.GetStyle())
			canvas.WriteStringInCanvasAt(selection, reverseStyle, api.NewPoint(x, y))
		} else {
			canvas.WriteStringInCanvasAt(selection, c.GetStyle(), api.NewPoint(x, y))
		}
	}
}

// -----------------------------------------------------------------------------
// Choose public methods
// -----------------------------------------------------------------------------

// GetSelected method returns the selected option.
func (c *Choose) GetSelected() int {
	return c.selected
}

// SetSelection method sets the selected option.
func (c *Choose) SetSelected(index int) {
	c.selected = index
	c.updateCanvas()
}

// Update method executes all choose functionality every tick time. Keyboard
// inut is scanned in order to move the selection index and proceed to select
// any option.
func (c *Choose) Update(event tcell.Event, scene engine.IScene) {
	defer c.Entity.Update(event, scene)
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

var _ engine.IObject = (*Choose)(nil)
var _ engine.IFocus = (*Choose)(nil)
var _ engine.IEntity = (*Choose)(nil)
