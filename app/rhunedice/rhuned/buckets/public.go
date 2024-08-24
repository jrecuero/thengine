package buckets

import (
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/rhunes"
)

func DefaultBuckets() []api.IBucket {
	defaultBuckets := []api.IBucket{
		api.NewBucket(api.AttackName, rhunes.AttackRhune),
		api.NewBucket(api.DefenseName, rhunes.DefenseRhune),
		api.NewBucket(api.SkillName, rhunes.SkillRhune),
	}
	return defaultBuckets
}
