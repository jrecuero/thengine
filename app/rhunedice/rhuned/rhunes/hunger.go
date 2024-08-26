package rhunes

import "github.com/jrecuero/thengine/app/rhunedice/rhuned/api"

func NewHungerRhune() *api.Rhune {
	return api.NewRhune(
		api.HungerName,
		api.HungerShort,
		"Hunger rhune is used to measure avatar hunger and thrist",
		api.BaseRhune,
		api.HungerBucket,
		nil)
}
