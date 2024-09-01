package api

type ETurnState string

const (
	NoneTurn               ETurnState = "none"
	InitTurn               ETurnState = "init"
	StartTurn              ETurnState = "start"
	AvatarUpdateTurn       ETurnState = "avatar-update"
	AvatarUpdateBucketTurn ETurnState = "avatar-update-turn"
	RollDiceTurn           ETurnState = "roll-dice"
	BucketSelectionTurn    ETurnState = "bucket-selection"
	ExecuteBucketTurn      ETurnState = "execute-bucket"
	EndTurn                ETurnState = "end"
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
//   - ExecuteBucketTurn: selected roll diceset are applied on avatar buckets
//     and executed.
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

	t.player.UpdateTurn()
	for _, enemy := range t.activeEnemies {
		enemy.UpdateTurn()
	}

	return AvatarUpdateBucketTurn
}

func (t *TurnHandler) AvatarUpdateBucket() ETurnState {
	t.state = AvatarUpdateBucketTurn

	t.player.UpdateBucketTurn()
	for _, enemy := range t.activeEnemies {
		enemy.UpdateBucketTurn()
	}

	return RollDiceTurn
}

func (t *TurnHandler) BucketSelection() ETurnState {
	t.state = BucketSelectionTurn

	t.player.SelectBucketTurn()
	for _, enemy := range t.activeEnemies {
		enemy.SelectBucketTurn()
	}

	return ExecuteBucketTurn
}

func (t *TurnHandler) End() ETurnState {
	t.state = EndTurn

	t.player.EndTurn()
	for _, enemy := range t.activeEnemies {
		enemy.EndTurn()
	}

	return InitTurn
}

func (t *TurnHandler) ExecuteBucket() ETurnState {
	t.state = ExecuteBucketTurn

	t.player.ExecuteButcketTurn()
	for _, enemy := range t.activeEnemies {
		enemy.ExecuteButcketTurn()
	}

	return EndTurn
}

func (t *TurnHandler) Init() ETurnState {
	t.state = InitTurn

	return StartTurn
}

func (t *TurnHandler) RollDice() ETurnState {
	t.state = RollDiceTurn

	t.player.RollDiceTurn()
	for _, enemy := range t.activeEnemies {
		enemy.RollDiceTurn()
	}

	return BucketSelectionTurn
}

func (t *TurnHandler) Run() {
	var newState ETurnState = NoneTurn

	switch t.state {
	case InitTurn:
		newState = t.Init()
	case StartTurn:
		newState = t.Start()
	case AvatarUpdateTurn:
		newState = t.AvatarUpdate()
	case AvatarUpdateBucketTurn:
		newState = t.AvatarUpdateBucket()
	case RollDiceTurn:
		newState = t.RollDice()
	case BucketSelectionTurn:
		newState = t.BucketSelection()
	case ExecuteBucketTurn:
		newState = t.ExecuteBucket()
	case EndTurn:
		newState = t.End()
	}

	t.state = newState
}

func (t *TurnHandler) Start() ETurnState {
	t.state = StartTurn

	t.populateActiveEnemies()

	t.player.StartTurn()
	for _, enemy := range t.activeEnemies {
		enemy.StartTurn()
	}

	return AvatarUpdateTurn
}
