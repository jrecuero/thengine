// rect.go contains everything required to handle a rectangle in the
// application. A rectangle is identified by an origin point at he top left
// of the rectangle and a size with width and height.
package api

import (
	"fmt"
)

// -----------------------------------------------------------------------------
//
// Rect
//
// -----------------------------------------------------------------------------

// Rect structure defines all attributes required to identify a rectangle by
// its position and size.
// Origin *Point instance identifies the rectangle top left corner.
// Size *Size instance identifies rectangle width and height.
type Rect struct {
	Origin *Point
	Size   *Size
}

// NewRect function creates a new Rect instance at the given origin and with
// the given size.
func NewRect(origin *Point, size *Size) *Rect {
	return &Rect{
		Origin: origin,
		Size:   size,
	}
}

// CloneRect function creates a new Rect instance using same attribute values
// as the given Rect instance.
func CloneRect(rect *Rect) *Rect {
	return &Rect{
		Origin: rect.Origin,
		Size:   rect.Size,
	}
}

// -----------------------------------------------------------------------------
// Rect public methods
// -----------------------------------------------------------------------------

// Clone method clones all attributed from the given Rect to the instance.
func (r *Rect) Clone(rect *Rect) {
	r.Origin = rect.Origin
	r.Size = rect.Size
}

// Set method assigns new origin and size to the instance.
func (r *Rect) Set(origin *Point, size *Size) {
	r.Origin = origin
	r.Size = size
}

// SetOrigin method assigns a new origin to the instance.
func (r *Rect) SetOrigin(origin *Point) {
	r.Origin = origin
}

// SetSize method assigns a new size to the instance.
func (r *Rect) SetSize(size *Size) {
	r.Size = size
}

// Get method returns the instance origin and size.
func (r *Rect) Get() (*Point, *Size) {
	return r.Origin, r.Size
}

// GetOrigin method returns the instance origin.
func (r *Rect) GetOrigin() *Point {
	return r.Origin
}

// GetSize method returns the instance size.
func (r *Rect) GetSize() *Size {
	return r.Size
}

// IsEqual method checks if the given Rect is equal to the instance bases in
// the same origin coordinates and the same size.
func (r *Rect) IsEqual(rect *Rect) bool {
	return r.Origin.IsEqual(rect.Origin) && r.Size.IsEqual(rect.Size)
}

// IsInside method checks if the given point is inside the rectangle.
func (r *Rect) IsInside(point *Point) bool {
	return point.X >= r.Origin.X &&
		point.X < (r.Origin.X+r.Size.W) &&
		point.Y >= r.Origin.Y &&
		point.Y < (r.Origin.Y+r.Size.H)
}

// IsBorder method chekcs if the given point is any border of the rectangle.
func (r *Rect) IsBorder(point *Point) bool {
	if (point.X == r.Origin.X) || (point.X == (r.Origin.X + r.Size.W - 1)) {
		return (point.Y >= r.Origin.Y) && (point.Y <= (r.Origin.Y + r.Size.H - 1))
	}
	if (point.Y == r.Origin.Y) || (point.Y == (r.Origin.Y + r.Size.H - 1)) {
		return (point.X >= r.Origin.X) && (point.X <= (r.Origin.X + r.Size.W - 1))
	}
	return false
}

// ToString method returns instance information as a string.
func (r *Rect) ToString() string {
	return fmt.Sprintf("%s/%s", r.Origin.ToString(), r.Size.ToString())
}

// SaveToDict method saves the instance information as a map.
func (r *Rect) SaveToDict() map[string]any {
	result := map[string]any{}
	result["origin"] = r.Origin.SaveToDict()
	result["size"] = r.Size.SaveToDict()
	return result
}
