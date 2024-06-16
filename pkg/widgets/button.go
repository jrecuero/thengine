// button.go contains all attributes and methods required for handling basic
// buttons.
package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// Button
//
// -----------------------------------------------------------------------------

// Button structure defines a baseline for any button entity.
// Entity: contains an Entity instance.
// Focus: contains a Focus instance.
// Label: label of the button to be displayed on the screen.
// clicked: flag to indicate if the button was clicked.
// command: instance of the command to be executed.
// args: arguments to pass to the command to be executed.
type Button struct {
	*Widget
	label   string
	clicked bool
}

// NewButton function creates a new Button instance with the given name,
// position, size and foreground and background colors.
func NewButton(name string, position *api.Point, size *api.Size, style *tcell.Style, label string) *Button {
	button := &Button{
		Widget:  NewWidget(name, position, size, style),
		label:   label,
		clicked: false,
	}
	button.SetFocusType(engine.SingleFocus)
	button.SetFocusEnable(true)
	button.updateCanvas()
	return button
}

// -----------------------------------------------------------------------------
// Button private methods
// -----------------------------------------------------------------------------

func (b *Button) execute(args ...any) {
	tools.Logger.WithField("module", "button").
		WithField("method", "execute").
		Debugf("%s %+v", b.GetName(), args)
	if b.GetWidgetCallback() != nil {
		b.GetWidgetCallback()(b, args...)
	}
}

// updateCanvas method updates the buttonwidget canvas with the label
// information.
func (b *Button) updateCanvas() {
	canvas := engine.NewCanvasFromString(b.label, b.GetStyle())
	b.SetCanvas(canvas)
}

// -----------------------------------------------------------------------------
// Button public methods
// -----------------------------------------------------------------------------

// AcquireFocus method acquires focus for the entity.
func (b *Button) AcquireFocus() (bool, error) {
	tools.Logger.WithField("module", "button").
		WithField("method", "AcquireFocus").
		Debugf("%s", b.GetName())
	ok, err := b.Entity.AcquireFocus()
	if err == nil {
		reverseStyle := tools.ReverseStyle(b.GetStyle())
		b.SetStyle(reverseStyle)
	}
	return ok, err
}

// GetLabel method returns the Button instance string.
func (t *Button) GetLabel() string {
	return t.label
}

// SetLabel method sets the Button instance string.
func (t *Button) SetLabel(label string) {
	t.label = label
	t.updateCanvas()
}

func (b *Button) Refresh() {
	b.updateCanvas()
}

// ReleaseFocus method release the focus for the entity.
func (b *Button) ReleaseFocus() (bool, error) {
	ok, err := b.Entity.ReleaseFocus()
	if err == nil {
		reverseStyle := tools.ReverseStyle(b.GetStyle())
		b.SetStyle(reverseStyle)
	}
	return ok, err
}

// Update method executes all the button functionality every tick time. Button
// callback function will be called if the button was clicked.
func (b *Button) Update(event tcell.Event, scene engine.IScene) {
	defer b.Entity.Update(event, scene)
	if !b.HasFocus() {
		return
	}
	actions := []*KeyboardAction{
		{
			Key:      tcell.KeyEnter,
			Callback: b.execute,
			Args:     b.GetWidgetCallbackArgs(),
		},
	}
	b.HandleKeyboardForActions(event, actions)
}

var _ engine.IObject = (*Button)(nil)
var _ engine.IFocus = (*Button)(nil)
var _ engine.IEntity = (*Button)(nil)
