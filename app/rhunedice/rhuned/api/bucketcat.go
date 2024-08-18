package api

type EBucketCat int

const (
	AtkBucket EBucketCat = iota
	DefBucket
	SklBucket
	StaBucket
	HltBucket
	StpBucket
	HngBucket
	ExtBucket
	NilBucket
)

func (b EBucketCat) String() string {
	switch b {
	case AtkBucket:
		return "attack"
	case DefBucket:
		return "defense"
	case SklBucket:
		return "skill"
	case StaBucket:
		return "stamina"
	case HltBucket:
		return "health"
	case StpBucket:
		return "step"
	case HngBucket:
		return "hunger"
	case ExtBucket:
		return "extra"
	default:
		return "invalid bucket cat"
	}
}

func (b EBucketCat) IsBase() bool {
	switch b {
	case AtkBucket:
		return true
	case DefBucket:
		return true
	case SklBucket:
		return true
	case StaBucket:
		return true
	case HltBucket:
		return true
	case StpBucket:
		return true
	case HngBucket:
		return true
	case ExtBucket:
		return false
	default:
		return false
	}
}

func (b EBucketCat) IsExtra() bool {
	switch b {
	case AtkBucket:
		return false
	case DefBucket:
		return false
	case SklBucket:
		return false
	case StaBucket:
		return false
	case HltBucket:
		return false
	case StpBucket:
		return false
	case HngBucket:
		return false
	case ExtBucket:
		return true
	default:
		return false
	}
}
