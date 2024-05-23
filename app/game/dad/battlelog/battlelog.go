// battlelog package contains a dedicated logging information for all battle
// data.
package battlelog

var (
	BLog *BattleLog = NewBattleLog()
)

// -----------------------------------------------------------------------------
//
// BattleLog
//
// -----------------------------------------------------------------------------

// BattleLog structure contains all required attributes and method to track all
// battle log information.
type BattleLog struct {
	cache []string
	index int
}

func NewBattleLog() *BattleLog {
	return &BattleLog{
		cache: []string{},
		index: 0,
	}
}

// -----------------------------------------------------------------------------
// BattleLog public methods
// -----------------------------------------------------------------------------

func (l *BattleLog) Push(str string) {
	if str != "" {
		l.cache = append(l.cache, str)
	}
}

func (l *BattleLog) Pop() string {
	var result string
	if l.IsAny() {
		result = l.cache[l.index]
		l.index++
	}
	return result
}

func (l *BattleLog) Cache() []string {
	return l.cache
}

func (l *BattleLog) Clear() {
	l.cache = []string{}
	l.index = 0
}

func (l *BattleLog) IsAny() bool {
	return len(l.cache) != l.index
}
