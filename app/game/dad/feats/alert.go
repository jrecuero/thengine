// alert.go package implement Alert feat for any character.
package feats

import (
	"github.com/jrecuero/thengine/app/game/dad/constants"
	"github.com/jrecuero/thengine/app/game/dad/rules"
)

// -----------------------------------------------------------------------------
//
// Alert
//
// -----------------------------------------------------------------------------

func NewAlert() *rules.Feat {
	f := rules.NewFeat("alert")
	f.SetDescription(`Always on the lookout for danger, you gain the following
benefits:
- You gain a +5 bonus to initiative.
x You can't be surprised while you are conscious.
x Other creatures don’t gain advantage on attack rolls against you as a result
of being unseen by you.`)
	effect := map[string]any{constants.InitiativeRoll: 5}
	f.SetEffects(effect)
	return f
}
