// battlelog package contains a dedicated logging information for all battle
// data.
package battlelog

import (
	"fmt"
	"strings"
)

const (
	debugStr = "[DEBUG]"
	infoStr  = "[INFO]"
)

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

func (l *BattleLog) PushDebug(str string) {
	l.Push(fmt.Sprintf("%s %s", debugStr, str))
}

func (l *BattleLog) PushInfo(str string) {
	l.Push(fmt.Sprintf("%s %s", infoStr, str))
}

func (l *BattleLog) Pop() string {
	var result string
	if l.IsAny() {
		result = l.cache[l.index]
		l.index++
	}
	return result
}

func (l *BattleLog) PopDebug() string {
	str := l.Pop()
	if strings.HasPrefix(str, debugStr) {
		return str
	}
	return ""
}

func (l *BattleLog) PopInfo() string {
	str := l.Pop()
	if strings.HasPrefix(str, infoStr) {
		return str
	}
	return ""
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
