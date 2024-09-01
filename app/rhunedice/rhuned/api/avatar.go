package api

import "fmt"

type IAvatar interface {
	BucketSelectionTurn()
	EndTurn()
	ExecuteButcketTurn()
	GetActions() []IAction
	GetBuckets() IBucketSet
	GetDiceSet() IDiceSet
	GetEquipment() IEquipment
	GetName() string
	GetStats() IStatSet
	GetSelected() IBucketSet
	IsActive() bool
	RollDiceTurn()
	SelectBucketTurn()
	SetActions([]IAction)
	SetActive(bool)
	SetBuckets(IBucketSet)
	SetDiceSet(IDiceSet)
	SetEquipment(IEquipment)
	SetName(string)
	SetStats(IStatSet)
	SetSelected(IBucketSet)
	StartTurn()
	String() string
	UpdateBucketTurn()
	UpdateTurn()
}

type Avatar struct {
	actions   []IAction
	active    bool
	buckets   IBucketSet
	diceset   IDiceSet
	equipment IEquipment
	name      string
	stats     IStatSet
	selected  IBucketSet
}

func NewAvatar(
	name string,
	stats IStatSet,
	diceset IDiceSet,
	buckets IBucketSet,
	equipment IEquipment,
	actions []IAction) *Avatar {

	return &Avatar{
		actions:   actions,
		active:    false,
		buckets:   buckets,
		diceset:   diceset,
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
		bucketName := stat.GetBucketName()
		bucket := a.buckets.GetBucketByName(bucketName)
		bucket.Inc(stat.GetValue())
	}
}

func (a *Avatar) BucketSelectionTurn() {
}

func (a *Avatar) EndTurn() {
}

func (a *Avatar) ExecuteButcketTurn() {
}

func (a *Avatar) GetActions() []IAction {
	return a.actions
}

func (a *Avatar) GetBuckets() IBucketSet {
	return a.buckets
}

func (a *Avatar) GetDiceSet() IDiceSet {
	return a.diceset
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

func (a *Avatar) IsActive() bool {
	return a.active
}

func (a *Avatar) RollDiceTurn() {
}

func (a *Avatar) SelectBucketTurn() {
}

func (a *Avatar) SetActions(actions []IAction) {
	a.actions = actions
}

func (a *Avatar) SetActive(active bool) {
	a.active = active
}

func (a *Avatar) SetBuckets(buckets IBucketSet) {
	a.buckets = buckets
}

func (a *Avatar) SetDiceSet(diceset IDiceSet) {
	a.diceset = diceset
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
	return fmt.Sprintf("%s %s %s %s", a.name, a.stats, a.diceset, a.buckets)
}

func (a *Avatar) UpdateBucketTurn() {
	a.updateBucketsWithStats()
	a.updateBucketsWithEquipment()
}

func (a *Avatar) UpdateTurn() {
}

var _ IAvatar = (*Avatar)(nil)
