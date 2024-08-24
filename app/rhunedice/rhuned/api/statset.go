package api

import "fmt"

type IStatSet interface {
	GetName() string
	GetStatByName(string) IStat
	GetStats() []IStat
	SetName(string)
	SetStats([]IStat)
	String() string
}

type StatSet struct {
	stats []IStat
	name  string
}

func NewStatSet(name string, stats []IStat) *StatSet {
	return &StatSet{
		stats: stats,
		name:  name,
	}
}

func (s StatSet) getStatAndIndex(name string) (IStat, int) {
	if stat, index, found := FindByNameWithIndex(s.stats, name); found {
		return stat, index
	}
	return nil, -1
}

func (s *StatSet) GetName() string {
	return s.name
}

func (s *StatSet) GetStatByName(name string) IStat {
	stat, _ := s.getStatAndIndex(name)
	return stat
}

func (s *StatSet) GetStats() []IStat {
	return s.stats
}

func (s *StatSet) SetName(name string) {
	s.name = name
}

func (s *StatSet) SetStats(stats []IStat) {
	s.stats = stats
}

func (s *StatSet) String() string {
	return fmt.Sprintf("%s %s", s.name, s.stats)
}
