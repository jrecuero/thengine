package api

import "fmt"

type IStat interface {
	GetBucket() IBucket
	GetDescription() string
	GetName() string
	GetShort() string
	GetValue() int
	SetBucket(IBucket)
	SetDescription(string)
	SetName(string)
	SetShort(string)
	SetValue(int)
	String() string
}

type Stat struct {
	bucket      IBucket
	description string
	name        string
	short       string
	value       int
}

func (s *Stat) GetBucket() IBucket {
	return s.bucket
}

func (s *Stat) GetDescription() string {
	return s.description
}

func (s *Stat) GetName() string {
	return s.name
}

func (s *Stat) GetShort() string {
	return s.short
}

func (s *Stat) GetValue() int {
	return s.value
}

func (s *Stat) SetBucket(bucket IBucket) {
	s.bucket = bucket
}

func (s *Stat) SetDescription(description string) {
	s.description = description
}

func (s *Stat) SetName(name string) {
	s.name = name
}

func (s *Stat) SetShort(short string) {
	s.short = short
}

func (s *Stat) SetValue(value int) {
	s.value = value
}

func (s *Stat) String() string {
	return fmt.Sprintf("%s %s %d", s.name, s.bucket, s.value)
}

var _ IStat = (*Stat)(nil)
