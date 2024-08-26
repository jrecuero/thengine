package api

import "fmt"

type IStat interface {
	GetBucketName() string
	GetDescription() string
	GetName() string
	GetShort() string
	GetValue() int
	SetBucketName(string)
	SetDescription(string)
	SetName(string)
	SetShort(string)
	SetValue(int)
	String() string
}

type Stat struct {
	bucketName  string
	description string
	name        string
	short       string
	value       int
}

func NewStat(name string, short string, description string,
	bucketName string, value int) *Stat {
	return &Stat{
		bucketName:  bucketName,
		description: description,
		name:        name,
		short:       short,
		value:       value,
	}
}

func (s *Stat) GetBucketName() string {
	return s.bucketName
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

func (s *Stat) SetBucketName(bucketName string) {
	s.bucketName = bucketName
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
	return fmt.Sprintf("%s %s %d", s.name, s.bucketName, s.value)
}

var _ IStat = (*Stat)(nil)
