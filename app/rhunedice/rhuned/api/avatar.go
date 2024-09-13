package api

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/tools"
)

type IAvatar interface {
	BucketSelectionTurn()
	EndTurn()
	ExecuteButcketTurn(IBucket, ...any)
	GetActions() []IAction
	GetBuckets() IBucketSet
	GetBucketTotalValueByName(string) int
	GetDiceSet() IDiceSet
	GetEquipment() IEquipment
	GetKnowledge() IKnowledge
	GetName() string
	GetRollDiceBuckets() []IBucket
	GetStats() IStatSet
	GetSelected() []IBucket
	IsActive() bool
	NextSelected() IBucket
	RemoveNextSelected()
	RollDiceTurn()
	//SelectBucketTurn()
	SetActions([]IAction)
	SetActive(bool)
	SetBuckets(IBucketSet)
	SetDiceSet(IDiceSet)
	SetEquipment(IEquipment)
	SetKnowledge(IKnowledge)
	SetName(string)
	SetStats(IStatSet)
	SetSelected([]IBucket)
	StartTurn()
	String() string
	UpdateBucketTurn()
	UpdateTurn()
}

type Avatar struct {
	actions     []IAction
	active      bool
	buckets     IBucketSet
	diceset     IDiceSet
	equipment   IEquipment
	knowledge   IKnowledge
	name        string
	rollfaces   []IFace
	rollbuckets []IBucket
	stats       IStatSet
	selected    []IBucket
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
		knowledge: NewKnowledge(0),
		name:      name,
		rollfaces: nil,
		stats:     stats,
		selected:  nil,
	}
}

func (a *Avatar) resetBuckets() {
	for _, bucket := range a.buckets.GetBuckets() {
		bucket.SetValue(0)
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

func (a *Avatar) updateStatsWithBuckets() {
	for _, bucket := range a.buckets.GetBuckets() {
		bucketName := bucket.GetName()
		if stat := a.stats.GetStatByName(bucketName); stat != nil {
			stat.SetValue(bucket.GetValue())
		}
	}
}

func (a *Avatar) BucketSelectionTurn() {

	//if a.selected != nil {
	//    a.buckets.UpdateBucketsFromBuckets(a.selected)
	//}
	tools.Logger.WithField("module", "avatar").
		WithField("method", "BucketSelectionTurn").
		Tracef("%s bucket selection %s", a.GetName(), a.buckets)
}

func (a *Avatar) EndTurn() {

	tools.Logger.WithField("module", "avatar").
		WithField("method", "EndTurn").
		Tracef("%s end turn", a.GetName())

	a.updateStatsWithBuckets()
}

func (a *Avatar) ExecuteButcketTurn(selBucket IBucket, args ...any) {

	tools.Logger.WithField("module", "avatar").
		WithField("method", "ExecuteBucketTurn").
		Tracef("%s execute bucket %#+v", a.GetName(), args)

	other := args[0].(IAvatar)
	cat := selBucket.GetCat().(EBucketCat)
	buckets := a.buckets.GetBucketsForCat(cat)
	for _, bucket := range buckets {
		totalSt := fmt.Sprintf("%d + %d", bucket.GetValue(), selBucket.GetValue())
		switch cat {
		case AttackBucket:
			attack := bucket.GetValue() + selBucket.GetValue()
			defense := other.GetBucketTotalValueByName(DefenseName)
			damage := attack - defense
			if damage < 0 {
				damage = 0
			}
			tools.Logger.WithField("module", "avatar").
				WithField("method", "ExecuteBucketTurn").
				Tracef("%s execute %s bucket for %s vs %d. damage: %d",
					a.GetName(), cat, totalSt, defense, damage)
			otherHealth := other.GetBuckets().GetBucketByName(HealthName)
			otherHealth.Dec(damage)
		case DefenseBucket:
			tools.Logger.WithField("module", "avatar").
				WithField("method", "ExecuteBucketTurn").
				Tracef("%s execute %s bucket for %s",
					a.GetName(), cat, totalSt)
		case SkillBucket:
			skill := bucket.GetValue() + selBucket.GetValue()
			defense := other.GetBucketTotalValueByName(DefenseName)
			damage := skill - defense
			if damage < 0 {
				damage = 0
			}
			tools.Logger.WithField("module", "avatar").
				WithField("method", "ExecuteBucketTurn").
				Tracef("%s execute %s bucket for %s vs %d. damage: %d",
					a.GetName(), cat, totalSt, defense, damage)
			otherHealth := other.GetBuckets().GetBucketByName(HealthName)
			otherHealth.Dec(damage)
		case StepBucket:
			steps := bucket.GetValue() + selBucket.GetValue()
			tools.Logger.WithField("module", "avatar").
				WithField("method", "ExecuteBucketTurn").
				Tracef("%s execute %s bucket for %s move %d steps",
					a.GetName(), cat, totalSt, steps)
		case HealthBucket:
			incHealth := bucket.GetValue() + selBucket.GetValue()
			health := a.GetBuckets().GetBucketByName(HealthName)
			health.Inc(incHealth)
			tools.Logger.WithField("module", "avatar").
				WithField("method", "ExecuteBucketTurn").
				Tracef("%s execute %s bucket for %s add %d health",
					a.GetName(), cat, totalSt, incHealth)
		case StaminaBucket:
			incStamina := bucket.GetValue() + selBucket.GetValue()
			stamina := a.GetBuckets().GetBucketByName(StaminaName)
			stamina.Inc(incStamina)
			tools.Logger.WithField("module", "avatar").
				WithField("method", "ExecuteBucketTurn").
				Tracef("%s execute %s bucket for %s add %d stamina",
					a.GetName(), cat, totalSt, incStamina)
		case ExtraBucket:
			rhune := selBucket.GetRhune()
			rhune.Execute(a)
			tools.Logger.WithField("module", "avatar").
				WithField("method", "ExecuteBucketTurn").
				Tracef("%s execute %s bucket with %s",
					a.GetName(), cat, bucket.GetRhune())
		}
	}
}

func (a *Avatar) GetActions() []IAction {
	return a.actions
}

func (a *Avatar) GetBuckets() IBucketSet {
	return a.buckets
}

func (a *Avatar) GetBucketTotalValueByName(name string) int {
	bucket := a.buckets.GetBucketByName(name)
	var selectedBucket IBucket
	for _, buck := range a.selected {
		if buck.GetName() == name {
			selectedBucket = buck
			break
		}
	}
	// resultA is the value for the avatar-bucket.
	var resultA int = bucket.GetValue()
	// resultB is the value for the roll-dice-selected-bucket.
	var resultB int
	if selectedBucket != nil {
		resultB = selectedBucket.GetValue()
	}
	tools.Logger.WithField("module", "avatar").
		WithField("method", "GetBucketTotalValueByName").
		Tracef("%s total %s is %d + %d",
			a.GetName(), name, resultA, resultB)
	return resultA + resultB
}

func (a *Avatar) GetDiceSet() IDiceSet {
	return a.diceset
}

func (a *Avatar) GetEquipment() IEquipment {
	return a.equipment
}

func (a *Avatar) GetKnowledge() IKnowledge {
	return a.knowledge
}

func (a *Avatar) GetName() string {
	return a.name
}

func (a *Avatar) GetRollDiceBuckets() []IBucket {
	return a.rollbuckets
}

func (a *Avatar) GetStats() IStatSet {
	return a.stats
}

func (a *Avatar) GetSelected() []IBucket {
	return a.selected
}

func (a *Avatar) IsActive() bool {
	return a.active
}

// NextSelected method returns any remaining bucket still left in the selected
// list.
func (a *Avatar) NextSelected() IBucket {
	if len(a.selected) != 0 {
		return a.selected[0]
	}
	return nil
}

// RemoveNextSelected method removes the first bucket in the list of selected
// buckets.
func (a *Avatar) RemoveNextSelected() {
	if len(a.selected) > 0 {
		a.selected = a.selected[1:]
	}
}

func (a *Avatar) RollDiceTurn() {
	if a.diceset != nil {
		a.rollfaces = a.diceset.Roll()

		a.rollbuckets = FacesToBuckets(a.rollfaces)
	}
	tools.Logger.WithField("module", "avatar").
		WithField("method", "RollDice").
		Tracef("%s roll dice %s", a.GetName(), a.rollfaces)
}

//func (a *Avatar) SelectBucketTurn() {
//    tools.Logger.WithField("module", "avatar").
//        WithField("method", "SelectBucketTurn").
//        Tracef("%s select bucket from %s", a.GetName(), a.rollbuckets)
//}

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

func (a *Avatar) SetKnowledge(knowledge IKnowledge) {
	a.knowledge = knowledge
}

func (a *Avatar) SetName(name string) {
	a.name = name
}

func (a *Avatar) SetStats(stats IStatSet) {
	a.stats = stats
}

func (a *Avatar) SetSelected(buckets []IBucket) {
	a.selected = buckets
}

func (a *Avatar) StartTurn() {

	tools.Logger.WithField("module", "avatar").
		WithField("method", "StartTurn").
		Tracef("%s start turn", a.GetName())

	a.rollfaces = nil
	a.selected = nil

	a.resetBuckets()
}

func (a *Avatar) String() string {
	return fmt.Sprintf("%s %s %s %s", a.name, a.stats, a.diceset, a.buckets)
}

func (a *Avatar) UpdateBucketTurn() {

	tools.Logger.WithField("module", "avatar").
		WithField("method", "UpdateBucketTurn").
		Tracef("%s update bucket", a.GetName())

	a.updateBucketsWithStats()
	a.updateBucketsWithEquipment()
}

func (a *Avatar) UpdateTurn() {

	tools.Logger.WithField("module", "avatar").
		WithField("method", "UpdateTurn").
		Tracef("%s update turn", a.GetName())
}

var _ IAvatar = (*Avatar)(nil)
