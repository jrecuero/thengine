// combobox.go module contains all attributes and methods required to implament
// a basic combobox widget.
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
// ComboBox
//
// -----------------------------------------------------------------------------

// ComboBox structure defines a baseline for any combobox widget.
type ComboBox struct {
	*Widget
	selections     []string
	filtered       []string
	selectionIndex int
	scroller       *Scroller
	inputStr       string
}

// NewComboBox function creates a new ComboBox instance.
func NewComboBox(name string, position *api.Point, size *api.Size, style *tcell.Style, selections []string, selectionIndex int) *ComboBox {
	tools.Logger.WithField("module", "combo-box").
		WithField("function", "NewComboBox").
		Infof("%s %s %s %+v", name, position.ToString(), size.ToString(), selections)
	selectionsLength := len(selections)
	// Add padding for every menu item to fill the whole horizontal length.
	paddingSelections := make([]string, selectionsLength)
	for i, s := range selections {
		paddingSelections[i] = fmt.Sprintf("%-*s", size.W-2, s)
	}
	//tools.Logger.WithField("module", "combo-box").
	//    WithField("function", "NewComboBox").
	//    Debugf("%s %+v", name, paddingSelections)
	comboBox := &ComboBox{
		Widget:         NewWidget(name, position, size, style),
		selections:     paddingSelections,
		filtered:       paddingSelections,
		selectionIndex: selectionIndex,
		scroller:       nil,
		inputStr:       "",
	}
	comboBox.scroller = NewVerticalScroller(selectionsLength, size.H-3)
	//tools.Logger.WithField("module", "combo-box").
	//    WithField("function", "NewComboBox").
	//    Debugf("scroller %s", comboBox.scroller.ToString())
	comboBox.SetFocusType(engine.SingleFocus)
	comboBox.SetFocusEnable(true)
	comboBox.updateCanvas()
	return comboBox
}

// -----------------------------------------------------------------------------
// ComboBox private methods
// -----------------------------------------------------------------------------

// execute method runs functionality based on the input from the keyboard, like
// moving selection or entering chracters in the input string.
func (c *ComboBox) execute(args ...any) {
	tools.Logger.WithField("module", "combo-box").WithField("method", "execute").Debugf("%s %+v", c.GetName(), args)
	switch args[0].(string) {
	case "up":
		if c.selectionIndex > 0 {
			c.selectionIndex--
			c.updateCanvas()
		}
	case "down":
		if c.selectionIndex < (len(c.filtered) - 1) {
			c.selectionIndex++
			c.updateCanvas()
		}
	case "run":
	}
}

// updateCanvas method updates the combo box canvas with proper selections to be
// displayed and the proper selected option.
func (c *ComboBox) updateCanvas() {
	// update the scroller with the selection index.
	c.scroller.Update(c.selectionIndex)
	canvas := c.GetCanvas()
	c.scroller.CreateIter()
	canvas.WriteRectangleInCanvasAt(nil, nil, c.GetStyle(), engine.CanvasRectSingleLine)
	canvas.WriteStringInCanvasAt(c.inputStr, c.GetStyle(), api.NewPoint(1, 1))
	isEmptyFilter := (len(c.filtered) == 0)
	for i, x, y := 0, 1, 2; c.scroller.IterHasNext() || i < c.GetSize().H-3; i, y = i+1, y+1 {
		index, _ := c.scroller.IterGetNext()
		var selection string
		if !isEmptyFilter && index <= c.scroller.EndSelection {
			selection = c.filtered[index]
		} else {
			selection = strings.Repeat(" ", c.GetSize().W-2)
		}
		//tools.Logger.WithField("module", "combo-box").
		//    WithField("function", "updateCanvas").
		//    Debugf("%d:%d '%s'", index, i, selection)
		if !isEmptyFilter && index == c.selectionIndex {
			reverseStyle := tools.ReverseStyle(c.GetStyle())
			canvas.WriteStringInCanvasAt(selection, reverseStyle, api.NewPoint(x, y))
		} else {
			canvas.WriteStringInCanvasAt(selection, c.GetStyle(), api.NewPoint(x, y))
		}
	}
}

// updateInputStr method update the user input string and proceed to filter the
// list of options.
func (c *ComboBox) updateInputStr() {
	if c.updateFilterOptions() {
		c.selectionIndex = 0
		c.scroller = NewVerticalScroller(len(c.filtered), c.GetSize().H-3)
		//tools.Logger.WithField("module", "combo-box").
		//    WithField("function", "NewComboBox").
		//    Debugf("scroller %s", c.scroller.ToString())
	}
}

// updateFilterOptions method filters list of options based on the user input
// string.
func (c *ComboBox) updateFilterOptions() bool {
	if len(c.inputStr) == 0 {
		c.filtered = c.selections
		return true
	} else {
		c.filtered = []string{}
		for _, selection := range c.selections {
			if strings.HasPrefix(selection, c.inputStr) {
				c.filtered = append(c.filtered, selection)
			}
		}
		return true
	}
}

// -----------------------------------------------------------------------------
// ComboBox public methods
// -----------------------------------------------------------------------------

// GetSelection method returns the option for the selected index.
func (c *ComboBox) GetSelection() string {
	return strings.TrimSpace(c.selections[c.selectionIndex])
}

// Update method executes all combobox functionality every tick time. Keyboard
// inut is scanned in order to move the selection index and proceed to select
// any option.
func (c *ComboBox) Update(event tcell.Event, scene engine.IScene) {
	defer c.Entity.Update(event, scene)
	if !c.HasFocus() {
		return
	}
	switch ev := event.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyUp:
			c.execute("up")
		case tcell.KeyDown:
			c.execute("down")
		case tcell.KeyEnter:
			c.execute("run")
		case tcell.KeyDEL:
			fallthrough
		case tcell.KeyBackspace:
			if lenInputStr := len(c.inputStr); lenInputStr > 0 {
				c.inputStr = c.inputStr[:lenInputStr-1]
			}
			c.updateInputStr()
			c.updateCanvas()
		case tcell.KeyRune:
			inputRune := string(ev.Rune())
			c.inputStr += inputRune
			c.updateInputStr()
			c.updateCanvas()
		}
	}
}

var _ engine.IObject = (*ComboBox)(nil)
var _ engine.IFocus = (*ComboBox)(nil)
var _ engine.IEntity = (*ComboBox)(nil)
