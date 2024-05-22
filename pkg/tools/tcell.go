// tcell.go contains all general functionality related with tcell module.
package tools

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

// -----------------------------------------------------------------------------
// Public functions
// -----------------------------------------------------------------------------

// IsEqualStyle function checks if two styles are the same.
func IsEqualStyle(style1, style2 *tcell.Style) bool {
	fg1, bg1, attrs1 := style1.Decompose()
	fg2, bg2, attrs2 := style2.Decompose()
	return (fg1 == fg2) && (bg1 == bg2) && (attrs1 == attrs2)
}

// ReverseStyle function reverses the style switching foreground with
// background colors.
func ReverseStyle(style *tcell.Style) *tcell.Style {
	if style == nil {
		return nil
	}
	fg, bg, attrs := style.Decompose()
	reversed := tcell.StyleDefault.Foreground(bg).Background(fg).Attributes(attrs)
	return &reversed
}

// StyleToString function display style as a string.
func StyleToString(style *tcell.Style) string {
	fg, bg, attrs := style.Decompose()
	return fmt.Sprintf("%s/%s/%d", fg.String(), bg.String(), attrs)
}
