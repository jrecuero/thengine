package api

import "fmt"

type IAvatar interface {
	GetActions() []IAction
	GetBuckets() IBucketSet
	GetEquipment() IEquipment
	GetName() string
	GetStats() IStatSet
	GetSelected() IBucketSet
	SetActions([]IAction)
	SetBuckets(IBucketSet)
	SetEquipment(IEquipment)
	SetName(string)
	SetStats(IStatSet)
	SetSelected(IBucketSet)
	StartTurn()
	String() string
}

type Avatar struct {
	actions   []IAction
	buckets   IBucketSet
	equipment IEquipment
	name      string
	stats     IStatSet
	selected  IBucketSet
}

func NewAvatar(name string, stats IStatSet, buckets IBucketSet,
	equipment IEquipment, actions []IAction) *Avatar {
	return &Avatar{
		actions:   actions,
		buckets:   buckets,
		equipment: equipment,
		name:      name,
		stats:     stats,
		selected:  nil,
	}
}

func (a *Avatar) updateBucketsWithEquipment() {
	if armor := a.GetEquipment().GetArmor(); armor != nil {
		bucket := a.GetBuckets().GetBucketByName(armor.GetBucketName())
		bucket.Inc(armor.GetValue())
	}
}

func (a *Avatar) updateBucketsWithStats() {
	for _, stat := range a.stats.GetStats() {
		stat.GetBucket().Inc(stat.GetValue())
	}
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

func (a *Avatar) GetStats() IStatSet {
	return a.stats
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

func (a *Avatar) SetStats(stats IStatSet) {
	a.stats = stats
}

func (a *Avatar) SetSelected(buckets IBucketSet) {
	a.selected = buckets
}

func (a *Avatar) StartTurn() {
	a.updateBucketsWithStats()
	a.updateBucketsWithEquipment()
}

func (a *Avatar) String() string {
	return fmt.Sprintf("%s %s %s", a.name, a.stats, a.buckets)
}

var _ IAvatar = (*Avatar)(nil)
