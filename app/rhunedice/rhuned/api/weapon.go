package api

type IWeapon interface {
}

type Weapon struct {
	description string
	name        string
	rhune       IRhune
}
