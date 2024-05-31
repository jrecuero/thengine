// gear.go contains all data and methods related with the unit gear that can be
// equiped. It relates to weapons, shields and any can or armament to be
// equiped by the unit
package rules

import (
	"fmt"

	"github.com/jrecuero/thengine/app/game/dad/battlelog"
)

// -----------------------------------------------------------------------------
//
// IGear
//
// -----------------------------------------------------------------------------

// IGear interfaces defines all methods any Geara structure should implement.
type IGear interface {
	AC() int
	GetAccessories() []any
	GetArms() IArmor
	GetBody() IArmor
	GetFeet() IArmor
	GetHands() IArmor
	GetHead() IArmor
	GetLegs() IArmor
	GetMainHand() IHandheld
	GetOffHand() IHandheld
	RollDamage() int
	SetAccessories(...any)
	SetArms(IArmor)
	SetBody(IArmor)
	SetFeet(IArmor)
	SetHands(IArmor)
	SetHead(IArmor)
	SetLegs(IArmor)
	SetMainHand(IHandheld)
	SetOffHand(IHandheld)
}

// -----------------------------------------------------------------------------
//
// Gear
//
// -----------------------------------------------------------------------------

// Gear structure defines all parts of the unit that can equip any gear.
type Gear struct {
	mainhand    IHandheld
	offhand     IHandheld
	head        IArmor
	body        IArmor
	arms        IArmor
	hands       IArmor
	legs        IArmor
	feet        IArmor
	accessories []any
}

func NewGear() *Gear {
	return &Gear{
		mainhand:    nil,
		offhand:     nil,
		head:        nil,
		body:        nil,
		arms:        nil,
		hands:       nil,
		legs:        nil,
		feet:        nil,
		accessories: nil,
	}
}

// -----------------------------------------------------------------------------
// Gear public methods
// -----------------------------------------------------------------------------

func (g *Gear) AC() int {
	result := 0
	if g.GetArms() != nil {
		result += g.GetArms().GetAC()
	}
	if g.GetBody() != nil {
		result += g.GetBody().GetAC()
	}
	if g.GetFeet() != nil {
		result += g.GetFeet().GetAC()
	}
	if g.GetHands() != nil {
		result += g.GetHands().GetAC()
	}
	if g.GetHead() != nil {
		result += g.GetHead().GetAC()
	}
	if g.GetLegs() != nil {
		result += g.GetLegs().GetAC()
	}
	return result
}

func (g *Gear) GetAccessories() []any {
	return g.accessories
}

func (g *Gear) GetArms() IArmor {
	return g.arms
}

func (g *Gear) GetBody() IArmor {
	return g.body
}

func (g *Gear) GetFeet() IArmor {
	return g.feet
}

func (g *Gear) GetHands() IArmor {
	return g.hands
}

func (g *Gear) GetHead() IArmor {
	return g.head
}

func (g *Gear) GetLegs() IArmor {
	return g.legs
}

func (g *Gear) GetMainHand() IHandheld {
	return g.mainhand
}

func (g *Gear) GetOffHand() IHandheld {
	return g.offhand
}

func (g *Gear) RollDamage() int {
	mainHandDamage := 0
	offHandDamage := 0
	if g.mainhand != nil {
		mainHandDamage = g.mainhand.RollDamage()
		battlelog.BLog.Push(fmt.Sprintf("main-hand damage: %d", mainHandDamage))
	}
	if g.offhand != nil {
		offHandDamage = g.offhand.RollDamage()
		battlelog.BLog.Push(fmt.Sprintf("off-hand damage: %d", offHandDamage))
	}
	return mainHandDamage + offHandDamage
}

func (g *Gear) SetAccessories(gears ...any) {
	for _, gear := range gears {
		g.accessories = append(g.accessories, gear)
	}
}

func (g *Gear) SetArms(gear IArmor) {
	g.arms = gear
}

func (g *Gear) SetBody(gear IArmor) {
	g.body = gear
}

func (g *Gear) SetFeet(gear IArmor) {
	g.feet = gear
}

func (g *Gear) SetHands(gear IArmor) {
	g.head = gear
}

func (g *Gear) SetHead(gear IArmor) {
	g.hands = gear
}

func (g *Gear) SetLegs(gear IArmor) {
	g.legs = gear
}

func (g *Gear) SetMainHand(handheld IHandheld) {
	g.mainhand = handheld
}

func (g *Gear) SetOffHand(handheld IHandheld) {
	g.offhand = handheld
}

var _ IGear = (*Gear)(nil)
