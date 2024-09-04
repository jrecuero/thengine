package tests

import (
	"fmt"
	"testing"

	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/avatars"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/faces"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/stats"
)

func TestAvatar(t *testing.T) {
	// create dice faces
	diceFaces := []api.IFace{
		faces.AttackFace,
		faces.DefenseFace,
		faces.SkillFace,
	}
	fmt.Println(diceFaces)

	// create multiple dice
	dices := make([]api.IDice, 3)
	for i := range dices {
		diceName := fmt.Sprintf("dice/%d", i)
		dices[i] = api.NewDice(diceName, diceFaces)
	}
	fmt.Println(dices)

	// create dice-set
	diceset := api.NewDiceSet("diceset/1", dices)
	fmt.Println(diceset)

	// create buckets
	buckets := []api.IBucket{
		api.NewBucket(api.AttackName, api.AttackBucket),
		api.NewBucket(api.DefenseName, api.DefenseBucket),
		api.NewBucket(api.SkillName, api.SkillBucket),
	}

	fmt.Println("buckets ", buckets)

	// create bucket-set
	bucketset := api.NewBucketSet("bucket-set/1", buckets)
	fmt.Println("bucketset ", bucketset)

	// create stats
	avatarStats := []api.IStat{
		stats.NewAttack(1),
		stats.NewDefense(1),
		stats.NewSkill(1),
	}

	// create stat-set
	statset := api.NewStatSet("stats-set/1", avatarStats)

	// create avatar
	avatar := api.NewAvatar("avatar/1", statset, diceset, bucketset, nil, nil)
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

func TestDefaultEnemy(t *testing.T) {
	// create map from stat name to stat value
	statsmap := map[string]int{
		api.AttackName:  1,
		api.DefenseName: 1,
		api.SkillName:   1,
	}

	// create default enemy avatar
	avatar := avatars.DefaultEnemy("enemy/1", statsmap)
	fmt.Println("default-enemy ", avatar)

	avatar.StartTurn()
	fmt.Println("default-enemy ", avatar)
}
