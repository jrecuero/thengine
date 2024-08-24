package stats

import "github.com/jrecuero/thengine/app/rhunedice/rhuned/api"

func NewAttack(bucket api.IBucket, value int) *api.Stat {
	return api.NewStat(
		api.AttackName,
		api.AttackShort,
		"attack stat used to damage",
		bucket,
		value,
	)
}

func NewDefense(bucket api.IBucket, value int) *api.Stat {
	return api.NewStat(
		api.DefenseName,
		api.DefenseShort,
		"defense stat used to defend against damage",
		bucket,
		value,
	)
}

func NewSkill(bucket api.IBucket, value int) *api.Stat {
	return api.NewStat(
		api.SkillName,
		api.SkillShort,
		"skill is used to provide skill abilities",
		bucket,
		value,
	)
}

func DefaultStats(givenStats map[string]int, bucketset api.IBucketSet) []api.IStat {
	defaultStats := []api.IStat{}
	for statName, statValue := range givenStats {
		var newStat api.IStat = nil
		switch statName {
		case api.AttackName:
			newStat = NewAttack(
				bucketset.GetBucketByName(api.AttackName), statValue)
		case api.DefenseName:
			newStat = NewDefense(
				bucketset.GetBucketByName(api.DefenseName), statValue)
		case api.SkillName:
			newStat = NewSkill(
				bucketset.GetBucketByName(api.SkillName), statValue)
		}
		if newStat != nil {
			defaultStats = append(defaultStats, newStat)
		}
	}
	return defaultStats
}
