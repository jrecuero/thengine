// proficiency.go package
package rules

type ProficiencyType string

const (
	ToolProficiency        ProficiencyType = "tool"
	WeaponProficiency      ProficiencyType = "weapon"
	SkillProficiency       ProficiencyType = "skill"
	SavingThrowProficiency ProficiencyType = "saving throw"
	ArmorProficiency       ProficiencyType = "armor"
)

// -----------------------------------------------------------------------------
//
// IProficiency
//
// -----------------------------------------------------------------------------

type IProficiency interface {
	GetDescription() string
	GetName() string
	GetType() ProficiencyType
	SetDescription(string)
	SetName(string)
	SetType(ProficiencyType)
}

// -----------------------------------------------------------------------------
//
// Proficiency
//
// -----------------------------------------------------------------------------

type Proficiency struct {
	description string
	name        string
	ptype       ProficiencyType
}

// NewProficiency function creates a new Proficiency instance.
func NewProficiency(name string, ptype ProficiencyType) *Proficiency {
	p := &Proficiency{
		description: name,
		name:        name,
		ptype:       ptype,
	}
	return p
}

// -----------------------------------------------------------------------------
// Proficiency public methods
// -----------------------------------------------------------------------------

func (p *Proficiency) GetDescription() string {
	return p.description
}

func (p *Proficiency) GetName() string {
	return p.name
}

func (p *Proficiency) GetType() ProficiencyType {
	return p.ptype
}

func (p *Proficiency) SetDescription(description string) {
	p.description = description
}

func (p *Proficiency) SetName(name string) {
	p.name = name
}

func (p *Proficiency) SetType(ptype ProficiencyType) {
	p.ptype = ptype
}

var _ IProficiency = (*Proficiency)(nil)
