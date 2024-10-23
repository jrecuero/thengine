package api

import "github.com/jrecuero/thengine/pkg/tools"

type ETurnState string

const (
	NoneTurn                ETurnState = "none"
	InitTurn                ETurnState = "init"
	StartTurn               ETurnState = "start"
	AvatarUpdateTurn        ETurnState = "avatar-update"
	AvatarUpdateBucketTurn  ETurnState = "avatar-update-turn"
	RollDiceTurn            ETurnState = "roll-dice"
	BucketSelectionTurn     ETurnState = "bucket-selection"
	PlayerExecuteBucketTurn ETurnState = "player-execute-bucket"
	EnemyExecuteBucketTurn  ETurnState = "enemy-execute-bucket"
	EndTurn                 ETurnState = "end"
)

// TurnHandler struct contains all attributes and logic required for handling
// the application.
// It is mostly in charge of the state machine with this flow of states:
//
//   - InitTurn: initial state for the handler. Everything should be initialized
//     at this time, and this state should not be repeated.
//   - StartTurn: firt state for every turn, where everything should be setup to
//     handle a new turn.
//   - AvatarUpdateTurn: every avatar involved in the turn initializes all their
//     internal information for the turn, based on configuration or previous turns
//     effects.
//   - AvatarUpdateBucketTurn: every avatar updates their buckets with their
//     particular stats.
//   - RollDiceTurn: roll all avatars dice sets.
//   - BucketSelection: every avatar selects roll diceset buckets to be used in
//     the turn.
//   - PlayerExecuteBucketTurn: selected roll diceset are applied on avatar
//     buckets and executed.
//   - EnemyExecuteBucketTurn: selected roll diceset are applied on enemy
//     avatar.
//   - EndTurn: handler ends the turn and any avatar updates with all results.
type TurnHandler struct {
	player        IAvatar
	enemies       []IAvatar
	activeEnemies []IAvatar
	state         ETurnState
}

func NewTurnHandler(player IAvatar, enemies []IAvatar) *TurnHandler {
	return &TurnHandler{
		player:        player,
		enemies:       enemies,
		activeEnemies: nil,
		state:         NoneTurn,
	}
}

// populateActiveEnemies private method loads all available active enemies.
func (t *TurnHandler) populateActiveEnemies() {
	t.activeEnemies = []IAvatar{}
	for _, enemy := range t.enemies {
		if active := enemy.IsActive(); active {
			t.activeEnemies = append(t.activeEnemies, enemy)
		}
	}
}

// AddEnemy public method adds a new enemy to the list of available enemies.
func (t *TurnHandler) AddEnemy(enemy IAvatar) {
	t.enemies = append(t.enemies, enemy)
}

// AddPlayer public method adds the active player.
func (t *TurnHandler) AddPlayer(player IAvatar) {
	t.player = player
}

// AvatarUpdate public method updates player and all active enemies.
func (t *TurnHandler) AvatarUpdate() ETurnState {
	t.state = AvatarUpdateTurn

	tools.Logger.WithField("module", "turn-handler").
		WithField("method", "AvatarUpdate").
		Tracef("Avatar update stage")

	t.player.UpdateTurn()
	for _, enemy := range t.activeEnemies {
		enemy.UpdateTurn()
	}

	return AvatarUpdateBucketTurn
}

// AvatarUpdateBucket public method updates player and all active enemies
// buckets.
func (t *TurnHandler) AvatarUpdateBucket() ETurnState {
	t.state = AvatarUpdateBucketTurn

	tools.Logger.WithField("module", "turn-handler").
		WithField("method", "AvatarUpdateBucket").
		Tracef("Avatar update bucket stage")

	t.player.UpdateBucketTurn()
	for _, enemy := range t.activeEnemies {
		enemy.UpdateBucketTurn()
	}

	return RollDiceTurn
}

// BucketSelection public method allows the user to select up to two buckets
// that have been rolled down.
func (t *TurnHandler) BucketSelection() ETurnState {
	t.state = BucketSelectionTurn

	tools.Logger.WithField("module", "turn-handler").
		WithField("method", "BucketSelection").
		Tracef("Bucket selection stage")

	t.player.BucketSelectionTurn()
	for _, enemy := range t.activeEnemies {
		enemy.BucketSelectionTurn()
	}

	return PlayerExecuteBucketTurn
}

// End public method ends the turn, calling any player or active enemies end
// turn behavior.
func (t *TurnHandler) End() ETurnState {
	t.state = EndTurn

	tools.Logger.WithField("module", "turn-handler").
		WithField("method", "End").
		Tracef("End stage")

	t.player.EndTurn()
	for _, enemy := range t.activeEnemies {
		enemy.EndTurn()
	}

	return StartTurn
}

// EnemyExecuteBucket public method executes active enemies selected buckets.
func (t *TurnHandler) EnemyExecuteBucket(args ...any) ETurnState {
	t.state = EnemyExecuteBucketTurn

	tools.Logger.WithField("module", "turn-handler").
		WithField("method", "ExecuteBucket").
		Tracef("Execute bucket stage %v", args)

	for _, enemy := range t.activeEnemies {
		if nextBucket := enemy.NextSelected(); nextBucket != nil {
			enemy.ExecuteButcketTurn(nextBucket, t.player)
			enemy.RemoveNextSelected()
		}
	}

	// if at least one enemy had a bucket left to execute, don't move to the
	// next state.
	for _, enemy := range t.activeEnemies {
		if nextBucket := enemy.NextSelected(); nextBucket != nil {
			return t.state
		}
	}
	return EndTurn
}

// GetEnemies public method returns all available enemies.
func (t *TurnHandler) GetEnemies() []IAvatar {
	return t.enemies
}

// GetPlayer public method returns the player.
func (t *TurnHandler) GetPlayer() IAvatar {
	return t.player
}

// GetState public method returns the active turn state.
func (t *TurnHandler) GetState() ETurnState {
	return t.state
}

// Init public method initializes the turn. It should be called only one time
// at initialization time
func (t *TurnHandler) Init() ETurnState {
	t.state = InitTurn

	tools.Logger.WithField("module", "turn-handler").
		WithField("method", "Init").
		Tracef("Initialize turn stage")

	return StartTurn
}

// PlayerExecuteBucket public method executes player selected buckets.
func (t *TurnHandler) PlayerExecuteBucket(args ...any) ETurnState {
	t.state = PlayerExecuteBucketTurn

	tools.Logger.WithField("module", "turn-handler").
		WithField("method", "ExecuteBucket").
		Tracef("Execute bucket stage %v", args)

	if nextBucket := t.player.NextSelected(); nextBucket != nil {
		t.player.ExecuteButcketTurn(nextBucket, t.activeEnemies[0])
		t.player.RemoveNextSelected()
	}

	if nextBucket := t.player.NextSelected(); nextBucket != nil {
		return t.state
	}
	return EnemyExecuteBucketTurn
}

// RollDice public method rolls player and active enemies dice.
func (t *TurnHandler) RollDice() ETurnState {
	t.state = RollDiceTurn

	tools.Logger.WithField("module", "turn-handler").
		WithField("method", "RollDice").
		Tracef("roll-dice stage")

	t.player.RollDiceTurn()
	for _, enemy := range t.activeEnemies {
		enemy.RollDiceTurn()
	}

	return BucketSelectionTurn
}

// Run public method runs the turn.
func (t *TurnHandler) Run(args ...any) {
}

// RunState public method returns if the new stage should run (true) or it has
// to be blocked (false).
func (t *TurnHandler) RunState(args ...any) bool {
	var newState ETurnState = NoneTurn
	var run bool = true

	switch t.state {
	case InitTurn:
		newState = t.Init()
	case StartTurn:
		newState = t.Start()
	case AvatarUpdateTurn:
		newState = t.AvatarUpdate()
	case AvatarUpdateBucketTurn:
		newState = t.AvatarUpdateBucket()
		run = false
	case RollDiceTurn:
		// user input in order to run this stage
		newState = t.RollDice()
		run = false
	case BucketSelectionTurn:
		// user input/selection in order to run this stage
		newState = t.BucketSelection()
	case PlayerExecuteBucketTurn:
		// user input in order to run this stage
		newState = t.PlayerExecuteBucket(args)
	case EnemyExecuteBucketTurn:
		newState = t.EnemyExecuteBucket(args)
	case EndTurn:
		newState = t.End()
		run = false
	}

	t.state = newState
	return run
}

// SetState public method sets the active turn state.
func (t *TurnHandler) SetState(state ETurnState) {
	t.state = state
}

// Start public method starts the turn.
func (t *TurnHandler) Start() ETurnState {
	t.state = StartTurn

	tools.Logger.WithField("module", "turn-handler").
		WithField("method", "Start").
		Tracef("Start stage")

	t.populateActiveEnemies()

	t.player.StartTurn()
	for _, enemy := range t.activeEnemies {
		enemy.StartTurn()
	}

	return AvatarUpdateTurn
}
