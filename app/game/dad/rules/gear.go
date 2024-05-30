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
	GetArms() any
	GetBody() any
	GetFeet() any
	GetHead() any
	GetLegs() any
	GetMainHand() IHandheld
	GetOffHand() IHandheld
	RollDamage() int
	SetAccessories(...any)
	SetArms(any)
	SetBody(any)
	SetFeet(any)
	SetHead(any)
	SetLegs(any)
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
	head        any
	body        any
	arms        any
	legs        any
	feet        any
	accessories []any
}

func NewGear() *Gear {
	return &Gear{
		mainhand:    nil,
		offhand:     nil,
		head:        nil,
		body:        nil,
		arms:        nil,
		legs:        nil,
		feet:        nil,
		accessories: nil,
	}
}

// -----------------------------------------------------------------------------
// Gear public methods
// -----------------------------------------------------------------------------

func (g *Gear) AC() int {
	return 0
}

func (g *Gear) GetAccessories() []any {
	return g.accessories
}

func (g *Gear) GetArms() any {
	return g.arms
}

func (g *Gear) GetBody() any {
	return g.body
}

func (g *Gear) GetFeet() any {
	return g.feet
}

func (g *Gear) GetHead() any {
	return g.head
}

func (g *Gear) GetLegs() any {
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

func (g *Gear) SetArms(gear any) {
	g.arms = gear
}

func (g *Gear) SetBody(gear any) {
	g.body = gear
}

func (g *Gear) SetFeet(gear any) {
	g.feet = gear
}

func (g *Gear) SetHead(gear any) {
	g.head = gear
}

func (g *Gear) SetLegs(gear any) {
	g.legs = gear
}

func (g *Gear) SetMainHand(handheld IHandheld) {
	g.mainhand = handheld
}

func (g *Gear) SetOffHand(handheld IHandheld) {
	g.offhand = handheld
}

var _ IGear = (*Gear)(nil)
