package spells

import "github.com/jrecuero/thengine/app/game/dad/rules"

const (
	GuidingBoltName = "spell/evocation/guiding-bold"
)

func init() {
	sections := []string{rules.DbSectionSpells}
	rules.DBase.Add(sections, rules.NewDatabaseEntry(GuidingBoltName, NewGuidingBolt))
}

// -----------------------------------------------------------------------------
//
// GuidingBolt
//
// -----------------------------------------------------------------------------

// NewGuidingBolt creates an Spell instance for a Guiding Bold.
// A flash of light streaks toward a creature of your choice within range. Make
// a ranged spell attack against the target. On a hit, the target takes 4d6
// radiant damage, and the next attack roll made against this target before the
// end of your next turn has advantage, thanks to the mystical dim light
// glittering on the target until then
func NewGuidingBolt() *rules.Spell {
	guidingBolt := rules.NewSpell(GuidingBoltName, "guiding bolt", rules.EvocationMagic, 1, rules.NewDamage(rules.DiceThrow4d6, rules.Magical))
	guidingBolt.SetDescription(`A flash of light streaks toward a creature of
 your choice within range. Make a ranged spell attack against the target. 
On a hit, the target takes 4d6 radiant damage, and the next attack roll made 
 against this target before the end of your next turn has advantage, thanks 
 to the mystical dim light glittering on the target until then.`)
	guidingBolt.SetComponents([]rules.MagicComponent{rules.VerbalComponent, rules.SomaticComponent})
	guidingBolt.SetCastingTime(1)                                                  // 1 action
	guidingBolt.SetRange(120)                                                      // TODO: 120 feet
	guidingBolt.SetDuration(1)                                                     // 1 round
	guidingBolt.SetHigherLevel(rules.NewDamage(rules.DiceThrow1d6, rules.Magical)) // 1d6 per higher level
	return guidingBolt
}
