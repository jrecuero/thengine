package api

var (
	AttackRhune  = NewAttackRhune()
	DefenseRhune = NewDefenseRhune()
	SkillRhune   = NewSkillRhune()
)

type IRhune interface {
	Execute(IAvatar) (any, error)
	GetCat() IComparable
	GetDescription() string
	GetName() string
	GetShort() string
}

type Rhune struct {
	cat         ERhuneCat
	description string
	execute     func(IAvatar) (any, error)
	name        string
	short       string
}

func NewAttackRhune() *Rhune {
	return &Rhune{
		cat:         BaseRhune,
		description: "Attack rhune is used to damage",
		execute:     nil,
		name:        "attack",
		short:       "ATK",
	}
}

func NewDefenseRhune() *Rhune {
	return &Rhune{
		cat:         BaseRhune,
		description: "Base rhune is used to defense against damage",
		execute:     nil,
		name:        "defense",
		short:       "DEF",
	}
}

func NewSkillRhune() *Rhune {
	return &Rhune{
		cat:         BaseRhune,
		description: "Skill rhune provides avatar skill abilities",
		execute:     nil,
		name:        "skill",
		short:       "SKL",
	}
}

func (r *Rhune) Execute(avatar IAvatar) (any, error) {
	if r.execute != nil {
		return r.execute(avatar)
	}
	return nil, nil
}

func (r *Rhune) GetCat() IComparable {
	return r.cat
}

func (r *Rhune) GetDescription() string {
	return r.description
}

func (r *Rhune) GetName() string {
	return r.name
}

func (r *Rhune) GetShort() string {
	return r.short
}

var _ IRhune = (*Rhune)(nil)
