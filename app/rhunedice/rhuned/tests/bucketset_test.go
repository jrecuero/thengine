package tests

import (
	"fmt"
	"testing"

	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
)

func TestBucketSet(t *testing.T) {
	// create dice faces
	faces := []api.IFace{
		api.AttackFace,
		api.DefenseFace,
		api.SkillFace,
	}
	fmt.Println(faces)

	// create multiple dice
	dices := make([]api.IDice, 3)
	for i := range dices {
		diceName := fmt.Sprintf("dice/%d", i)
		dices[i] = api.NewDice(diceName, faces)
	}
	fmt.Println(dices)

	// create dice-set
	diceset := api.NewDiceSet("diceset/1", dices)
	fmt.Println(diceset)

	// roll dice-set
	roll := diceset.Roll()
	fmt.Println(roll)

	// create buckets
	buckets := []api.IBucket{
		api.NewBucket("atk/1", api.AttackRhune),
		api.NewBucket("def/1", api.DefenseRhune),
		api.NewBucket("skl/1", api.SkillRhune),
	}
	fmt.Println(buckets)

	// create bucket-set
	bucketset := api.NewBucketSet("bucket-set/1", buckets)
	fmt.Println(bucketset)

	// update bucket-set with dice-set roll
	bucketset.UpdateBucketsFromDiceSetRoll(roll)
	fmt.Println(bucketset)

	fmt.Println()
}
