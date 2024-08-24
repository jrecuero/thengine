package api

type IAccessory interface {
	IEquipmentPiece
}

type Accessory struct {
	*EquipmentPiece
}

func NewAccessory(name string, description string, bucketname string, value int) *Accessory {
	return &Accessory{
		EquipmentPiece: NewEquipmentPiece(name, description, bucketname, value),
	}
}

var _ IEquipmentPiece = (*Accessory)(nil)
var _ IAccessory = (*Accessory)(nil)
