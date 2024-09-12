package api

type EBucketCat int

const (
	AttackBucket EBucketCat = iota
	DefenseBucket
	SkillBucket
	StaminaBucket
	HealthBucket
	StepBucket
	ExtraBucket
	NilBucket
)

func (b EBucketCat) String() string {
	switch b {
	case AttackBucket:
		return AttackName
	case DefenseBucket:
		return DefenseName
	case SkillBucket:
		return SkillName
	case StaminaBucket:
		return StaminaName
	case HealthBucket:
		return HealthName
	case StepBucket:
		return StepName
	case ExtraBucket:
		return ExtraName
	case NilBucket:
		return NilName
	default:
		return NilName
	}
}

func (b EBucketCat) IsBase() bool {
	switch b {
	case AttackBucket:
		return true
	case DefenseBucket:
		return true
	case SkillBucket:
		return true
	case StaminaBucket:
		return true
	case HealthBucket:
		return true
	case StepBucket:
		return true
	case ExtraBucket:
		return false
	default:
		return false
	}
}

func (b EBucketCat) IsExtra() bool {
	switch b {
	case AttackBucket:
		return false
	case DefenseBucket:
		return false
	case SkillBucket:
		return false
	case StaminaBucket:
		return false
	case HealthBucket:
		return false
	case StepBucket:
		return false
	case ExtraBucket:
		return true
	default:
		return false
	}
}

func (b EBucketCat) Equal(other IComparable) bool {
	if o, ok := other.(EBucketCat); ok {
		return b == o
	}
	return false
}
