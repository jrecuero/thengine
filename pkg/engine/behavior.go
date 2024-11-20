package engine

// -----------------------------------------------------------------------------
// Package constants
// -----------------------------------------------------------------------------
const (
	BehaviorConsume string = "consume"
	BehaviorDraw    string = "draw"
	BehaviorInit    string = "init"
	BehaviorNotify  string = "notify"
	BehaviorStart   string = "start"
	BehaviorStop    string = "stop"
	BehaviorUpdate  string = "update"
)

// -----------------------------------------------------------------------------
//
// IBehavior
//
// -----------------------------------------------------------------------------

type IBehavior interface {
	GetBehaviorFor(string) any
	SetBehaviorFor(string, any)
}

// -----------------------------------------------------------------------------
//
// Behavior
//
// -----------------------------------------------------------------------------

type Behavior struct {
	behaviors map[string]any
}

// -----------------------------------------------------------------------------
// New Behavior functions
// -----------------------------------------------------------------------------

// NewBehavior function creates a new Behavior instance.
func NewBehavior() *Behavior {
	return &Behavior{
		behaviors: make(map[string]any),
	}
}

// -----------------------------------------------------------------------------
// Behavior public methods
// -----------------------------------------------------------------------------

func (b *Behavior) GetBehaviorFor(name string) any {
	return b.behaviors[name]
}

func (b *Behavior) SetBehaviorFor(name string, f any) {
	b.behaviors[name] = f
}

var _ IBehavior = (*Behavior)(nil)
