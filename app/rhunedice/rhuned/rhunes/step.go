package rhunes

import "github.com/jrecuero/thengine/app/rhunedice/rhuned/api"

func NewStepRhune() *api.Rhune {
	return api.NewRhune(
		api.StepName,
		api.StepShort,
		"Step rhune is used to measure how many steps can move the avatar",
		api.BaseRhune,
		api.StepBucket,
		nil)
}
