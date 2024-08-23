package rhunes

import "github.com/jrecuero/thengine/app/rhunedice/rhuned/api"

func NewAttackRhune() *api.Rhune {
	return api.NewRhune(
		"attack",
		"ATK",
		"Attack rhune is used to damage",
		api.BaseRhune,
		nil)
}
