package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

const (
	// Widget names
	SnakeWidgetName            = "widget/snake/1"
	TextPointsWidgetName       = "text/points/1"
	TextHighScoreWidgetName    = "text/high-score/1"
	BoxWidgetName              = "widget/box/1"
	TimerFoodWidgetName        = "timer/food/1"
	TimerFoodPieceWidgetName   = "timer/food-piece/%d"
	TextGameOverWidgetName     = "text/game-over/1"
	BulletWidgetName           = "widget/bullet/%d"
	TimerGaugeBullerWidgetName = "timer-gauge/bullet/1"

	// Scene names
	MainSceneName     = "scene/main/1"
	GameOverSceneName = "scene/game-over/1"

	// Handler names
	AppHanderName = "/handler/app/1"

	SnakePointsTopic = "snake/points/topic"
	GameOverTopic    = "snake/game-over/topic"
)

var (
	FoodPieceCounter int = 0
	BulletCounter    int = 0
)

type Bullet struct {
	*widgets.Widget
	direction string
	x         float64
	y         float64
	parent    *snake
}

func NewBullet(position *api.Point, style *tcell.Style, direction string, parent *snake) *Bullet {
	name := fmt.Sprintf(BulletWidgetName, BulletCounter)
	bullet := &Bullet{
		Widget:    widgets.NewWidget(name, position, api.NewSize(1, 1), style),
		direction: direction,
		x:         float64(position.X),
		y:         float64(position.Y),
		parent:    parent,
	}
	bullet.GetCanvas().WriteStringInCanvas(string(tcell.RuneBullet), style)
	bullet.SetSolid(true)
	return bullet
}

func (b *Bullet) PublishCollision(points int) {
	message := &engine.Message{
		Topic:   SnakePointsTopic,
		Src:     b.GetName(),
		Dst:     "broadcast",
		Content: points,
	}
	engine.GetMailbox().Publish(SnakePointsTopic, message)
}

func (b *Bullet) Update(ecent tcell.Event, scene engine.IScene) {
	var vx, vy int
	switch b.direction {
	case "up":
		vx = 0
		vy = -1
	case "down":
		vx = 0
		vy = 1
	case "left":
		vx = -1
		vy = 0
	case "right":
		vx = 1
		vy = 0
	}
	speed := 2.0
	b.x += float64(vx) / speed
	b.y += float64(vy) / speed
	intX := int(b.x)
	intY := int(b.y)
	if (intX != b.GetPosition().X) || (intY != b.GetPosition().Y) {
		if (intX >= 80) || (intX < 0) || (intY >= 20) || (intY < 0) {
			b.parent.Bullet = nil
			scene.RemoveEntity(b)
			return
		}
		b.SetPosition(api.NewPoint(intX, intY))
	}

	collisions := scene.CheckCollisionWith(b)
	for _, ent := range collisions {
		if _, ok := ent.(*TimerFoodPiece); ok {
			tools.Logger.WithField("module", "snake").
				WithField("struct", "Bullet").
				WithField("function", "Update").
				Debugf("bullet collider %s", b.GetCollider().GetRect().ToString())
			tools.Logger.WithField("module", "snake").
				WithField("struct", "Bullet").
				WithField("function", "Update").
				Debugf("collision collider %s", ent.GetCollider().GetRect().ToString())
			b.parent.Bullet = nil
			scene.RemoveEntity(ent)
			scene.RemoveEntity(b)
			b.PublishCollision(ent.(*TimerFoodPiece).Points)
			return
		}
	}
}

type snake struct {
	*widgets.Sprite
	x         float64
	y         float64
	speed     float64
	direction string
	alive     bool
	Bullet    *Bullet
}

func NewSnake(position *api.Point, style *tcell.Style) *snake {
	cell := engine.NewCell(style, '#')
	snake := &snake{
		Sprite: widgets.NewSprite(SnakeWidgetName, position,
			[]*widgets.SpriteCell{widgets.NewSpriteCell(api.NewPoint(5, 5), cell)}),
		speed:     10.0,
		direction: "none",
		alive:     true,
		Bullet:    nil,
	}
	snake.x = 5.0
	snake.y = 5.0
	snake.SetFocusType(engine.SingleFocus)
	snake.SetFocusEnable(true)
	return snake
}

func (s *snake) Move(args ...any) {
	direction := args[0].(string)
	// Don't allow to reverse vertical direction.
	if (direction == "up") && (s.direction == "down") ||
		(direction == "down") && (s.direction == "up") {
		return
	}
	// Don't allow to reverse horizontal direction.
	if (direction == "left") && (s.direction == "right") ||
		(direction == "right") && (s.direction == "left") {
		return
	}
	s.direction = args[0].(string)
}

func (s *snake) Shoot(args ...any) {
	tools.Logger.WithField("module", "snake").
		WithField("struct", "snake").
		WithField("function", "Shoot").
		Debugf("snake is shooting %s", s.direction)
	if s.Bullet == nil {
		scene := args[0].(engine.IScene)
		BulletCounter++
		spriteCell := s.GetSpriteCells()[0]
		position := api.ClonePoint(s.GetPosition())
		position.Add(spriteCell.GetPosition())
		s.Bullet = NewBullet(position, s.GetStyle(), s.direction, s)
		scene.AddEntity(s.Bullet)
	}
}

func (s *snake) PublishCollision(points int) {
	message := &engine.Message{
		Topic:   SnakePointsTopic,
		Src:     s.GetName(),
		Dst:     "broadcast",
		Content: points,
	}
	engine.GetMailbox().Publish(SnakePointsTopic, message)
}

func (s *snake) PublishGameOver() {
	message := &engine.Message{
		Topic:   GameOverTopic,
		Src:     s.GetName(),
		Dst:     AppHanderName,
		Content: "game-over",
	}
	engine.GetMailbox().Publish(GameOverTopic, message)
}

func (s *snake) Update(event tcell.Event, scene engine.IScene) {
	if !s.HasFocus() {
		return
	}
	if !s.alive {
		return
	}
	actions := []*widgets.KeyboardAction{
		{Key: tcell.KeyUp, Callback: s.Move, Args: []any{"up"}},
		{Key: tcell.KeyDown, Callback: s.Move, Args: []any{"down"}},
		{Key: tcell.KeyRight, Callback: s.Move, Args: []any{"right"}},
		{Key: tcell.KeyLeft, Callback: s.Move, Args: []any{"left"}},
		{Rune: ' ', Callback: s.Shoot, Args: []any{scene}},
	}
	s.HandleKeyboardForActions(event, actions)
	var vx, vy int
	if s.direction != "none" {
		switch s.direction {
		case "up":
			vx = 0
			vy = -1
		case "down":
			vx = 0
			vy = 1
		case "left":
			vx = -1
			vy = 0
		case "right":
			vx = 1
			vy = 0
		}
		s.x += float64(vx) / s.speed
		s.y += float64(vy) / s.speed
		spriteCell := s.GetSpriteCells()[0]
		intX := int(s.x)
		intY := int(s.y)
		if (intX != spriteCell.GetPosition().X) || (intY != spriteCell.GetPosition().Y) {
			if (intX >= 80) || (intX < 0) || (intY >= 20) || (intY < 0) {
				s.alive = false
				tools.Logger.WithField("module", "snake").
					WithField("struct", "snake").
					WithField("function", "Update").
					Debugf("snake is dead at %d,%d", intX, intY)
				s.PublishGameOver()
				return
			}
			// Remove the last entry and place first with position increased
			// based in the direction the snake is running.
			lastCell := s.RemoveSpriteCellAt(widgets.AtTheEnd)
			lastCell.SetPosition(api.NewPoint(intX, intY))
			s.AddSpriteCellAt(0, lastCell)
			//tools.Logger.WithField("module", "snake").
			//    WithField("struct", "snake").
			//    WithField("function", "Update").
			//    Debugf("rotate %s", lastCell.GetPosition().ToString())
		}
		// check collision with itself.
		spriteCells := s.GetSpriteCells()
		spriteCell = spriteCells[0]
		for i := 1; i < len(s.GetSpriteCells()); i++ {
			if spriteCell.GetPosition().IsEqual(spriteCells[i].GetPosition()) {
				s.alive = false
				tools.Logger.WithField("module", "snake").
					WithField("struct", "snake").
					WithField("function", "Update").
					Debugf("snake is dead at %d", i)
				s.PublishGameOver()
				return
			}

		}

		collisions := scene.CheckCollisionWith(s)
		if len(collisions) != 0 {
			for _, ent := range collisions {
				if _, ok := ent.(*TimerFoodPiece); ok {
					spriteCell = s.GetSpriteCells()[0]
					scene.RemoveEntity(ent)
					// Update float64 position with new entry.
					s.x = float64(spriteCell.GetPosition().X + vx)
					s.y = float64(spriteCell.GetPosition().Y + vy)
					newSpriteCell := widgets.NewSpriteCell(api.NewPoint(int(s.x), int(s.y)), spriteCell.GetCell())
					s.AddSpriteCellAt(0, newSpriteCell)
					tools.Logger.WithField("module", "snake").
						WithField("struct", "snake").
						WithField("function", "Update").
						Debugf("collision with %s at %s", ent.GetName(), newSpriteCell.GetPosition().ToString())
					s.PublishCollision(ent.(*TimerFoodPiece).Points)
				}
			}
		}
	}
}

type TextPoints struct {
	*widgets.Text
	Points int
}

func (t *TextPoints) Consume() {
	if message, _ := engine.GetMailbox().Consume(SnakePointsTopic, t.GetName()); message != nil {
		t.Points += message.Content.(int)
		t.SetText(fmt.Sprintf("Points: %d", t.Points))
	}
}

type TextGameOver struct {
	*widgets.Text
}

func (t *TextGameOver) Exit(args ...any) {
	tools.Logger.WithField("module", "snake").
		WithField("struct", "TextGameOver").
		WithField("function", "Exit").
		Debugf("key pressed")
	message := &engine.Message{
		Topic:   GameOverTopic,
		Src:     t.GetName(),
		Dst:     AppHanderName,
		Content: "restart",
	}
	engine.GetMailbox().Publish(GameOverTopic, message)
}

func (t *TextGameOver) Update(event tcell.Event, scene engine.IScene) {
	if !t.HasFocus() {
		return
	}
	actions := []*widgets.KeyboardAction{
		{Key: tcell.KeyEnter, Callback: t.Exit, Args: nil},
	}
	t.HandleKeyboardForActions(event, actions)
}

type TimerFoodPiece struct {
	*widgets.Widget
	interval time.Duration
	time     time.Time
	Points   int
}

func NewTimerFoodPiece(name string, position *api.Point, style *tcell.Style, duration time.Duration, points int) *TimerFoodPiece {
	return &TimerFoodPiece{
		Widget:   widgets.NewWidget(name, position, api.NewSize(1, 1), style),
		interval: duration,
		Points:   points,
	}
}

func (t *TimerFoodPiece) Start() {
	t.time = time.Now()
}

func (t *TimerFoodPiece) Update(event tcell.Event, scene engine.IScene) {
	now := time.Now()
	if elapsed := now.Sub(t.time); elapsed < t.interval {
		return
	}
	scene.RemoveEntity(t)
}

type AppHandler struct {
	*engine.Entity
	Running   bool
	HighScore int
}

func (h *AppHandler) Consume() {
	if message, _ := engine.GetMailbox().Consume(GameOverTopic, h.GetName()); message != nil {
		tools.Logger.WithField("module", "snake").
			WithField("struct", "AppHandler").
			WithField("function", "Consume").
			Debugf("cosume message")

		content, ok := message.Content.(string)
		if !ok {
			return
		}
		sceneManager := engine.GetEngine().GetSceneManager()
		mainScene := sceneManager.GetSceneByName(MainSceneName)
		gameOverScene := sceneManager.GetSceneByName(GameOverSceneName)
		textPoints := mainScene.GetEntityByName(TextPointsWidgetName).(*TextPoints)
		if content == "game-over" {
			textGameOver := gameOverScene.GetEntityByName(TextGameOverWidgetName).(*TextGameOver)
			textGameOver.SetText(fmt.Sprintf("GAME OVER\nScore: %d\nHSCORE: %d", textPoints.Points, h.HighScore))
			sceneManager.DeactivateScene(mainScene)
			sceneManager.ActivateScene(gameOverScene)
		} else if content == "restart" {
			if textPoints.Points > h.HighScore {
				h.HighScore = textPoints.Points
			}
			sceneManager.DeactivateScene(gameOverScene)
			// Remove all entities from the scene.
			mainScene.Clean()
			// Remove all entities from the scene in the focus manager.
			engine.GetEngine().GetFocusManager().RemoveEntitiesInScene(mainScene)
			// Run scene setup and start.
			h.SetUpMainScene(mainScene)
			mainScene.Start()
			sceneManager.ActivateScene(mainScene)
		}
	}
}

func (h *AppHandler) SetUpMainScene(mainScene engine.IScene) {
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	styleTwo := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	styleThree := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorBlack)
	styleFour := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlue)
	styleFive := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorYellow)

	textPoints := &TextPoints{
		Text: widgets.NewText(TextPointsWidgetName, api.NewPoint(0, 0), api.NewSize(20, 1), &styleThree, "Points: 0"),
	}
	mainScene.AddEntity(textPoints)

	highScore := fmt.Sprintf("HSCORE: %d", h.HighScore)
	TextHighScore := widgets.NewText(TextHighScoreWidgetName, api.NewPoint(65, 0), api.NewSize(20, 1), &styleThree, highScore)
	mainScene.AddEntity(TextHighScore)

	gauge := widgets.NewTimerGauge(TimerGaugeBullerWidgetName, api.NewPoint(30, 0), api.NewSize(10, 1), &styleOne, 2*time.Second, 10)
	mainScene.AddEntity(gauge)

	box := engine.NewEntity(BoxWidgetName, api.NewPoint(0, 1), api.NewSize(80, 20), &styleTwo)
	box.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &styleTwo, engine.CanvasRectSingleLine)
	mainScene.AddEntity(box)

	foodTimer := widgets.NewTimer(TimerFoodWidgetName, 5*time.Second, widgets.ForeverTimer)
	foodTimer.SetWidgetCallback(func(entity engine.IEntity, args ...any) bool {
		x := rand.Intn(78) + 1
		y := rand.Intn(18) + 2
		FoodPieceCounter++
		foodPieceName := fmt.Sprintf(TimerFoodPieceWidgetName, FoodPieceCounter)
		duration := rand.Intn(30) + 10
		var points int
		var style tcell.Style
		if duration < 20 {
			points = duration + 20
			style = styleOne
		} else if duration < 30 {
			points = duration - 10
			style = styleFour
		} else if duration < 40 {
			points = duration - 30
			style = styleFive
		}
		timeDuration := time.Duration(duration) * time.Second
		food := NewTimerFoodPiece(foodPieceName, api.NewPoint(x, y), &style, timeDuration, points)
		food.GetCanvas().WriteStringInCanvas("*", &style)
		food.SetSolid(true)
		mainScene.AddEntity(food)
		tools.Logger.WithField("module", "snake").
			WithField("struct", "AppHandler").
			WithField("function", "SetUpMainScene").
			Debugf("foodTimer generates new food %d %d", duration, points)
		return true
	})
	mainScene.AddEntity(foodTimer)

	snake := NewSnake(api.NewPoint(0, 1), &styleTwo)
	mainScene.AddEntity(snake)
}

func (h *AppHandler) SetUp(appEngine *engine.Engine) {
	appEngine.InitResources()
	camera := engine.NewCamera(api.NewPoint(0, 0), api.NewSize(90, 30))
	//styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	//styleTwo := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	//styleThree := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorBlack)
	styleFour := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlue)
	//styleFive := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorRed)
	mainScene := engine.NewScene(MainSceneName, camera)

	h.SetUpMainScene(mainScene)

	// TODO this is required in the initialized part
	engine.GetMailbox().CreateTopic(SnakePointsTopic)
	engine.GetMailbox().Subscribe(SnakePointsTopic, TextPointsWidgetName)

	engineScene, err := appEngine.CreateEngineScene()
	if err != nil {
		panic(err)
	}
	engineScene.AddEntity(theAppHandler)
	appEngine.GetSceneManager().SetSceneAsActive(engineScene)

	appEngine.GetSceneManager().AddScene(mainScene)
	appEngine.GetSceneManager().ActivateScene(mainScene)

	gameOverScene := engine.NewScene(GameOverSceneName, camera)
	gameOverText := &TextGameOver{
		Text: widgets.NewText(TextGameOverWidgetName, api.NewPoint(40, 10), api.NewSize(20, 1), &styleFour, "GAME OVER"),
	}
	gameOverText.SetFocusType(engine.SingleFocus)
	gameOverText.SetFocusEnable(true)
	gameOverScene.AddEntity(gameOverText)
	appEngine.GetSceneManager().AddScene(gameOverScene)

	engine.GetMailbox().CreateTopic(GameOverTopic)
	engine.GetMailbox().Subscribe(GameOverTopic, AppHanderName)

	appEngine.Init()
	appEngine.Start()
	theAppHandler.Running = true
}

var (
	theAppHandler *AppHandler = &AppHandler{
		Entity:  engine.NewNamedEntity(AppHanderName),
		Running: false,
	}
)

func demoSnake(dryRun bool) {
	tools.Logger.WithField("module", "main").
		WithField("dry-mode", dryRun).
		Debugf("ThEngine demo-snake")
	fmt.Println("ThEngine demo-snake")
	appEngine := engine.GetEngine()

	theAppHandler.SetUp(appEngine)

	appEngine.Run(60.0)
}
