// statuseffect.go package contains all status effects constants.
package constants

const (
	// IsForever is the duration for any permanent status.
	IsForever = -1

	// -----------------------------------------------------------------------------
	// Blinded: A blinded creature can’t see and automatically fails any
	// ability check that requires sight.
	Blinded = "blinded"

	// Charmed: A charmed creature can’t attack the charmer or target the
	// charmer with harmful abilities or magical effects.
	Charmed = "charmed"

	// Deafened: A deafened creature can’t hear and automatically fails any
	// ability check that requires hearing.
	Deafened = "deafened"

	// Frightened: A frightened creature has disadvantage on ability checks and
	// attack rolls while the source of its fear is within line of sight.
	Frightened = "frightened"

	// Grappled: A grappled creature’s speed becomes 0, and it can’t benefit
	// from any bonus to its speed.
	Grappled = "grappled"

	// Incapacitated: An incapacitated creature can’t take actions or reactions.
	Incapacitated = "incapacitated"

	// Invisible: An invisible creature is impossible to see without the aid of
	// magic or a special sense.
	Invisible = "invisible"

	// Paralyzed: A paralyzed creature is incapacitated and can’t move or speak.
	Paralyzed = "paralized"

	// Petrified: A petrified creature is transformed, along with any nonmagical
	// object it is wearing or carrying, into a solid inanimate substance (usually
	// stone).
	Petrified = "petrified"

	// Poisoned: A poisoned creature has disadvantage on attack rolls and ability
	// checks.
	Poisoned = "poisoned"

	// Prone: A prone creature’s only movement option is to crawl, unless it
	// stands up and thereby ends the condition.
	Prone = "prone"

	// Restrained: A restrained creature’s speed becomes 0, and it can’t
	// benefit from any bonus to its speed.
	Restrained = "restrained"

	// Stunned: A stunned creature is incapacitated, can’t move, and can speak
	// only falteringly.
	Stunned = "stunned"

	// Unconscious: An unconscious creature is incapacitated, can’t move or
	// speak, and is unaware of its surroundings.
	Unconscious = "unconscious"
)
