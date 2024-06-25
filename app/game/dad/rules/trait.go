// trait.go
package rules

// -----------------------------------------------------------------------------
//
// ITrait
//
// -----------------------------------------------------------------------------

type ITrait interface {
	IActivable
	GetDescription() string
	GetName() string
	SetDescription(string)
	SetName(string)
}

// -----------------------------------------------------------------------------
//
// Trait
//
// -----------------------------------------------------------------------------

type Trait struct {
	*Activable
	description string
	name        string
}

// NewTrait function creates a new Trait instance.
func NewTrait(name string, ispassive bool, issustained bool, isactivated bool) *Trait {
	t := &Trait{
		Activable: NewActivable(ispassive, issustained, isactivated),
		name:      name,
	}
	return t
}

// -----------------------------------------------------------------------------
// Trait public methods
// -----------------------------------------------------------------------------

func (t *Trait) GetDescription() string {
	return t.description
}

func (t *Trait) GetName() string {
	return t.name
}

func (t *Trait) SetDescription(description string) {
	t.description = description
}

func (t *Trait) SetName(name string) {
	t.name = name
}

var _ IActivable = (*Trait)(nil)
var _ ITrait = (*Trait)(nil)
