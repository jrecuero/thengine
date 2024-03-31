// collider.go contains all logic required for handling collision. Collisions
// occur between entities that are physical and dynamic. Any physical entity
// can collide with any other. Dynamic entities are those that can move across
// the scenario.
package engine

// -----------------------------------------------------------------------------
// Package private functions
// -----------------------------------------------------------------------------

// checkCollisionBetweenEntities function checks if there are ny collision
// between two sets of entities.
// If inTile flag is set, then, coordinates for checking collisions are taken
// relative to the tile where entities belong.
func checkCollisionBetweenEntities(entOne IEntity, entTwo IEntity) bool {
	entOneX, entOneY := entOne.GetPosition().Get()
	entTwoX, entTwoY := entTwo.GetPosition().Get()
	entOneW, entOneH := entOne.GetSize().Get()
	entTwoW, entTwoH := entTwo.GetSize().Get()
	if entOneX < entTwoX+entTwoW && entOneX+entOneW > entTwoX &&
		entOneY < entTwoY+entTwoH && entOneY+entOneH > entTwoY {
		return true
	}
	return false
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
