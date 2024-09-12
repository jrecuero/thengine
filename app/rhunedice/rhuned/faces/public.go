package faces

import (
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/rhunes"
)

var (
	// Base faces
	AttackFace  = api.NewFace(rhunes.AttackRhune)
	DefenseFace = api.NewFace(rhunes.DefenseRhune)
	SkillFace   = api.NewFace(rhunes.SkillRhune)
	StaminaFace = api.NewFace(rhunes.StaminaRhune)
	HealthFace  = api.NewFace(rhunes.HealthRhune)
	StepFace    = api.NewFace(rhunes.StepRhune)
	NilFace     = api.NewFace(rhunes.NilRhune)

	// Extra faces
	ClothesFace   = api.NewFace(rhunes.ClothesRhune)
	KnowledgeFace = api.NewFace(rhunes.KnowledgeRhune)
)
