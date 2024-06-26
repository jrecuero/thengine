// inventory.go package contains all attributes and methods to implement an
// inventory used by any Played Character.
package rules

// -----------------------------------------------------------------------------
//
// IInventory
//
// -----------------------------------------------------------------------------

type IInventory interface {
	AddAccessories(...any) error
	AddArmors(...IBattleGear) error
	AddConsumables(...IConsumable) error
	AddCurrencies(...any) error
	AddMagicItems(...any) error
	AddMaterials(...any) error
	AddMiscellaneous(...any) error
	AddTools(...any) error
	AddWeapons(...IBattleGear) error
	GetAccessoryByName(string) any
	GetAccessories() []any
	GetArmorByName(string) IBattleGear
	GetArmors() []IBattleGear
	GetConsumableByName(string) IConsumable
	GetConsumables() []IConsumable
	GetCurrencyByName(string) any
	GetCurrencies() []any
	GetMagicItemByname(string) any
	GetMagicItems() []any
	GetMaterialByName(string) any
	GetMaterials() []any
	GetMiscellaneousByName(string) any
	GetMiscellaneous() []any
	GetName() string
	GetToolByName(string) any
	GetTools() []any
	GetWeaponByName(string) IBattleGear
	RemoveAccessory(any) error
	RemoveArmor(IBattleGear) error
	RemoveConsumable(IConsumable) error
	RemoveCurrency(any) error
	RemoveMagicItem(any) error
	RemoveMaterial(any) error
	RemoveMiscellaneous(any) error
	RemoveTool(any) error
	RemoveWeapon(IBattleGear) error
	GetWeapons() []IBattleGear
	SetAccessories([]any)
	SetArmors([]IBattleGear)
	SetConsumables([]IConsumable)
	SetCurrencies([]any)
	SetMagicItems([]any)
	SetMaterials([]any)
	SetMiscellaneous([]any)
	SetName(string)
	SetTools([]any)
	SetWeapons([]IBattleGear)
}

// -----------------------------------------------------------------------------
//
// Inventory
//
// -----------------------------------------------------------------------------

type Inventory struct {
	accessories   []any
	armors        []IBattleGear
	consumables   []IConsumable
	currencies    []any
	magicItems    []any
	materials     []any
	miscellaneous []any
	name          string
	tools         []any
	weapons       []IBattleGear
}

func NewInventory(name string) *Inventory {
	return &Inventory{
		accessories:   nil,
		armors:        nil,
		consumables:   nil,
		currencies:    nil,
		magicItems:    nil,
		materials:     nil,
		miscellaneous: nil,
		name:          name,
		tools:         nil,
		weapons:       nil,
	}
}

// -----------------------------------------------------------------------------
// Inventory private methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// Inventory public methods
// -----------------------------------------------------------------------------

func (i *Inventory) AddAccessories(accessories ...any) error {
	i.accessories = append(i.accessories, accessories...)
	return nil
}

func (i *Inventory) AddArmors(armors ...IBattleGear) error {
	i.armors = append(i.armors, armors...)
	return nil
}

func (i *Inventory) AddConsumables(consumables ...IConsumable) error {
	i.consumables = append(i.consumables, consumables...)
	return nil

}

func (i *Inventory) AddCurrencies(currencies ...any) error {
	i.currencies = append(i.currencies, currencies...)
	return nil

}

func (i *Inventory) AddMagicItems(magicItems ...any) error {
	i.magicItems = append(i.magicItems, magicItems...)
	return nil

}

func (i *Inventory) AddMaterials(materials ...any) error {
	i.materials = append(i.materials, materials...)
	return nil

}

func (i *Inventory) AddMiscellaneous(miscellaneous ...any) error {
	i.miscellaneous = append(i.miscellaneous, miscellaneous...)
	return nil

}

func (i *Inventory) AddTools(tools ...any) error {
	i.tools = append(i.tools, tools...)
	return nil

}

func (i *Inventory) AddWeapons(weapons ...IBattleGear) error {
	i.weapons = append(i.weapons, weapons...)
	return nil

}

func (i *Inventory) GetAccessoryByName(name string) any {
	return nil
}

func (i *Inventory) GetAccessories() []any {
	return nil
}

func (i *Inventory) GetArmorByName(name string) IBattleGear {
	return nil
}

func (i *Inventory) GetArmors() []IBattleGear {
	return i.armors
}

func (i *Inventory) GetConsumableByName(string) IConsumable {
	return nil
}

func (i *Inventory) GetConsumables() []IConsumable {
	return i.consumables
}

func (i *Inventory) GetCurrencyByName(string) any {
	return nil
}

func (i *Inventory) GetCurrencies() []any {
	return i.currencies
}

func (i *Inventory) GetMagicItemByname(string) any {
	return nil
}

func (i *Inventory) GetMagicItems() []any {
	return i.magicItems
}

func (i *Inventory) GetMaterialByName(string) any {
	return nil
}

func (i *Inventory) GetMaterials() []any {
	return i.materials
}

func (i *Inventory) GetMiscellaneous() []any {
	return i.miscellaneous
}

func (i *Inventory) GetMiscellaneousByName(string) any {
	return nil
}

func (i *Inventory) GetName() string {
	return i.name
}

func (i *Inventory) GetToolByName(string) any {
	return nil
}

func (i *Inventory) GetTools() []any {
	return i.tools
}

func (i *Inventory) GetWeaponByName(string) IBattleGear {
	return nil
}

func (i *Inventory) GetWeapons() []IBattleGear {
	return i.weapons
}

func (i *Inventory) RemoveAccessory(accessory any) error {
	for index, traversed := range i.accessories {
		if traversed == accessory {
			i.accessories = append(i.accessories[:index], i.accessories[index+1:]...)
		}
	}
	return nil
}

func (i *Inventory) RemoveArmor(armor IBattleGear) error {
	for index, traversed := range i.armors {
		if traversed == armor {
			i.armors = append(i.armors[:index], i.armors[index+1:]...)
		}
	}
	return nil
}

func (i *Inventory) RemoveConsumable(consumable IConsumable) error {
	for index, traversed := range i.consumables {
		if traversed == consumable {
			i.consumables = append(i.consumables[:index], i.consumables[index+1:]...)
		}
	}
	return nil

}

func (i *Inventory) RemoveCurrency(currency any) error {
	for index, traversed := range i.currencies {
		if traversed == currency {
			i.currencies = append(i.currencies[:index], i.currencies[index+1:]...)
		}
	}
	return nil

}

func (i *Inventory) RemoveMagicItem(magicItem any) error {
	for index, traversed := range i.magicItems {
		if traversed == magicItem {
			i.magicItems = append(i.magicItems[:index], i.magicItems[index+1:]...)
		}
	}
	return nil

}

func (i *Inventory) RemoveMaterial(material any) error {
	for index, traversed := range i.materials {
		if traversed == material {
			i.materials = append(i.materials[:index], i.materials[index+1:]...)
		}
	}
	return nil

}

func (i *Inventory) RemoveMiscellaneous(miscellaneous any) error {
	for index, traversed := range i.miscellaneous {
		if traversed == miscellaneous {
			i.miscellaneous = append(i.miscellaneous[:index], i.miscellaneous[index+1:]...)
		}
	}
	return nil

}

func (i *Inventory) RemoveTool(tool any) error {
	for index, traversed := range i.tools {
		if traversed == tool {
			i.tools = append(i.tools[:index], i.tools[index+1:]...)
		}
	}
	return nil

}

func (i *Inventory) RemoveWeapon(weapon IBattleGear) error {
	for index, traversed := range i.weapons {
		if traversed == weapon {
			i.weapons = append(i.weapons[:index], i.weapons[index+1:]...)
		}
	}
	return nil
}

func (i Inventory) SetAccessories(accessories []any) {
	i.accessories = accessories
}

func (i *Inventory) SetArmors(armors []IBattleGear) {
	i.armors = armors
}

func (i *Inventory) SetConsumables(consumables []IConsumable) {
	i.consumables = consumables
}

func (i *Inventory) SetCurrencies(currencies []any) {
	i.currencies = currencies
}

func (i *Inventory) SetMagicItems(magicItems []any) {
	i.magicItems = magicItems
}

func (i *Inventory) SetMaterials(materials []any) {
	i.materials = materials
}

func (i *Inventory) SetMiscellaneous(miscellaneous []any) {
	i.miscellaneous = miscellaneous
}

func (i *Inventory) SetName(name string) {
	i.name = name
}

func (i *Inventory) SetTools(tools []any) {
	i.tools = tools
}

func (i *Inventory) SetWeapons(weapons []IBattleGear) {
	i.weapons = weapons
}

var _ IInventory = (*Inventory)(nil)
