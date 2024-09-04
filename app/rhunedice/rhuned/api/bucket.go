package api

import (
	"fmt"
	"math"
)

type IBucket interface {
	Dec(any)
	Inc(any)
	GetCat() IComparable
	GetLimit() int
	GetName() string
	GetRhune() IRhune
	GetValue() int
	SetLimit(int)
	SetName(string)
	SetRhune(IRhune)
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

func NewBucket(name string, cat EBucketCat) *Bucket {
	return &Bucket{
		cat:   cat,
		limit: math.MaxInt,
		name:  name,
		rhune: nil,
		value: 0,
	}
}

func NewBucketWithRhune(name string, cat EBucketCat, rhune IRhune) *Bucket {
	return &Bucket{
		cat:   cat,
		limit: math.MaxInt,
		name:  name,
		rhune: rhune,
		value: 0,
	}
}

func NewBucketWithValue(name string, cat EBucketCat, val int) *Bucket {
	return &Bucket{
		cat:   cat,
		limit: math.MaxInt,
		name:  name,
		rhune: nil,
		value: val,
	}
}

func (b *Bucket) Dec(value any) {
	if b.rhune != nil {
		b.rhune = nil
	}
	if b.value > 0 {
		b.value -= value.(int)
	}
}

func (b *Bucket) Inc(value any) {
	if b.rhune != nil {
		b.rhune = value.(IRhune)
	}
	if b.value < b.limit {
		b.value += value.(int)
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

func (b *Bucket) SetRhune(rhune IRhune) {
	b.rhune = rhune
}

func (b *Bucket) SetValue(value int) {
	b.value = value
}

func (b *Bucket) String() string {
	return fmt.Sprintf("%s %s %d", b.name, b.cat, b.value)
}

var _ IBucket = (*Bucket)(nil)
