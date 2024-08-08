// location.go contains all data and logic related with any location present in
// the storyboard.
package storyboard

// -----------------------------------------------------------------------------
//
// ILocation
//
// -----------------------------------------------------------------------------

type ILocation interface {
	GetDescription() string
	GetKeyEvents() []IEvent
	GetName() string
	SetDescription(string)
	SetKeyEvents([]IEvent)
	SetName(string)
}

// -----------------------------------------------------------------------------
//
// Location
//
// -----------------------------------------------------------------------------

type Location struct {
	description string
	keyEvents   []IEvent
	name        string
}

// -----------------------------------------------------------------------------
// Location public methods
// -----------------------------------------------------------------------------

func (l *Location) GetDescription() string {
	return l.description
}

func (l *Location) GetKeyEvents() []IEvent {
	return l.keyEvents
}

func (l *Location) GetName() string {
	return l.name
}

func (l *Location) SetDescription(description string) {
	l.description = description
}

func (l *Location) SetKeyEvents(keyEvents []IEvent) {
	l.keyEvents = keyEvents
}

func (l *Location) SetName(name string) {
	l.name = name
}

var _ ILocation = (*Location)(nil)
