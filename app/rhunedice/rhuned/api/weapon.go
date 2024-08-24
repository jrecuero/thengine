package api

type IWeapon interface {
	IEquipmentPiece
}

type Weapon struct {
	*EquipmentPiece
}

func NewWeapon(name string, description string, bucketname string, value int) *Weapon {
	return &Weapon{
		EquipmentPiece: NewEquipmentPiece(name, description, bucketname, value),
	}
}

var _ IEquipmentPiece = (*Weapon)(nil)
var _ IWeapon = (*Weapon)(nil)
