// level.go package contains all data and logic required to handle level and
// experience for any player character.
package rules

import "math"

// -----------------------------------------------------------------------------
//
// ILevel
//
// -----------------------------------------------------------------------------

// ILevel interface defines all methods required for any level.
type ILevel interface {
	DecExperience(int) int
	GetExperience() int
	GetScore() int
	GetToGive() int
	GetToNext() int
	IncExperience(int) int
	LevelDown(int, bool) int
	LevelUp(int, bool) int
	SetExperience(int)
	SetScore(int)
	SetToGive(int)
}

// -----------------------------------------------------------------------------
//
// Level
//
// -----------------------------------------------------------------------------

// Level struct defines all attributes and methods required for level a unit
// and experience need.
//
// Level refers to a character's overall experience and power. Characters start
// at 1st level and can advance up to a maximum of 20th level through gaining
// experience points (XP) from combat, exploration, and completing quests.
//
// As a character gains levels, they become more powerful and gain new
// abilities, spells, hit points, and proficiencies. For example, a 3rd-level
// wizard gains access to 2nd-level spells, while a 6th-level fighter gains an
// extra attack when using the Attack action.
//
// A character's level also affects their proficiency bonus, which is a bonus
// added to certain rolls based on their level. The proficiency bonus starts at
// +2 for 1st-level characters and increases to a maximum of +6 at 17th level.
//
// Level is an important aspect of character progression in D&D and can affect
// a character's abilities, equipment, and overall effectiveness in combat and
// other challenges.
type Level struct {
	experience int // level experience.
	score      int // level score.
	togive     int // experience to give to player character
}

// NewLevel functions creates a new Level instance.
func NewLevel(score int, exp int, togive int) *Level {
	level := &Level{
		experience: exp,
		score:      score,
		togive:     togive,
	}
	tonext := level.GetToNext()
	level.SetExperience(int(math.Min(float64(exp), float64(tonext-1))))
	return level
}

// -----------------------------------------------------------------------------
// Level public methods
// -----------------------------------------------------------------------------

// DecExperience method decreases the experience score with the given value.
func (l *Level) DecExperience(exp int) int {
	l.experience -= exp
	for l.experience < l.ToNext(l.GetScore()) {
		l.LevelDown(1, false)
	}
	return l.experience
}

// GetExperience method returns the actual experience score.
func (l *Level) GetExperience() int {
	return l.experience
}

// GetScore method returns the actual level score.
func (l *Level) GetScore() int {
	return l.score
}

// GetToGive method returns the experience the player character will get when
// defeating the unit with this Level instance.
func (l *Level) GetToGive() int {
	return l.togive
}

// GetToNext method returns the experience score required to get to the next
// level.
func (l *Level) GetToNext() int {
	return l.ToNext(l.GetScore() + 1)
}

// IncExperience method increaases the experience score with the given value.
func (l *Level) IncExperience(exp int) int {
	l.experience += exp
	for l.experience > l.GetToNext() {
		l.LevelUp(1, false)
	}
	return l.experience
}

// LevelDown method levels down the level score the given value.
func (l *Level) LevelDown(score int, updateExp bool) int {
	for i := 0; i < score; i++ {
		l.score--
		if updateExp {
			l.experience = l.ToNext(l.score)
		}
	}
	return l.score
}

// LevelUp method levels up the level score the given value.
func (l *Level) LevelUp(score int, updateExp bool) int {
	for i := 0; i < score; i++ {
		if updateExp {
			l.experience += l.GetToNext()
		}
		l.score++
	}
	return l.score
}

// SetExperience method sets the actual experience score.
func (l *Level) SetExperience(exp int) {
	l.experience = exp
}

// SetScore method sets the actual level score.
func (l *Level) SetScore(score int) {
	l.score = score
}

// SetToGive method sets the experience the player character will get when
// defeating the unit with this Level instance.
func (l *Level) SetToGive(togive int) {
	l.togive = togive
}

// ToNext method returns the default value for the experience required to
// reach a given level.
func (l *Level) ToNext(level int) int {
	return level * 1000
}

var _ ILevel = (*Level)(nil)
