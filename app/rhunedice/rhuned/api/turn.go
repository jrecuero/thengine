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

func (t *TurnHandler) populateActiveEnemies() {
	t.activeEnemies = []IAvatar{}
	for _, enemy := range t.enemies {
		if active := enemy.IsActive(); active {
			t.activeEnemies = append(t.activeEnemies, enemy)
		}
	}
}

func (t *TurnHandler) AddEnemy(enemy IAvatar) {
	t.enemies = append(t.enemies, enemy)
}

func (t *TurnHandler) AddPlayer(player IAvatar) {
	t.player = player
}

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

func (t *TurnHandler) GetEnemies() []IAvatar {
	return t.enemies
}

func (t *TurnHandler) GetPlayer() IAvatar {
	return t.player
}

func (t *TurnHandler) GetState() ETurnState {
	return t.state
}

func (t *TurnHandler) Init() ETurnState {
	t.state = InitTurn

	tools.Logger.WithField("module", "turn-handler").
		WithField("method", "Init").
		Tracef("Initialize turn stage")

	return StartTurn
}

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

func (t *TurnHandler) Run(args ...any) {
}

// RunState method returns if the new stage should run (true) or it has to be
// blocked (false).
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

func (t *TurnHandler) SetState(state ETurnState) {
	t.state = state
}

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
