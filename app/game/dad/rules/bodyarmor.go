// bodyarmor.go contains all pieces of armor that can be equiped in any part of
// the body.
package rules

// -----------------------------------------------------------------------------
//
// HeadGear
//
// -----------------------------------------------------------------------------

type HeadGear struct {
	*Armor
}

func NewHeadGear(name string, cost int, weight int) *HeadGear {
	return &HeadGear{
		Armor: NewArmor(name, cost, weight, "head"),
	}
}

var _ IDamage = (*HeadGear)(nil)
var _ IBattleGear = (*HeadGear)(nil)
var _ IArmor = (*HeadGear)(nil)

// -----------------------------------------------------------------------------
//
// BodyGear
//
// -----------------------------------------------------------------------------

type BodyGear struct {
	*Armor
}

func NewBodyGear(name string, cost int, weight int) *BodyGear {
	return &BodyGear{
		Armor: NewArmor(name, cost, weight, "body"),
	}
}

var _ IDamage = (*BodyGear)(nil)
var _ IBattleGear = (*BodyGear)(nil)
var _ IArmor = (*BodyGear)(nil)

// -----------------------------------------------------------------------------
//
// ArmsGear
//
// -----------------------------------------------------------------------------

type ArmsGear struct {
	*Armor
}

func NewArmsGear(name string, cost int, weight int) *ArmsGear {
	return &ArmsGear{
		Armor: NewArmor(name, cost, weight, "arms"),
	}
}

var _ IDamage = (*ArmsGear)(nil)
var _ IBattleGear = (*ArmsGear)(nil)
var _ IArmor = (*ArmsGear)(nil)

// -----------------------------------------------------------------------------
//
// HandsGear
//
// -----------------------------------------------------------------------------

type HandsGear struct {
	*Armor
}

func NewHandsGear(name string, cost int, weight int) *HandsGear {
	return &HandsGear{
		Armor: NewArmor(name, cost, weight, "hands"),
	}
}

var _ IDamage = (*HandsGear)(nil)
var _ IBattleGear = (*HandsGear)(nil)
var _ IArmor = (*HandsGear)(nil)

// -----------------------------------------------------------------------------
//
// LegsGear
//
// -----------------------------------------------------------------------------

type LegsGear struct {
	*Armor
}

func NewLegsGear(name string, cost int, weight int) *LegsGear {
	return &LegsGear{
		Armor: NewArmor(name, cost, weight, "legs"),
	}
}

var _ IDamage = (*LegsGear)(nil)
var _ IBattleGear = (*LegsGear)(nil)
var _ IArmor = (*LegsGear)(nil)

// -----------------------------------------------------------------------------
//
// FeetGear
//
// -----------------------------------------------------------------------------

type FeetGear struct {
	*Armor
}

func NewFeetGear(name string, cost int, weight int) *FeetGear {
	return &FeetGear{
		Armor: NewArmor(name, cost, weight, "feet"),
	}
}

var _ IDamage = (*FeetGear)(nil)
var _ IBattleGear = (*FeetGear)(nil)
var _ IArmor = (*FeetGear)(nil)
