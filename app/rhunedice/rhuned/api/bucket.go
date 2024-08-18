package api

import (
	"fmt"
	"math"
)

type IBucket interface {
	Dec(int)
	Inc(int)
	GetCat() EBucketCat
	GetLimit() int
	GetName() string
	GetRhune() IRhune
	GetValue() int
	SetLimit(int)
	SetName(string)
	SetValue(int)
	String() string
}

type Bucket struct {
	cat   EBucketCat
	limit int
	name  string
	rhune IRhune
	value int
}

func NewBucket(name string, rhune IRhune) *Bucket {
	return &Bucket{
		cat:   getBucketCatFromRhune(rhune),
		limit: math.MaxInt,
		name:  name,
		rhune: rhune,
		value: 0,
	}
}

func (b *Bucket) Dec(value int) {
	if b.value > 0 {
		b.value -= value
	}
}

func (b *Bucket) Inc(value int) {
	if b.value < b.limit {
		b.value += value
	}
}

func (b *Bucket) GetCat() EBucketCat {
	return b.cat
}

func (b *Bucket) GetLimit() int {
	return b.limit
}

func (b *Bucket) GetName() string {
	return b.name
}

func (b *Bucket) GetRhune() IRhune {
	return b.rhune
}

func (b *Bucket) GetValue() int {
	return b.value
}

func (b *Bucket) SetLimit(limit int) {
	b.limit = limit
}

func (b *Bucket) SetName(name string) {
	b.name = name
}

func (b *Bucket) SetValue(value int) {
	b.value = value
}

func (b *Bucket) String() string {
	return fmt.Sprintf("%s %s %d", b.name, b.cat, b.value)
}

var _ IBucket = (*Bucket)(nil)
