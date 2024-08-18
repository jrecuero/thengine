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
	for index, bucket := range b.buckets {
		if bucket.GetName() == name {
			return bucket, index
		}
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
	result, _ := b.getBucketAndIndex(name)
	return result
}

func (b *BucketSet) GetBuckets() []IBucket {
	return b.buckets
}

func (b *BucketSet) GetBucketsForCat(cat EBucketCat) []IBucket {
	result := []IBucket{}
	for _, bucket := range b.buckets {
		if bucket.GetCat() == cat {
			result = append(result, bucket)
		}
	}
	return result
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

func (b *BucketSet) UpdateBucketsFromDiceSetRoll(roll []IFace) {
	for _, face := range roll {
		rhune := face.GetRhune()
		cat := getBucketCatFromRhune(rhune)
		for _, bucket := range b.GetBucketsForCat(cat) {
			bucket.Inc(1)
		}
	}
}

func (b *BucketSet) String() string {
	return fmt.Sprintf("%s %s", b.name, b.buckets)
}

var _ IBucketSet = (*BucketSet)(nil)
