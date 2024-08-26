package rhunes

import "github.com/jrecuero/thengine/app/rhunedice/rhuned/api"

func NewExtraRhune() *api.Rhune {
	return api.NewRhune(
		api.ExtraName,
		api.ExtraShort,
		"Extra rhune is used for any extra-functionality",
		api.BaseRhune,
		api.ExtraBucket,
		nil)
}
