// shield.go contains all data and methods common to any shield to be used.
package rules

// -----------------------------------------------------------------------------
//
// IShield
//
// -----------------------------------------------------------------------------

// IShield interface defines all methods any Shield structure should be
// implementing.
type IShield interface {
	IHandheld
}

// -----------------------------------------------------------------------------
//
// Shield
//
// -----------------------------------------------------------------------------

// Shield structure defines all attributes and methods for the basic shield.
type Shield struct {
	*Handheld
}

func NewShield(name string, uname string, cost int, weight int, ac int) *Shield {
	shield := &Shield{
		Handheld: NewHandheld(name, uname, cost, weight, NewShieldHandheldType()),
	}
	shield.SetAC(ac)
	return shield
}
