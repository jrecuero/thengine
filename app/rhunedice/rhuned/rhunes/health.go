package rhunes

import "github.com/jrecuero/thengine/app/rhunedice/rhuned/api"

func NewHealthRhune() *api.Rhune {
	return api.NewRhune(
		api.HealthName,
		api.HealthShort,
		"Health rhune is used to measure avatar life",
		api.BaseRhune,
		api.HealthBucket,
		nil)
}
