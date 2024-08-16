// point.go contains everything required to identify a point in the screeen or
// the canvas using the horizontal and vertical position in the canvas.
package api

import (
	"fmt"
	"math"
)

// -----------------------------------------------------------------------------
//
// Point
//
// -----------------------------------------------------------------------------

// Point structure identifies any position in the screen or canvas by the
// horizontal and vertical values.
// X integer with the horizontal location.
// Y integer with the vertical location.
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// NewPoint function creates a new Point instaces based on the given horizontal
// and vertical locations.
func NewPoint(x int, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

// ClonePoint functions creates a new Point instances with same attributes as
// the given Point.
func ClonePoint(point *Point) *Point {
	return &Point{
		X: point.X,
		Y: point.Y,
	}
}

// -----------------------------------------------------------------------------
// Point public methods
// -----------------------------------------------------------------------------

// Add method adds the given point coordinates to the point instance.
func (p *Point) Add(point *Point) {
	p.X += point.X
	p.Y += point.Y
}

// AddScale method adds given X and Y values to the point instance.
func (p *Point) AddScale(x int, y int) {
	p.X += x
	p.Y += y
}

// Clone method clones all attributes from the given Point instance.
func (p *Point) Clone(point *Point) {
	p.X = point.X
	p.Y = point.Y
}

// Distance method returns the distance with the given point.
func (p *Point) Distance(point *Point) (float64, int, int) {
	dx := point.X - p.X
	dy := point.Y - p.Y
	distance := math.Sqrt(float64(dx*dx + dy*dy))
	return distance, dx, dy
}

// Get method returns horizontal and vertical location for the instance.
func (p *Point) Get() (int, int) {
	return p.X, p.Y
}

func (p *Point) GetAdjacentPoints() []*Point {
	result := []*Point{
		NewPoint(p.X, p.Y-1),
		NewPoint(p.X, p.Y+1),
		NewPoint(p.X-1, p.Y),
		NewPoint(p.X+1, p.Y),
	}
	return result
}

// IsAdjacent method returns if the given Point is adjacent to the instance, it
// means if it is just on top, bottom, left or right.
func (p *Point) IsAdjacent(point *Point) bool {
	return ((p.X == point.X) && (p.Y+1 == point.Y)) ||
		((p.X == point.X) && (p.Y-1 == point.Y)) ||
		((p.X+1 == point.X) && (p.Y == point.Y)) ||
		((p.X-1 == point.X) && (p.Y == point.Y))
}

// IsEqual method returns if the given Point is equal than the instance, based
// on the same horizontal and vertival coordinates.
func (p *Point) IsEqual(point *Point) bool {
	return (p.X == point.X) && (p.Y == point.Y)
}

// SaveToDict method saves the instance information as a map.
func (p *Point) SaveToDict() map[string]any {
	result := map[string]any{}
	result["x"] = p.X
	result["y"] = p.Y
	return result
}

// Set method assigns new horizontal and vertical locations with given values.
func (p *Point) Set(x int, y int) {
	p.X = x
	p.Y = y
}

// Subtract method subtracts the given point coordinates to the point instance.
func (p *Point) Subtract(point *Point) {
	p.X -= point.X
	p.Y -= point.Y
}

// ToString method returns instance information as a string.
func (p *Point) ToString() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}
