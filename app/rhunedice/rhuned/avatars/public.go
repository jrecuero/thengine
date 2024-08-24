package avatars

import (
	"fmt"

	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/buckets"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/stats"
)

func DefaultAvatar(name string, givenStats map[string]int) *api.Avatar {
	// create default buckets
	defaultBuckets := buckets.DefaultBuckets()

	// create default bucket-set
	bucketSetName := fmt.Sprintf("bucket-set/%s", name)
	defaultBucketSet := api.NewBucketSet(bucketSetName, defaultBuckets)

	// create stats
	defaultStats := stats.DefaultStats(givenStats, defaultBucketSet)

	// create stat-set
	statSetName := fmt.Sprintf("stat-set/%s", name)
	defaultStatSet := api.NewStatSet(statSetName, defaultStats)

	// create avatar
	defaultAvatar := api.NewAvatar(
		name,
		defaultStatSet,
		defaultBucketSet,
		api.NewEquipment(nil, nil, nil),
		nil)

	return defaultAvatar
}
