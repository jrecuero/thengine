package api

import (
	"fmt"
	"math"
)

type IBucket interface {
	Dec(int)
	Inc(int)
	GetCat() IComparable
	GetLimit() int
	GetName() string
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
	value int
}

func NewBucket(name string, cat EBucketCat) *Bucket {
	return &Bucket{
		cat:   cat,
		limit: math.MaxInt,
		name:  name,
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

func (b *Bucket) GetCat() IComparable {
	return b.cat
}

func (b *Bucket) GetLimit() int {
	return b.limit
}

func (b *Bucket) GetName() string {
	return b.name
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
