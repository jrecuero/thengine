// textInput.go contains all attributes and methods required to implement a
// generic text box input widget.
package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// TextInput
//
// -----------------------------------------------------------------------------

// TextInput structure defines all attributes and methods required for any
// generic text input widget.
type TextInput struct {
	*Widget
	inputStr string
}

// NewTextInput function creates a new TextInput instance widget.
func NewTextInput(name string, position *api.Point, size *api.Size, style *tcell.Style, defaultStr string) *TextInput {
	textInput := &TextInput{
		Widget:   NewWidget(name, position, size, style),
		inputStr: defaultStr,
	}
	textInput.updateCanvas()
	textInput.SetFocusType(engine.SingleFocus)
	textInput.SetFocusEnable(true)
	return textInput
}

// -----------------------------------------------------------------------------
// TextInput private methods
// -----------------------------------------------------------------------------

// updateCanvas method updates the text inputwidget canvas with the string
// information.
func (t *TextInput) updateCanvas() {
	//canvas := engine.NewCanvasFromString(t.inputStr, t.GetStyle())
	//t.SetCanvas(canvas)
	canvas := t.GetCanvas()
	cell := engine.NewCell(t.GetStyle(), ' ')
	canvas.FillWithCell(cell)
	canvas.WriteStringInCanvas(t.inputStr, t.GetStyle())
}

// -----------------------------------------------------------------------------
// TextInput public methods
// -----------------------------------------------------------------------------

// AcquireFocus method acquires focus for the entity.
func (t *TextInput) AcquireFocus() (bool, error) {
	tools.Logger.WithField("module", "input-text").WithField("function", "AcquireFocus").Infof("%s", t.GetName())
	ok, err := t.Entity.AcquireFocus()
	if err == nil {
		lenInputStr := len(t.inputStr)
		col := t.GetPosition().X + lenInputStr
		row := t.GetPosition().Y
		display := engine.GetEngine().GetDisplay()
		display.ShowCursor(col, row)
	}
	return ok, err
}

// GetInputText method returns the input text string.
func (t *TextInput) GetInputText() string {
	return t.inputStr
}

// ReleaseFocus method release the focus for the entity.
func (t *TextInput) ReleaseFocus() (bool, error) {
	ok, err := t.Entity.ReleaseFocus()
	if err == nil {
		display := engine.GetEngine().GetDisplay()
		display.HideCursor()
	}
	return ok, err
}

// SetInputText method sets the input text string.
func (t *TextInput) SetInputText(str string) {
	t.inputStr = str
	t.updateCanvas()
}

// Update method runs every cycle to update the text input.
func (t *TextInput) Update(event tcell.Event) {
	tools.Logger.WithField("module", "input-text").WithField("function", "Update").Debugf("%s %+v", t.GetName(), event)
}

var _ engine.IEntity = (*TextInput)(nil)
