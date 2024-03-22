// tcell.go contains all general functionality related with tcell module.
package tools

import "github.com/gdamore/tcell/v2"

// -----------------------------------------------------------------------------
// Public functions
// -----------------------------------------------------------------------------

// IsEqualStyle function checks if two styles are the same.
func IsEqualStyle(style1, style2 *tcell.Style) bool {
	fg1, bg1, attrs1 := style1.Decompose()
	fg2, bg2, attrs2 := style2.Decompose()
	return (fg1 == fg2) && (bg1 == bg2) && (attrs1 == attrs2)
}
