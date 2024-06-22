// trait.go
package rules

// -----------------------------------------------------------------------------
//
// ITrait
//
// -----------------------------------------------------------------------------

type ITrait interface {
	GetBenefits() []any
	GetDescription() string
	GetName() string
	SetBenefits([]any)
	SetDescription(string)
	SetName(string)
}

// -----------------------------------------------------------------------------
//
// Trait
//
// -----------------------------------------------------------------------------

type Trait struct {
	benefits    []any
	description string
	name        string
}

// NewTrait function creates a new Trait instance.
func NewTrait(name string) *Trait {
	t := &Trait{
		benefits:    nil,
		description: name,
		name:        name,
	}
	return t
}

// -----------------------------------------------------------------------------
// Trait public methods
// -----------------------------------------------------------------------------

func (t *Trait) GetBenefits() []any {
	return t.benefits
}

func (t *Trait) GetDescription() string {
	return t.description
}

func (t *Trait) GetName() string {
	return t.name
}

func (t *Trait) SetBenefits(benefits []any) {
	t.benefits = benefits
}

func (t *Trait) SetDescription(description string) {
	t.description = description
}

func (t *Trait) SetName(name string) {
	t.name = name
}

var _ ITrait = (*Trait)(nil)
