package tests

import (
	"fmt"
	"testing"

	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/faces"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/rhunes"
)

func TestBucketSet(t *testing.T) {
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

	// roll dice-set
	roll := diceset.Roll()
	fmt.Println(roll)

	// create buckets
	buckets := []api.IBucket{
		api.NewBucket("atk/1", rhunes.AttackRhune),
		api.NewBucket("def/1", rhunes.DefenseRhune),
		api.NewBucket("skl/1", rhunes.SkillRhune),
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
