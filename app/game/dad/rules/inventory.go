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
	RemoveCurrencie(any) error
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

// -----------------------------------------------------------------------------
// Inventory private methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// Inventory public methods
// -----------------------------------------------------------------------------

func (i *Inventory) AddAccessories(...any) error {
	return nil
}

func (i *Inventory) AddArmors(...IBattleGear) error {
	return nil
}

func (i *Inventory) AddConsumables(...IConsumable) error {
	return nil

}

func (i *Inventory) AddCurrencies(...any) error {
	return nil

}

func (i *Inventory) AddMagicItems(...any) error {
	return nil

}

func (i *Inventory) AddMaterials(...any) error {
	return nil

}

func (i *Inventory) AddMiscellaneous(...any) error {
	return nil

}

func (i *Inventory) AddTools(...any) error {
	return nil

}

func (i *Inventory) AddWeapons(...IBattleGear) error {
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

func (i *Inventory) RemoveAccessory(any) error {
	return nil
}

func (i *Inventory) RemoveArmor(IBattleGear) error {
	return nil
}

func (i *Inventory) RemoveConsumable(IConsumable) error {
	return nil

}

func (i *Inventory) RemoveCurrencie(any) error {
	return nil

}

func (i *Inventory) RemoveMagicItem(any) error {
	return nil

}

func (i *Inventory) RemoveMaterial(any) error {
	return nil

}

func (i *Inventory) RemoveMiscellaneous(any) error {
	return nil

}

func (i *Inventory) RemoveTool(any) error {
	return nil

}

func (i *Inventory) RemoveWeapon(IBattleGear) error {
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
