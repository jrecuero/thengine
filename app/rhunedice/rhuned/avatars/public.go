package avatars

import (
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/buckets"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/dicesets"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/stats"
)

func DefaultAvatar(name string, givenStats map[string]int) *api.Avatar {
	// create default bucket-set
	defaultBucketSet := buckets.NewDefaultBucketSet(name)

	// create avatar
	defaultAvatar := api.NewAvatar(
		name,
		stats.NewDefaultStatSet(name, givenStats, defaultBucketSet),
		dicesets.NewFiveAndOneDiceSet(name),
		defaultBucketSet,
		api.NewEquipment(nil, nil, nil),
		nil)

	return defaultAvatar
}
