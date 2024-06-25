// interfaces.go contains all common interfaces to dad package.
package rules

// -----------------------------------------------------------------------------
//
// IRollBonus
//
// -----------------------------------------------------------------------------

// IRollBonus interface defines methods required to be implemented for any
// structure that can apply any bonus to any roll.
type IRollBonus interface {
	GetRollBonusForAction(string) any
}
