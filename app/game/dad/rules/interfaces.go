// interfaces.go contains all common interfaces to dad package.
package rules

// -----------------------------------------------------------------------------
//
// IDieRollBonus
//
// -----------------------------------------------------------------------------

// IDieRollBonus interfaces defines methods required to be implemented for any
// structure that can apply a bonus to a die-roll.
type IDieRollBonus interface {
	DieRollBonus(string) int
}
