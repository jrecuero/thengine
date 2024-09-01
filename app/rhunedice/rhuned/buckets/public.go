package buckets

import (
	"fmt"

	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
)

func DefaultBuckets() []api.IBucket {
	defaultBuckets := []api.IBucket{
		api.NewBucket(api.AttackName, api.AttackBucket),
		api.NewBucket(api.DefenseName, api.DefenseBucket),
		api.NewBucket(api.SkillName, api.SkillBucket),
		api.NewBucket(api.StaminaName, api.StaminaBucket),
		api.NewBucket(api.HealthName, api.HealthBucket),
		api.NewBucket(api.StepName, api.StepBucket),
		api.NewBucket(api.HungerName, api.HungerBucket),
		api.NewBucket(api.ExtraName, api.ExtraBucket),
	}
	return defaultBuckets
}

func NewDefaultBucketSet(name string) *api.BucketSet {
	// create default bucket-set
	bucketSetName := fmt.Sprintf("bucket-set/default/%s", name)
	defaultBucketSet := api.NewBucketSet(bucketSetName, DefaultBuckets())

	return defaultBucketSet
}
