package api

type IArmor interface {
	IEquipmentPiece
}

type Armor struct {
	*EquipmentPiece
}

func NewArmor(name string, description string, bucketname string, value int) *Armor {
	return &Armor{
		EquipmentPiece: NewEquipmentPiece(name, description, bucketname, value),
	}
}

var _ IEquipmentPiece = (*Armor)(nil)
var _ IArmor = (*Armor)(nil)
