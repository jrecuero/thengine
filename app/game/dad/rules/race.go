// race.go package
package rules

// -----------------------------------------------------------------------------
//
// IRace
//
// -----------------------------------------------------------------------------

type IRace interface {
	GetDescription() string
	GetFeats() []IFeat
	GetName() string
	GetTraits() []ITrait
	SetDescription(string)
	SetFeats([]IFeat)
	SetName(string)
	SetTraits([]ITrait)
}

// -----------------------------------------------------------------------------
//
// Race
//
// -----------------------------------------------------------------------------

type Race struct {
	description string
	feats       []IFeat
	name        string
	traits      []ITrait
}

// NewRace creates a new Race instance.
func NewRace(name string) *Race {
	r := &Race{
		description: name,
		feats:       nil,
		name:        name,
		traits:      nil,
	}
	return r
}

// -----------------------------------------------------------------------------
// Race public methods
// -----------------------------------------------------------------------------

func (r *Race) GetDescription() string {
	return r.description
}

func (r *Race) GetFeats() []IFeat {
	return r.feats
}

func (r *Race) GetName() string {
	return r.name
}

func (r *Race) GetTraits() []ITrait {
	return r.traits
}

func (r *Race) SetDescription(description string) {
	r.description = description
}

func (r *Race) SetFeats(feats []IFeat) {
	r.feats = feats
}

func (r *Race) SetName(name string) {
	r.name = name
}

func (r *Race) SetTraits(traits []ITrait) {
	r.traits = traits
}

var _ IRace = (*Race)(nil)
