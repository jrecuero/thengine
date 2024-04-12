// size.go contains everything required to identify the size for any object
// added to the application. Size is measure by its width and height.
package api

import "fmt"

// -----------------------------------------------------------------------------
//
// Size
//
// -----------------------------------------------------------------------------

// Size structure identifies the size for any object in the game by its width
// and height
// W integer with the object width
// H integer with the object heght
type Size struct {
	W int
	H int
}

// NewSize function creates a new Size instance with the given width and
// height.
func NewSize(w int, h int) *Size {
	return &Size{
		W: w,
		H: h,
	}
}

// CloneSize function creates a new Size instance with same attribute values
// as the given Size instance.
func CloneSize(size *Size) *Size {
	return &Size{
		W: size.W,
		H: size.H,
	}
}

// -----------------------------------------------------------------------------
// Size public methods
// -----------------------------------------------------------------------------

// Clone method clones all attributes from the given Size instance.
func (s *Size) Clone(size *Size) {
	s.W = size.W
	s.H = size.H
}

// Set method assigns new width and height values to the instance.
func (s *Size) Set(w int, h int) {
	s.W = w
	s.H = h
}

// Get method returns width and height for the instance.
func (s *Size) Get() (int, int) {
	return s.W, s.H
}

// IsEqual method returns if the given Size is equal than the instance,
// based on the same width and height.
func (s *Size) IsEqual(size *Size) bool {
	return (s.W == size.W) && (s.H == size.H)
}

// IsZeroSize method returns if the Size instance has zero width and height.
func (s *Size) IsZeroSize() bool {
	return (s.W == 0) && (s.H == 0)
}

// ToString method returns size information as a string.
func (s *Size) ToString() string {
	return fmt.Sprintf("(%d-%d)", s.W, s.H)
}

// SaveToDict method saves the instance information as a map.
func (s *Size) SaveToDict() map[string]any {
	result := map[string]any{}
	result["w"] = s.W
	result["h"] = s.H
	return result
}
