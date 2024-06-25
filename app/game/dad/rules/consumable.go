// consumable.go package contains all data and logic required to implement any
// consumable item.
package rules

import "github.com/jrecuero/thengine/pkg/tools"

// -----------------------------------------------------------------------------
//
// IConsumable
//
// -----------------------------------------------------------------------------

type IConsumable interface {
	IRollBonus
	GetCost() int
	GetDescription() string
	GetEffects() map[string]any
	GetName() string
	GetWeight() int
	SeCost(int)
	SetDescription(string)
	SetEffects(map[string]any)
	SetName(string)
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
	weight      int
}

// -----------------------------------------------------------------------------
// Consumable public methods
// -----------------------------------------------------------------------------

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

func (c *Consumable) SetWeight(weight int) {
	c.weight = weight
}

var _ IConsumable = (*Consumable)(nil)
