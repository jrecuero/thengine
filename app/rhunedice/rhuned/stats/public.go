package stats

import (
	"fmt"

	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
)

func DefaultStats(givenStats map[string]int, bucketset api.IBucketSet) []api.IStat {
	defaultStats := []api.IStat{}
	for statName, statValue := range givenStats {
		var newStat api.IStat = nil
		switch statName {
		case api.AttackName:
			newStat = NewAttack(statValue)
		case api.DefenseName:
			newStat = NewDefense(statValue)
		case api.HealthName:
			newStat = NewHealth(statValue)
		case api.HungerName:
			newStat = NewHunger(statValue)
		case api.SkillName:
			newStat = NewSkill(statValue)
		case api.StaminaName:
			newStat = NewStamina(statValue)
		case api.StepName:
			newStat = NewStep(statValue)
		}
		if newStat != nil {
			defaultStats = append(defaultStats, newStat)
		}
	}
	return defaultStats
}

func NewAttack(value int) *api.Stat {
	return api.NewStat(
		api.AttackName,
		api.AttackShort,
		"attack stat used to damage",
		api.AttackName,
		value,
	)
}

func NewDefaultStatSet(name string, givenStats map[string]int, bucketset api.IBucketSet) *api.StatSet {
	defaultStats := DefaultStats(givenStats, bucketset)
	statSetName := fmt.Sprintf("stat-set/default/%s", name)
	defaultStatSet := api.NewStatSet(statSetName, defaultStats)
	return defaultStatSet
}

func NewDefense(value int) *api.Stat {
	return api.NewStat(
		api.DefenseName,
		api.DefenseShort,
		"defense stat used to defend against damage",
		api.DefenseName,
		value,
	)
}

func NewHealth(value int) *api.Stat {
	return api.NewStat(
		api.HealthName,
		api.HealthShort,
		"health is used to provide avatar life",
		api.HealthName,
		value,
	)
}

func NewHunger(value int) *api.Stat {
	return api.NewStat(
		api.HungerName,
		api.HungerShort,
		"hunger is used to measure avatar hunger and thirst",
		api.HungerName,
		value,
	)
}

func NewSkill(value int) *api.Stat {
	return api.NewStat(
		api.SkillName,
		api.SkillShort,
		"skill is used to provide skill abilities",
		api.SkillName,
		value,
	)
}

func NewStamina(value int) *api.Stat {
	return api.NewStat(
		api.StaminaName,
		api.StaminaShort,
		"stamina is used to measure avatar tireness",
		api.StaminaName,
		value,
	)
}

func NewStep(value int) *api.Stat {
	return api.NewStat(
		api.StepName,
		api.StepShort,
		"setp is used to provide avatar movement",
		api.StepName,
		value,
	)
}
