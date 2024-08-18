package api

func getBucketCatFromRhune(rhune IRhune) EBucketCat {
	switch rhune.GetShort() {
	case "ATK":
		return AtkBucket
	case "DEF":
		return DefBucket
	case "SKL":
		return SklBucket
	default:
		return NilBucket
	}
}
