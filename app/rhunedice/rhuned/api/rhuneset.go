package api

import "fmt"

type IRhuneSet interface {
	AddRhune(IRhune) error
	GetName() string
	GetRhuneByName(string) IRhune
	GetRhunes() []IRhune
	GetRhunesForCat(ERhuneCat) []IRhune
	RemoveRhune(IRhune)
	SetName(string)
	SetRhunes([]IRhune)
	String() string
}

type RhuneSet struct {
	rhunes []IRhune
	name   string
}

func (r *RhuneSet) getRhuneAndIndex(name string) (IRhune, int) {
	//for index, rhune := range r.rhunes {
	//    if rhune.GetName() == name {
	//        return rhune, index
	//    }
	//}
	//return nil, -1
	if rhune, index, found := FindByNameWithIndex(r.rhunes, name); found {
		return rhune, index
	}
	return nil, -1
}

func (r *RhuneSet) AddRhune(rhune IRhune) error {
	return nil
}

func (r *RhuneSet) GetRhunesForCat(cat ERhuneCat) []IRhune {
	return FindByCat(r.rhunes, cat)
}

func (r *RhuneSet) String() string {
	return fmt.Sprintf("%s %s", r.name, r.rhunes)
}

//var _ IRhuneSet = (*RhuneSet)(nil)
