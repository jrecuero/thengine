package rhunes

import "github.com/jrecuero/thengine/app/rhunedice/rhuned/api"

func NewSkillRhune() *api.Rhune {
	return api.NewRhune(
		api.SkillName,
		api.SkillShort,
		"Skill rhune provides avatar skill abilities",
		api.BaseRhune,
		nil)
}
