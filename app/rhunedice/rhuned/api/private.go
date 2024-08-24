package api

func getBucketCatFromRhune(rhune IRhune) EBucketCat {
	switch rhune.GetShort() {
	case AttackShort:
		return AtkBucket
	case DefenseShort:
		return DefBucket
	case SkillShort:
		return SklBucket
	default:
		return NilBucket
	}
}
