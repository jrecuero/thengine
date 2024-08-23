package faces

import (
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/rhunes"
)

var (
	AttackFace  = api.NewFace(rhunes.AttackRhune)
	DefenseFace = api.NewFace(rhunes.DefenseRhune)
	SkillFace   = api.NewFace(rhunes.SkillRhune)
)
