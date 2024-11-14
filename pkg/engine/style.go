// style.go wraps-up tcell.Style in order to provide further functionality for
// handling colors and attributes.
package engine

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

// -----------------------------------------------------------------------------
// Package public methods
// -----------------------------------------------------------------------------

// CloneStyle function clones the given tcell.Style to a new instance with same
// values.
func CloneStyle(style *tcell.Style) *tcell.Style {
	fg, bg, attrs := style.Decompose()
	newStyle := NewStyle(fg, bg, attrs)
	return newStyle
}

// CompareStyle function checks if foreground, background and attribute colors
// are equal in both styles.
func CompareStyle(style1 *tcell.Style, style2 *tcell.Style) bool {
	if style1 == style2 {
		return true
	}
	if (style1 == nil) || (style2 == nil) {
		return false
	}
	return (GetForegroundFromStyle(style1) == GetForegroundFromStyle(style2)) &&
		(GetBackgroundColorFromStyle(style1) == GetBackgroundColorFromStyle(style2)) &&
		(GetAttrsFromStyle(style1) == GetAttrsFromStyle(style2))
}

// CompareStyleColors function checks if foreground and background colors are
// equal in both styles.
func CompareStyleColor(style1 *tcell.Style, style2 *tcell.Style) bool {
	return (GetForegroundFromStyle(style1) == GetForegroundFromStyle(style2)) &&
		(GetBackgroundColorFromStyle(style1) == GetBackgroundColorFromStyle(style2))
}

// GetAttrsFromStyle function returns color attributes from the given style.
func GetAttrsFromStyle(s *tcell.Style) tcell.AttrMask {
	_, _, attrs := s.Decompose()
	return attrs
}

// GetBackgroundColorFromStyle function returns the background color from the
// given style.
func GetBackgroundColorFromStyle(s *tcell.Style) tcell.Color {
	_, bg, _ := s.Decompose()
	return bg
}

// GetForegroundFromStyle function returns the foreground color from the given
// style.
func GetForegroundFromStyle(s *tcell.Style) tcell.Color {
	fg, _, _ := s.Decompose()
	return fg
}

// NewStyle function creates a new tcell.Style instance.
func NewStyle(fg tcell.Color, bg tcell.Color, attrs tcell.AttrMask) *tcell.Style {
	style := tcell.StyleDefault.Foreground(fg).Background(bg).Attributes((attrs))
	return &style
}

func StyleToString(s *tcell.Style) string {
	if s == nil {
		return "[nil:nil:nil]"
	}
	fg, bg, attrs := s.Decompose()
	return fmt.Sprintf("[%s:%s:%d]", fg.String(), bg.String(), attrs)
}
