package api

type IAccessory interface {
}

type Accessory struct {
	description string
	name        string
	rhune       IRhune
}
