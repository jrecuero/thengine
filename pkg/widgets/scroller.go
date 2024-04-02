// scroller.go contains the internal scroller structure and functionality that
// allows other widget to scroll, horizontally or verticaly, through a
// selection of entries.
package widgets

// -----------------------------------------------------------------------------
//
// Scroller
//
// -----------------------------------------------------------------------------

// Scroller structure defines the baseline for any widget using a scroller.
// It can be an horizontal or vertical scroller.
// Vertical scroller sets SelectionLength as always ONE, because it is based on
// lines, where MaxLength is the number of lines that can be displayed at one
// time and TotalSelecionLength is the total number of selections.
// StartSelection and EndSelection represent the index values for the first and
// the last selection to be displayed.
// Horizontal scroller in other way is based in the number of characters and
// not in the number of lines. SelectionLength is the length every selection
// should take, where MaxLength is the maximum number of characters for all
// selections to be displayed at any time, where TotalSelectionLength is the
// total number of characters for all possible selections. StartSelection and
// EndSelection represent the number of characters to shift for the first
// selection and the last selection to be displayed.
type Scroller struct {
	TotalSelectionLength int
	MaxLength            int
	SelectionLength      int
	StartSelection       int
	EndSelection         int
}

// NewScroller function creates a new Scroller instance.
func NewScroller(totalSelectionLength int, maxLength int, selectionLength int) *Scroller {
	scroller := &Scroller{
		TotalSelectionLength: totalSelectionLength,
		MaxLength:            maxLength,
		SelectionLength:      selectionLength,
	}
	if totalSelectionLength > maxLength {
		scroller.StartSelection = 0
		scroller.EndSelection = (maxLength / selectionLength) - 1
	} else {
		scroller.StartSelection = 0
		scroller.EndSelection = (totalSelectionLength / selectionLength) - 1
	}
	return scroller
}

// NewVerticalScroller function creates a new vertival Scroller instance, where
// the selection Length is fixed to one.
func NewVerticalScroller(totalNbrOfSelections int, nbrOfSelectionToDisplay int) *Scroller {
	return NewScroller(totalNbrOfSelections, nbrOfSelectionToDisplay, 1)
}

// -----------------------------------------------------------------------------
// Scroller Public methods
// -----------------------------------------------------------------------------

// Update method update the scroller StartSelection and EndSelection based on
// the given selection value.
func (s *Scroller) Update(selectionIndex int) {
	if selectionIndex < s.StartSelection {
		s.StartSelection = selectionIndex
		s.EndSelection = s.StartSelection + (s.MaxLength / s.SelectionLength) - 1
	} else if selectionIndex > s.EndSelection {
		s.EndSelection = selectionIndex
		s.StartSelection = s.EndSelection - (s.MaxLength / s.SelectionLength) + 1
	}
}
