package dice

import (
	"fmt"

	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/faces"
)

func NewDefaultBaseDice(name string) *api.Dice {
	diceFaces := []api.IFace{
		faces.AttackFace,
		faces.DefenseFace,
		faces.SkillFace,
		faces.StaminaFace,
		faces.HealthFace,
		faces.StepFace,
		faces.HungerFace,
		faces.NilFace,
	}
	diceName := fmt.Sprintf("dice/base/%s", name)
	dice := api.NewDice(diceName, diceFaces)
	return dice
}

func NewDefaultExtraDice(name string) *api.Dice {
	diceFaces := []api.IFace{
		faces.NilFace,
	}
	dice := api.NewDice(name, diceFaces)
	return dice
}
