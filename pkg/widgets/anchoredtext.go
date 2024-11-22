// anchoredtext.go contains a sequence of anchored Text Widgets logic.
package widgets

import "github.com/jrecuero/thengine/pkg/api"

// -----------------------------------------------------------------------------
//
// AnchoredText
//
// -----------------------------------------------------------------------------

type AnchoredText struct {
	*Group
}

// -----------------------------------------------------------------------------
// New AnchoredText functions
// -----------------------------------------------------------------------------

func NewAnchoredText(name string, texts ...IWidget) *AnchoredText {
	group := &AnchoredText{
		Group: NewGroup(name, texts...),
	}
	group.Refresh()
	return group
}

// -----------------------------------------------------------------------------
// AnchoredText private methods
// -----------------------------------------------------------------------------

func (t *AnchoredText) updateGroup() {
	var anchor *api.Point
	for _, w := range t.GetWidgets() {
		t, ok := w.(*Text)
		if !ok {
			return
		}
		if anchor != nil {
			t.SetPosition(anchor)
		}
		anchor = t.SetAnchor()
	}
}

// -----------------------------------------------------------------------------
// AnchoredText public methods
// -----------------------------------------------------------------------------

// Refresh method refreshes the AnchoredText widget with latest attribute values.
func (t *AnchoredText) Refresh() {
	t.updateGroup()
}
