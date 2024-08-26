package rhunes

import "github.com/jrecuero/thengine/app/rhunedice/rhuned/api"

func NewDefenseRhune() *api.Rhune {
	return api.NewRhune(
		api.DefenseName,
		api.DefenseShort,
		"Base rhune is used to defense against damage",
		api.BaseRhune,
		api.DefenseBucket,
		nil)
}
