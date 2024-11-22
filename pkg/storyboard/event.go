// event.go package contains all data and logic related with any event in the
// storyboard.
package storyboard

// -----------------------------------------------------------------------------
//
// IEvent
//
// -----------------------------------------------------------------------------

type IEvent interface {
	GetActions() []IAction
	GetDesription() string
	GetDialogs() []IDialog
	GetTitle() string
	SetActions([]IAction)
	SetDesription(string)
	SetDialogs([]IDialog)
	SetTitle(string)
}

// -----------------------------------------------------------------------------
//
// Event
//
// -----------------------------------------------------------------------------

type Event struct {
	actions    []IAction
	desription string
	dialogs    []IDialog
	title      string
}

// -----------------------------------------------------------------------------
// Event public methods
// -----------------------------------------------------------------------------

func (e *Event) GetActions() []IAction {
	return e.actions
}

func (e *Event) GetDesription() string {
	return e.desription
}

func (e *Event) GetDialogs() []IDialog {
	return e.dialogs
}

func (e *Event) GetTitle() string {
	return e.title
}

func (e *Event) SetActions(actions []IAction) {
	e.actions = actions
}

func (e *Event) SetDesription(description string) {
	e.desription = description
}

func (e *Event) SetDialogs(dialogs []IDialog) {
	e.dialogs = dialogs
}

func (e *Event) SetTitle(title string) {
	e.title = title
}

var _ IEvent = (*Event)(nil)
