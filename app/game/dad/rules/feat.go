// feat.go module provides structures and methods to represent and manage
// feats.
// A feat is a special feature that represents a character's unique abilities
// or training. Feats provide characters with new capabilities or enhance
// existing ones, allowing for greater customization and specialization.
// They are an optional rule that can be chosen instead of an Ability Score
// Improvement at certain levels.
package rules

// -----------------------------------------------------------------------------
//
// IFeat
//
// -----------------------------------------------------------------------------

// IFeat interface provides all meethods any feat have to implement.
type IFeat interface {
	GetDescription() string
	GetEffects() map[string]any
	GetName() string
	GetPrerequisites() []any
	MeetPrerequisites(IUnit) bool
	RollEffects(IUnit)
	SetDescription(string)
	SetEffects(map[string]any)
	SetName(string)
	SetPrerequisites([]any)
}

// -----------------------------------------------------------------------------
//
// Feat
//
// -----------------------------------------------------------------------------

// Feat structure represents all attributes and methods for any generic feat.
type Feat struct {
	description   string
	effects       map[string]any
	name          string
	prerequisites []any
}

// NewFeat function creates a new Feat instance.
func NewFeat(name string) *Feat {
	f := &Feat{
		description:   name,
		effects:       make(map[string]any),
		name:          name,
		prerequisites: nil,
	}
	return f
}

// -----------------------------------------------------------------------------
// Feat public methods
// -----------------------------------------------------------------------------

func (f *Feat) GetDescription() string {
	return f.description
}

func (f *Feat) GetEffects() map[string]any {
	return f.effects
}

func (f *Feat) GetName() string {
	return f.name
}

func (f *Feat) GetPrerequisites() []any {
	return f.prerequisites
}

func (f *Feat) MeetPrerequisites(unit IUnit) bool {
	return false
}

func (f *Feat) RollEffects(unit IUnit) {
}

func (f *Feat) SetDescription(description string) {
	f.description = description
}

func (f *Feat) SetEffects(effects map[string]any) {
	f.effects = effects
}

func (f *Feat) SetName(name string) {
	f.name = name
}

func (f *Feat) SetPrerequisites(prerequisites []any) {
	f.prerequisites = prerequisites
}

var _ IFeat = (*Feat)(nil)
