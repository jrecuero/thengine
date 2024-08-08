// action.go package contains all data and logic for any action involved in a
// storyboard event.
package storyboard

// -----------------------------------------------------------------------------
//
// IAction
//
// -----------------------------------------------------------------------------

type IAction interface {
	GetConsequences() map[string]any
	GetDescription() string
	GetName() string
	SetConsequences(map[string]any)
	SetDescription(string)
	SetName(string)
}

// -----------------------------------------------------------------------------
//
// Action
//
// -----------------------------------------------------------------------------

type Action struct {
	consequences map[string]any
	description  string
	name         string
}

// -----------------------------------------------------------------------------
// Action public methods
// -----------------------------------------------------------------------------

func (a *Action) GetConsequences() map[string]any {
	return a.consequences
}

func (a *Action) GetDescription() string {
	return a.description
}

func (a *Action) GetName() string {
	return a.name
}

func (a *Action) SetConsequences(consequences map[string]any) {
	a.consequences = consequences
}

func (a *Action) SetDescription(description string) {
	a.description = description
}

func (a *Action) SetName(name string) {
	a.name = name
}

var _ IAction = (*Action)(nil)
