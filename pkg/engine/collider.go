// collider.go contains all logic required for handling collision. Collisions
// occur between entities that are physical and dynamic. Any physical entity
// can collide with any other. Dynamic entities are those that can move across
// the scenario.
package engine

import "github.com/jrecuero/thengine/pkg/api"

// -----------------------------------------------------------------------------
//
// Collider
//
// -----------------------------------------------------------------------------

// Collider structure defines the data required for checking collisions.
type Collider struct {
	rect   *api.Rect
	points []*api.Point
}

// NewCollider function creates a Collider instance.
func NewCollider(rect *api.Rect, points []*api.Point) *Collider {
	return &Collider{
		rect:   rect,
		points: points,
	}
}

// -----------------------------------------------------------------------------
// Collider private functions
// -----------------------------------------------------------------------------

func (c *Collider) collidePointsWithPoints(collider *Collider) bool {
	for _, p1 := range c.GetPoints() {
		for _, p2 := range collider.GetPoints() {
			if p1.IsEqual(p2) {
				return true
			}
		}
	}
	return false
}

func (c *Collider) collideRectWithPoints(collider *Collider) bool {
	for _, point := range collider.GetPoints() {
		if c.GetRect().IsInside(point) {
			return true
		}
	}
	return false
}

func (c *Collider) collideRectWithRect(collider *Collider) bool {
	return c.GetRect().IsRectIntersect(collider.GetRect())
}

// -----------------------------------------------------------------------------
// Collider public functions
// -----------------------------------------------------------------------------

func (c *Collider) CollideWith(collider *Collider) bool {
	if c.GetRect() != nil {
		if collider.GetRect() != nil {
			return c.collideRectWithRect(collider)
		}
		if collider.GetPoints() != nil {
			return c.collideRectWithPoints(collider)
		}
	}
	if c.GetPoints() != nil {
		if collider.GetRect() != nil {
			return collider.collideRectWithRect(c)
		}
		if collider.GetPoints() != nil {
			return c.collidePointsWithPoints(collider)
		}
	}
	return false
}

func (c *Collider) GetPoints() []*api.Point {
	return c.points
}

func (c *Collider) GetRect() *api.Rect {
	return c.rect
}

func (c *Collider) SetPoints(points []*api.Point) {
	c.points = points
}

func (c *Collider) SetRect(rect *api.Rect) {
	c.rect = rect
}

// -----------------------------------------------------------------------------
// Package private functions
// -----------------------------------------------------------------------------

// checkCollisionBetweenEntities function checks if there are ny collision
// between two sets of entities.
// If inTile flag is set, then, coordinates for checking collisions are taken
// relative to the tile where entities belong.
func checkCollisionBetweenEntities(entOne IEntity, entTwo IEntity) bool {
	entOneCollider := entOne.GetCollider()
	entTwoCollider := entTwo.GetCollider()
	return entOneCollider.CollideWith(entTwoCollider)
}

// checkCollisionsWorker function is a go routine for checking collisions
// between two sets of entities.
// If inTile flag is set, then, coordinates for checking collisions are taken
// relative to the tile where entities belong.
func checkCollisionsWorker(collisions []IEntity, dynamics <-chan IEntity, results chan<- int) {
	for d := range dynamics {
		for _, coll := range collisions {
			if coll == d {
				continue
			}
			if checkCollisionBetweenEntities(d, coll) {
				//d.Collide(coll)
				//coll.Collide(d)
			}
		}
		results <- 1
	}
}

// -----------------------------------------------------------------------------
// Package private functions
// -----------------------------------------------------------------------------

func CheckCollisionWith(entity IEntity, entities []IEntity) {
	if (len(entities) == 0) || (entity == nil) {
		return
	}
	for _, ent := range entities {
		if checkCollisionBetweenEntities(entity, ent) {
			//d.Collide(coll)
			//coll.Collide(d)
		}
	}
}

// checkCollisions function checks collisions between two sets of scene entities.
func CheckCollisions(phyEntities []IEntity, dynEntities []IEntity) {
	lenDynamics := len(dynEntities)
	if lenDynamics == 0 {
		return
	}
	dynamics := make(chan IEntity, lenDynamics)
	results := make(chan int, lenDynamics)
	for w := 0; w <= lenDynamics/3; w++ {
		go checkCollisionsWorker(phyEntities, dynamics, results)
	}
	for _, dyn := range dynEntities {
		dynamics <- dyn
	}
	close(dynamics)
	for r := 0; r < lenDynamics; r++ {
		<-results
	}
}
