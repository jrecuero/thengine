// shields.go contains all shields to be used.
package shields

import "github.com/jrecuero/thengine/app/game/dad/rules"

const (
	ShieldName = "shield/shield/shield"
)

func init() {
	sections := []string{rules.DbSectionGear, rules.DbSectionShield}
	rules.DBase.Add(sections, rules.NewDatabaseEntry(ShieldName, NewShield))
}

// -----------------------------------------------------------------------------
//
// Shield
//
// -----------------------------------------------------------------------------

func NewShield() *rules.Shield {
	return rules.NewShield(ShieldName, "shield", 10, 2, 2)
}
