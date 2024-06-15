// dialog.go module contains all required data and logix to create a dialog
// widget that can contains:
// - decorated box
// - list of text and text input aligned
// - list of buttons
package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// Dialog
//
// -----------------------------------------------------------------------------

type Dialog struct {
	*engine.ObjectUI
	background engine.IEntity
	texts      []*Text
	inputs     []*TextInput
	buttons    []*Button
	scene      engine.IScene
	outputs    []string
	opts       []any
}

func NewDialog(name string, position *api.Point, size *api.Size, style *tcell.Style, scene engine.IScene,
	texts []*Text, inputs []*TextInput, buttons []*Button, opts ...any) *Dialog {
	d := &Dialog{
		ObjectUI:   engine.NewObjectUI(name, position, size, style),
		background: engine.NewEntity("dialog/background", position, size, style),
		texts:      texts,
		inputs:     inputs,
		buttons:    buttons,
		scene:      scene,
		outputs:    nil,
		opts:       opts,
	}

	d.background.GetCanvas().FillWithCell(engine.NewCell(style, ' '))
	d.background.GetCanvas().WriteRectangleInCanvasAt(nil, nil, d.GetStyle(), engine.CanvasRectSingleLine)
	scene.AddEntity(d.background)

	maxW := d.populateTexts()

	d.populateTextInputs(maxW)

	d.populateButtons()

	return d
}

// -----------------------------------------------------------------------------
// Dialog private methods.
// -----------------------------------------------------------------------------

func (d *Dialog) populateButtons() {
	x, y := d.GetPosition().X, d.GetPosition().Y+d.GetSize().H-2
	var spacePerLabel int
	if len(d.GetButtons()) != 0 {
		spacePerLabel = d.GetSize().W / len(d.GetButtons())
	}
	nextW := 0
	for _, button := range d.GetButtons() {
		w := len(button.GetLabel())
		padding := (spacePerLabel - w) / 2
		button.SetPosition(api.NewPoint(x+nextW+padding, y))
		button.SetSize(api.NewSize(w, 3))
		button.SetCanvas(engine.NewCanvas(button.GetSize()))
		if button.GetStyle() == nil {
			button.SetStyle(d.GetStyle())
		}
		//tools.Logger.WithField("module", "dialog").
		//    WithField("method", "NewDialog").
		//    Debugf("x:%d len(label):%d spacePerLabel:%d padding:%d nextW:%d",
		//        x, w, spacePerLabel, padding, nextW)
		nextW += spacePerLabel
		button.Refresh()
		d.GetScene().AddEntity(button)
	}
}

func (d *Dialog) populateTexts() int {

	x, y := d.GetPosition().X+1, d.GetPosition().Y+1
	maxW := 0
	for i, text := range d.GetTexts() {
		text.SetPosition(api.NewPoint(x, y+i))
		w := len(text.GetText())
		maxW = tools.Max(maxW, w)
		text.SetSize(api.NewSize(w, 1))
		text.SetCanvas(engine.NewCanvas(text.GetSize()))
		if text.GetStyle() == nil {
			text.SetStyle(d.GetStyle())
		}
		text.Refresh()
		d.GetScene().AddEntity(text)
	}
	return maxW
}

func (d *Dialog) populateTextInputs(maxW int) {
	maxW += 2
	x, y := d.GetPosition().X+maxW, d.GetPosition().Y+1
	for i, input := range d.GetInputs() {
		input.SetPosition(api.NewPoint(x, y+i))
		w := d.GetSize().W - maxW - 2
		input.SetSize(api.NewSize(w, 1))
		input.SetCanvas(engine.NewCanvas(input.GetSize()))
		if style := input.GetStyle(); style == nil {
			style = tools.ReverseStyle(d.GetStyle())
			input.SetStyle(style)
		}
		input.Refresh()
		d.GetScene().AddEntity(input)
	}
}

// -----------------------------------------------------------------------------
// Dialog public methods.
// -----------------------------------------------------------------------------

func (d *Dialog) Close() []string {
	return d.GetOutputs()
}

func (d *Dialog) GetButtons() []*Button {
	return d.buttons
}

func (d *Dialog) GetInputs() []*TextInput {
	return d.inputs
}

func (d *Dialog) GetOpts() []any {
	return d.opts
}

func (d *Dialog) GetOutputs() []string {
	for _, input := range d.GetInputs() {
		d.outputs = append(d.outputs, input.GetInputText())
	}
	return d.outputs
}

func (d *Dialog) GetScene() engine.IScene {
	return d.scene
}

func (d *Dialog) GetTexts() []*Text {
	return d.texts
}

//func (d *Dialog) SetButtons(buttons []*Button) {
//    d.buttons = buttons
//}

//func (d *Dialog) SetInputs(inputs []*TextInput) {
//    d.inputs = inputs
//}

//func (d *Dialog) SetOpts(opts []any) {
//    d.opts = opts
//}

//func (d *Dialog) SetTexts(texts []*Text) {
//    d.texts = texts
//}
