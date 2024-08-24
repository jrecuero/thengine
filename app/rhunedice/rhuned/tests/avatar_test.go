package tests

import (
	"fmt"
	"testing"

	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/avatars"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/rhunes"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/stats"
)

func TestAvatar(t *testing.T) {
	// create buckets
	buckets := []api.IBucket{
		api.NewBucket(api.AttackName, rhunes.AttackRhune),
		api.NewBucket(api.DefenseName, rhunes.DefenseRhune),
		api.NewBucket(api.SkillName, rhunes.SkillRhune),
	}
	fmt.Println("buckets ", buckets)

	// create bucket-set
	bucketset := api.NewBucketSet("bucket-set/1", buckets)
	fmt.Println("bucketset ", bucketset)

	// create stats
	avatarStats := []api.IStat{
		stats.NewAttack(bucketset.GetBucketByName(api.AttackName), 1),
		stats.NewDefense(bucketset.GetBucketByName(api.DefenseName), 1),
		stats.NewSkill(bucketset.GetBucketByName(api.SkillName), 1),
	}

	// create stat-set
	statset := api.NewStatSet("stats-set/1", avatarStats)

	// create avatar
	avatar := api.NewAvatar("avatar/1", statset, bucketset, nil, nil)
	fmt.Println("avatar ", avatar)
}

func TestDefaultAvatar(t *testing.T) {
	// create map from stat name to stat value
	statsmap := map[string]int{
		api.AttackName:  2,
		api.DefenseName: 1,
		api.SkillName:   3,
	}

	// create default avatar
	avatar := avatars.DefaultAvatar("test/1", statsmap)
	fmt.Println("default-avatar ", avatar)

	avatar.StartTurn()
	fmt.Println("default-avatar ", avatar)
}
