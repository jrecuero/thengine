package rhunes

import "github.com/jrecuero/thengine/app/rhunedice/rhuned/api"

func NewStaminaRhune() *api.Rhune {
	return api.NewRhune(
		api.StaminaName,
		api.StaminaShort,
		"Stamina rhune is used to measure avatar tireless",
		api.BaseRhune,
		api.StaminaBucket,
		nil)
}
