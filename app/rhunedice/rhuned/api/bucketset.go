package api

import "fmt"

type IBucketSet interface {
	AddBucket(IBucket) error
	GetBucketByName(string) IBucket
	GetBuckets() []IBucket
	GetBucketsForCat(EBucketCat) []IBucket
	GetName() string
	RemoveBucket(IBucket)
	SetBuckets([]IBucket)
	SetName(string)
	String() string
	UpdateBucketsFromBuckets([]IBucket)
	UpdateBucketsFromDiceSetRoll([]IFace)
}

type BucketSet struct {
	buckets []IBucket
	name    string
}

func NewBucketSet(name string, buckets []IBucket) *BucketSet {
	return &BucketSet{
		buckets: buckets,
		name:    name,
	}
}

func (b *BucketSet) getBucketAndIndex(name string) (IBucket, int) {
	if bucket, index, found := FindByNameWithIndex(b.buckets, name); found {
		return bucket, index
	}
	return nil, -1
}

func (b *BucketSet) AddBucket(bucket IBucket) error {
	if found := b.GetBucketByName(bucket.GetName()); found != nil {
		return fmt.Errorf("Bucket %s found in BucketSet %s", bucket.GetName(), b.name)
	}
	b.buckets = append(b.buckets, bucket)
	return nil
}

func (b *BucketSet) GetBucketByName(name string) IBucket {
	bucket, _ := b.getBucketAndIndex(name)
	return bucket
}

func (b *BucketSet) GetBuckets() []IBucket {
	return b.buckets
}

func (b *BucketSet) GetBucketsForCat(cat EBucketCat) []IBucket {
	return FindByCat(b.buckets, cat)
}

func (b *BucketSet) GetName() string {
	return b.name
}

func (b *BucketSet) RemoveBucket(bucket IBucket) {
	if bucket, index := b.getBucketAndIndex(bucket.GetName()); bucket != nil {
		b.buckets = append(b.buckets[:index], b.buckets[index+1:]...)
	}
}

func (b *BucketSet) SetBuckets(buckets []IBucket) {
	b.buckets = buckets
}

func (b *BucketSet) SetName(name string) {
	b.name = name
}

// UpdateBucketsFromBuckets method updates the bucketset with the list of
// provided buckets.
func (b *BucketSet) UpdateBucketsFromBuckets(buckets []IBucket) {
	for _, bucket := range buckets {
		cat := bucket.GetCat().(EBucketCat)
		if cat == ExtraBucket {
			bucket.SetRhune(bucket.GetRhune())
		} else {
			for _, buck := range b.GetBucketsForCat(cat) {
				buck.Inc(bucket.GetValue())
			}
		}
	}
}

// UpdateBucketsFromDiceSetRoll method updates the bucketset with the list of
// faces from a rolldice.
func (b *BucketSet) UpdateBucketsFromDiceSetRoll(roll []IFace) {
	for _, face := range roll {
		rhune := face.GetRhune()
		cat := rhune.GetBucketCat()
		for _, bucket := range b.GetBucketsForCat(cat) {
			if cat != ExtraBucket {
				bucket.Inc(1)
			} else {
				bucket.SetRhune(rhune)
			}
		}
	}
}

func (b *BucketSet) String() string {
	return fmt.Sprintf("%s %s", b.name, b.buckets)
}

var _ IBucketSet = (*BucketSet)(nil)
