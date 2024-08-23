package api

import "fmt"

type IAvatar interface {
	GetActions() []IAction
	GetBuckets() IBucketSet
	GetEquipment() IEquipment
	GetName() string
	GetStats() []IStat
	GetSelected() IBucketSet
	SetActions([]IAction)
	SetBuckets(IBucketSet)
	SetEquipment(IEquipment)
	SetName(string)
	SetStats([]IStat)
	SetSelected(IBucketSet)
	String() string
}

type Avatar struct {
	actions   []IAction
	buckets   IBucketSet
	equipment IEquipment
	name      string
	rhunes    []IStat
	selected  IBucketSet
}

func (a *Avatar) GetActions() []IAction {
	return a.actions
}

func (a *Avatar) GetBuckets() IBucketSet {
	return a.buckets
}

func (a *Avatar) GetEquipment() IEquipment {
	return a.equipment
}

func (a *Avatar) GetName() string {
	return a.name
}

func (a *Avatar) GetStats() []IStat {
	return a.rhunes
}

func (a *Avatar) GetSelected() IBucketSet {
	return a.selected
}

func (a *Avatar) SetActions(actions []IAction) {
	a.actions = actions
}

func (a *Avatar) SetBuckets(buckets IBucketSet) {
	a.buckets = buckets
}

func (a *Avatar) SetEquipment(equipment IEquipment) {
	a.equipment = equipment
}

func (a *Avatar) SetName(name string) {
	a.name = name
}

func (a *Avatar) SetStats(rhunes []IStat) {
	a.rhunes = rhunes
}

func (a *Avatar) SetSelected(buckets IBucketSet) {
	a.selected = buckets
}

func (a *Avatar) String() string {
	return fmt.Sprintf("%s", a.name)
}

var _ IAvatar = (*Avatar)(nil)
