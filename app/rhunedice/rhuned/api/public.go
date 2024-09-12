package api

const (
	AttackName  = "attack"
	DefenseName = "defense"
	SkillName   = "skill"
	StaminaName = "stamina"
	HealthName  = "health"
	StepName    = "step"
	ExtraName   = "extra"
	NilName     = "nil"

	ClothesName   = "clothes"
	KnowledgeName = "knowledge"

	AttackShort  = "ATK"
	DefenseShort = "DEF"
	SkillShort   = "SKL"
	StaminaShort = "STA"
	HealthShort  = "HLT"
	StepShort    = "STP"
	ExtraShort   = "EXT"
	NilShort     = "NIL"

	ClothesShort   = "CL1"
	KnowledgeShort = "KNO"
)

func FacesToBuckets(faces []IFace) []IBucket {
	buckets := []IBucket{}
	ebuckets := map[EBucketCat]any{}

	for _, face := range faces {
		bucketCat := face.GetRhune().GetBucketCat()
		if bucketCat == NilBucket {
			continue
		}
		if bucketCat != ExtraBucket {
			if _, ok := ebuckets[bucketCat]; ok {
				ebuckets[bucketCat] = ebuckets[bucketCat].(int) + 1
			} else {
				ebuckets[bucketCat] = 1
			}
		} else {
			ebuckets[bucketCat] = face.GetRhune()
		}
	}
	for k, v := range ebuckets {
		var bucket IBucket
		if k != ExtraBucket {
			bucket = NewBucketWithValue(k.String(), k, v.(int))
		} else {
			bucket = NewBucketWithRhune(k.String(), k, v.(IRhune))
		}
		buckets = append(buckets, bucket)
	}

	return buckets
}
