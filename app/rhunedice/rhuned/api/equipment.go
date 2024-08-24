package api

import "fmt"

type IEquipment interface {
	GetAccessory() IAccessory
	GetArmor() IArmor
	GetWeapon() IWeapon
	SetAccessory(IAccessory)
	SetArmor(IArmor)
	SetWeapon(IWeapon)
	String() string
}

type Equipment struct {
	accessory IAccessory
	armor     IArmor
	weapon    IWeapon
}

func NewEquipment(weapon IWeapon, armor IArmor, accessory IAccessory) *Equipment {
	return &Equipment{
		accessory: accessory,
		armor:     armor,
		weapon:    weapon,
	}
}

func (e *Equipment) GetAccessory() IAccessory {
	return e.accessory
}

func (e *Equipment) GetArmor() IArmor {
	return e.armor
}

func (e *Equipment) GetWeapon() IWeapon {
	return e.weapon
}

func (e *Equipment) SetAccessory(accessory IAccessory) {
	e.accessory = accessory
}

func (e *Equipment) SetArmor(armor IArmor) {
	e.armor = armor
}

func (e *Equipment) SetWeapon(weapon IWeapon) {
	e.weapon = weapon
}

func (e *Equipment) String() string {
	return fmt.Sprintf("%s %s %s",
		e.weapon.GetName(), e.armor.GetName(), e.accessory.GetName())
}

var _ IEquipment = (*Equipment)(nil)
