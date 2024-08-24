package rhunes

import "github.com/jrecuero/thengine/app/rhunedice/rhuned/api"

func NewAttackRhune() *api.Rhune {
	return api.NewRhune(
		api.AttackName,
		api.AttackShort,
		"Attack rhune is used to damage",
		api.BaseRhune,
		nil)
}
