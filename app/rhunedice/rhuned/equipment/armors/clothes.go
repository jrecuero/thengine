package armors

import "github.com/jrecuero/thengine/app/rhunedice/rhuned/api"

func NewClothes() *api.Armor {
	clothes := api.NewArmor(
		api.ClothesName,
		"clothes are the most basic armor",
		api.DefenseName,
		1)
	return clothes
}
