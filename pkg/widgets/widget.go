// widget.go contains all common data and method for any widget instance
package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

// -----------------------------------------------------------------------------
// Package public types
// -----------------------------------------------------------------------------

// WidgetCallback type is the type for widget callback.
type WidgetCallback func(entity engine.IEntity, args ...any) bool

// WidgetArgs type is the type for the list of arguments passed to any widget
// callback.
type WidgetArgs []any

// KeyboardAction structure identifies the information to be passed for
// handling keyboard input actions.
type KeyboardAction struct {
	Key      tcell.Key
	Rune     rune
	Callback func(...any)
	Args     []any
}

// -----------------------------------------------------------------------------
//
// Widget
//
// -----------------------------------------------------------------------------

// Wdiget structure defines all attributes and method for any basic and common
// widget.
type Widget struct {
	*engine.Entity
	callback     WidgetCallback
	callbackArgs WidgetArgs
}

// NewWidget function creates a new Widget instance.
func NewWidget(name string, position *api.Point, size *api.Size, style *tcell.Style) *Widget {
	return &Widget{
		Entity: engine.NewEntity(name, position, size, style),
	}
}

// NewNamedWidget function creates a new Widget instance with only a name.
func NewNamedWidget(name string) *Widget {
	return &Widget{
		Entity: engine.NewNamedEntity(name),
	}
}

// -----------------------------------------------------------------------------
// Widget public methods
// -----------------------------------------------------------------------------

// GetWidgetCallback method returns the widget callback function.
func (w *Widget) GetWidgetCallback() WidgetCallback {
	return w.callback
}

// GetWidgetCallbackArgs method returns the widget callback arguments.
func (w *Widget) GetWidgetCallbackArgs() WidgetArgs {
	return w.callbackArgs
}

// HandleKeyboardInputForActions method handles keyboard inputs related with
// the given information provided. Input parameters provide the keys that have
// to be handled and the callbacks for each of them.
func (w *Widget) HandleKeyboardForActions(event tcell.Event, actions []*KeyboardAction) {
	switch ev := event.(type) {
	case *tcell.EventKey:
		for _, action := range actions {
			if (action.Key != 0) && (ev.Key() == action.Key) {
				action.Callback(action.Args...)
				return
			} else if (action.Rune != 0) && (ev.Rune() == action.Rune) {
				action.Callback(action.Args...)
				return
			}
		}
	}
}

// HandleKeyboardInputForString method handles keyboard input affecting the given
// string.
// - Any rune entered is added to the string.
// - Any delete removes the last characted in the string.
func (w *Widget) HandleKeyboardInputForString(event tcell.Event, str string) (string, bool, bool) {
	switch ev := event.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyDEL:
			fallthrough
		case tcell.KeyBackspace:
			if lenInputStr := len(str); lenInputStr > 0 {
				str = str[:lenInputStr-1]
				return str, true, false
			}
		case tcell.KeyRune:
			inputRune := string(ev.Rune())
			str += inputRune
			return str, true, false
		case tcell.KeyEnter:
			return str, true, true
		default:
		}
	}
	return str, false, false
}

// RunCallback method executes the widget callback. If not any arguments are
// provided, it used those in the widget arguments attributes.
func (w *Widget) RunCallback(args ...any) bool {
	if w.callback == nil {
		return true
	}
	if len(args) == 0 {
		args = w.callbackArgs
	}
	return w.callback(w, args...)
}

// SetWidgetCallback method sets a new widget callback and arguments to be
// passed.
func (w *Widget) SetWidgetCallback(f WidgetCallback, args ...any) {
	w.callback = f
	w.callbackArgs = args
}

// SetWidgetCallbackArgs method sets a new set of arguments to be passed to the
// widget callback.
func (w *Widget) SetWidgetCallbackArgs(args ...any) {
	w.callbackArgs = args
}

var _ engine.IEntity = (*Widget)(nil)
