package api

type IArmor interface {
}

type Armor struct {
	description string
	name        string
	rhune       IRhune
}
