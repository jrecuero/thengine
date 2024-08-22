package api

type IEquipment interface {
}

type Equipment struct {
	accessory IAccessory
	armor     IArmor
	weapon    IWeapon
}
