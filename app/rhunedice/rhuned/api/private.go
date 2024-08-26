package api

func getBucketCatFromRhune(rhune IRhune) EBucketCat {
	switch rhune.GetShort() {
	case AttackShort:
		return AttackBucket
	case DefenseShort:
		return DefenseBucket
	case SkillShort:
		return SkillBucket
	default:
		return NilBucket
	}
}
