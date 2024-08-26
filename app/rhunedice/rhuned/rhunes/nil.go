package rhunes

import "github.com/jrecuero/thengine/app/rhunedice/rhuned/api"

func NewNilRhune() *api.Rhune {
	return api.NewRhune(
		api.NilName,
		api.NilShort,
		"Nil rhune is an invalid rhune",
		api.BaseRhune,
		api.NilBucket,
		nil)
}
