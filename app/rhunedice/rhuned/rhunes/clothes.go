package rhunes

import (
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/equipment/armors"
)

func NewClothesRhune() *api.Rhune {
	clothesRhune := api.NewRhune(
		api.ClothesName,
		api.ClothesShort,
		"Clothes armor rhune",
		api.ExtraRhune,
		api.ExtraBucket,
		func(avatar api.IAvatar) (any, error) {
			clothes := armors.NewClothes()
			avatar.GetEquipment().SetArmor(clothes)
			return nil, nil
		},
	)
	return clothesRhune
}
