package consumables

import (
	"github.com/jrecuero/thengine/app/game/dad/constants"
	"github.com/jrecuero/thengine/app/game/dad/rules"
	"github.com/jrecuero/thengine/pkg/tools"
)

const (
	PotionName = "inventory/consumable/potion"
)

func init() {
	sections := []string{rules.DbSectionInventory, rules.DbSectionConsumables}
	rules.DBase.Add(sections, rules.NewDatabaseEntry(PotionName, NewPotion))
}

// -----------------------------------------------------------------------------
//
// Potion
//
// -----------------------------------------------------------------------------

func NewPotion(live int) *rules.Consumable {
	p := rules.NewConsumable(PotionName, "potion", 5, 1)
	effects := map[string]any{
		constants.ConsumableRoll: func(unit rules.IUnit) error {
			pcLive := unit.GetHitPoints().Inc(live)
			tools.Logger.WithField("module", "potion").
				WithField("method", "Consume/Potion").
				Debugf("Consume potion for %d -> %d", live, pcLive)
			return nil
		},
	}
	p.SetEffects(effects)
	return p
}
