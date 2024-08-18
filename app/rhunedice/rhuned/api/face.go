package api

import "fmt"

var (
	AttackFace  = NewFace(AttackRhune)
	DefenseFace = NewFace(DefenseRhune)
	SkillFace   = NewFace(SkillRhune)
)

type IFace interface {
	GetRhune() IRhune
	String() string
}

type Face struct {
	rhune IRhune
}

func NewFace(rhune IRhune) *Face {
	return &Face{
		rhune: rhune,
	}
}

func (f *Face) GetRhune() IRhune {
	return f.rhune
}

func (f *Face) String() string {
	return fmt.Sprintf(f.rhune.GetShort())
}

var _ IFace = (*Face)(nil)
