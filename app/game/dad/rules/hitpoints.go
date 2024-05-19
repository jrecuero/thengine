package rules

import "math"

// -----------------------------------------------------------------------------
//
// IHitPoints
//
// -----------------------------------------------------------------------------

// IHitPoints interface defines all methods required to be impleted for any hit
// points.
type IHitPoints interface {
	GetMaxScore() int
	SetMaxScore(int)
	GetScore() int
	SetScore(int)
	GetExtra() int
	SetExtra(int)
	Inc(int) int
	Dec(int) int
	IsAlive() bool
}

// -----------------------------------------------------------------------------
//
// HitPoints
//
// -----------------------------------------------------------------------------

// HitPoints struct defines all attributes and methods requried for a unit hit
// points.
//
// Hit points (HP) represent how much damage a character can withstand before
// falling unconscious or dying.
//
// Each character and creature has a certain number of hit points, which are
// determined by their class, level, and Constitution score. For example, a
// 1st-level fighter with a Constitution score of 14 has a base hit point total
// of 10 (the maximum roll of a 10-sided die) plus their Constitution modifier.
// As the character gains levels, they gain additional hit points based on their
// class and Constitution score.
//
// When a character or creature takes damage, their hit point total is reduced
// by the amount of damage taken. When a character's hit points are reduced to
// zero or less, they fall unconscious and are at risk of dying. If a
// character's hit points are reduced to negative their maximum hit points, they
// die outright.
//
// Characters can regain hit points through healing spells, magical potions, or
// natural rest over time. A character can also be stabilized with first aid if
// they are at 0 hit points, preventing them from dying but not restoring any
// hit points.
type HitPoints struct {
	maxScore int // life maximun score.
	score    int // life score.
	extra    int // life extra values.
}

// NewHitPoints function creates a new HitPoints instance.
func NewHitPoints(max int) *HitPoints {
	return &HitPoints{
		maxScore: max,
		score:    max,
	}
}

// -----------------------------------------------------------------------------
// HitPoints public methods
// -----------------------------------------------------------------------------

// GetMaxScore method returns the maximun hit points.
func (l *HitPoints) GetMaxScore() int {
	return l.maxScore
}

// SetMaxScore method get the maximun hit points.
func (l *HitPoints) SetMaxScore(score int) {
	l.maxScore = score
}

// GetScore method returns the hit points score.
func (l *HitPoints) GetScore() int {
	return l.score
}

// SetScore method sets the hit points score.
func (l *HitPoints) SetScore(score int) {
	l.score = score
}

// GetExtra method returns the extra maximun hit points score.
func (l *HitPoints) GetExtra() int {
	return l.extra
}

// SetExtra method sets the extra maximun hit points score.
func (l *HitPoints) SetExtra(extra int) {
	l.extra = extra
}

// Inc method adds the given value to the hit points score.
func (l *HitPoints) Inc(score int) int {
	l.score = int(math.Min(float64(l.score+score), float64(l.maxScore)))
	return l.score
}

// Dec method substracts the given value to the hit points score.
func (l *HitPoints) Dec(score int) int {
	l.score = int(math.Max(0, float64(l.score+score)))
	return l.score
}

// IsAlive method returns if hit points score is larger that zero.
func (l *HitPoints) IsAlive() bool {
	return l.score > 0
}

var _ IHitPoints = (*HitPoints)(nil)
