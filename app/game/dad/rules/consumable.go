// consumable.go package contains all data and logic required to implement any
// consumable item.
package rules

import (
	"github.com/jrecuero/thengine/app/game/dad/constants"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// IConsumable
//
// -----------------------------------------------------------------------------

type IConsumable interface {
	IRollBonus
	Consume(IUnit) error
	GetCost() int
	GetDescription() string
	GetEffects() map[string]any
	GetName() string
	GetUName() string
	GetWeight() int
	SeCost(int)
	SetDescription(string)
	SetEffects(map[string]any)
	SetName(string)
	SetUName(string)
	SetWeight(int)
}

// -----------------------------------------------------------------------------
//
// Consumable
//
// -----------------------------------------------------------------------------

type Consumable struct {
	cost        int
	description string
	effects     map[string]any
	name        string
	uname       string
	weight      int
}

func NewConsumable(name string, uname string, cost int, weight int) *Consumable {
	return &Consumable{
		cost:        cost,
		description: name,
		effects:     make(map[string]any),
		name:        name,
		uname:       uname,
		weight:      weight,
	}
}

// -----------------------------------------------------------------------------
// Consumable public methods
// -----------------------------------------------------------------------------

// Consume method implements the fact to use/consume the consumable item.
func (c *Consumable) Consume(unit IUnit) error {
	for key, effect := range c.effects {
		if key == constants.ConsumableRoll {
			err := (effect.(func(IUnit) error))(unit)
			return err
		}
	}
	return nil
}

func (c *Consumable) GetCost() int {
	return c.cost
}

func (c *Consumable) GetDescription() string {
	return c.description
}

func (c *Consumable) GetEffects() map[string]any {
	return c.effects
}

func (c *Consumable) GetName() string {
	return c.name
}

func (c *Consumable) GetRollBonusForAction(action string) any {
	for k, v := range c.GetEffects() {
		if k == action {
			tools.Logger.WithField("module", "consumable").
				WithField("method", "GetRollBonusForAction").
				Debugf("consumable %s bonus %v for %s", c.GetName(), v.(int), action)
			return v
		}
	}
	return nil
}

func (c *Consumable) GetUName() string {
	return c.uname
}

func (c *Consumable) GetWeight() int {
	return c.weight
}

func (c *Consumable) SeCost(cost int) {
	c.cost = cost
}

func (c *Consumable) SetDescription(description string) {
	c.description = description
}

func (c *Consumable) SetEffects(effects map[string]any) {
	c.effects = effects
}

func (c *Consumable) SetName(name string) {
	c.name = name
}

func (c *Consumable) SetUName(name string) {
	c.uname = name
}

func (c *Consumable) SetWeight(weight int) {
	c.weight = weight
}

var _ IConsumable = (*Consumable)(nil)
