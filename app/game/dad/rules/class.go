// class.go
package rules

// -----------------------------------------------------------------------------
//
// IClass
//
// -----------------------------------------------------------------------------

type IClass interface {
	GetDescription() string
	GetFeats() []IFeat
	GetHitDie() IDiceThrow
	GetName() string
	GetTraits() []ITrait
	RollHitDie() int
	SetDescription(string)
	SetFeats([]IFeat)
	SetHitDie(IDiceThrow)
	SetName(string)
	SetTraits([]ITrait)
}

// -----------------------------------------------------------------------------
//
// Class
//
// -----------------------------------------------------------------------------

type Class struct {
	description string
	feats       []IFeat
	hitDie      IDiceThrow
	name        string
	traits      []ITrait
}

// NewClass function creates a new Class instance.
func NewClass(name string, hitDie IDiceThrow) *Class {
	c := &Class{
		description: name,
		feats:       nil,
		hitDie:      hitDie,
		name:        name,
		traits:      nil,
	}
	return c
}

// -----------------------------------------------------------------------------
// Class public methods
// -----------------------------------------------------------------------------

func (c *Class) GetDescription() string {
	return c.description
}

func (c *Class) GetFeats() []IFeat {
	return c.feats
}

func (c *Class) GetHitDie() IDiceThrow {
	return c.hitDie
}

func (c *Class) GetName() string {
	return c.name
}

func (c *Class) GetTraits() []ITrait {
	return c.traits
}

func (c *Class) RollHitDie() int {
	if c.hitDie != nil {
		return c.hitDie.SureRoll()
	}
	return 1
}

func (c *Class) SetDescription(description string) {
	c.description = description
}

func (c *Class) SetFeats(feats []IFeat) {
	c.feats = feats
}

func (c *Class) SetHitDie(hitDie IDiceThrow) {
	c.hitDie = hitDie
}

func (c *Class) SetName(name string) {
	c.name = name
}

func (c *Class) SetTraits(traits []ITrait) {
	c.traits = traits
}

var _ IClass = (*Class)(nil)
