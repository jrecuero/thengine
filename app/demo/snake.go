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
	SnakeWidgetName      = "widget/snake/1"
	TextPointsWidgetName = "text/points/1"
	BoxWidgetName        = "widget/box/1"
	FoodTimerWidgetName  = "timer/food/1"
	TextGameOverName     = "text/game-over/1"

	// Scene names
	MainSceneName     = "scene/main/1"
	GameOverSceneName = "scene/game-over/1"

	// Handler names
	AppHanderName = "/handler/app/1"

	SnakePointsTopic = "snake/points/topic"
	GameOverTopic    = "snake/game-over/topic"
)

type snake struct {
	*widgets.Sprite
	x         float64
	y         float64
	speed     float64
	direction string
	alive     bool
}

func NewSnake(position *api.Point, style *tcell.Style) *snake {
	cell := engine.NewCell(style, '#')
	snake := &snake{
		Sprite: widgets.NewSprite(SnakeWidgetName, position,
			[]*widgets.SpriteCell{widgets.NewSpriteCell(api.NewPoint(5, 5), cell)}),
		speed:     10.0,
		direction: "none",
		alive:     true,
	}
	//snake.x = float64(snake.GetPosition().X)
	//snake.y = float64(snake.GetPosition().Y)
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

func (s *snake) PublishCollision() {
	message := &engine.Message{
		Topic:   SnakePointsTopic,
		Src:     s.GetName(),
		Dst:     "broadcast",
		Content: nil,
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
					Debugf("eat %s", newSpriteCell.GetPosition().ToString())
				s.PublishCollision()
			}
		}
	}
}

type TextPoints struct {
	*widgets.Text
	points int
}

func (t *TextPoints) Consume() {
	if message, _ := engine.GetMailbox().Consume(SnakePointsTopic, t.GetName()); message != nil {
		t.points += 10
		t.SetText(fmt.Sprintf("Points: %d", t.points))
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

type AppHandler struct {
	*engine.Entity
	Running bool
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
		if content == "game-over" {
			sceneManager.DeactivateScene(mainScene)
			sceneManager.ActivateScene(gameOverScene)
		} else if content == "restart" {
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

	textPoints := &TextPoints{
		Text: widgets.NewText(TextPointsWidgetName, api.NewPoint(0, 0), api.NewSize(20, 1), &styleThree, "Points: 0"),
	}
	mainScene.AddEntity(textPoints)

	box := engine.NewEntity(BoxWidgetName, api.NewPoint(0, 1), api.NewSize(80, 20), &styleTwo)
	box.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &styleTwo, engine.CanvasRectSingleLine)
	mainScene.AddEntity(box)

	foodTimer := widgets.NewTimer(FoodTimerWidgetName, 5*time.Second, widgets.ForeverTimer)
	foodTimer.SetWidgetCallback(func(entity engine.IEntity, args ...any) bool {
		x := rand.Intn(78) + 1
		y := rand.Intn(18) + 2
		food := engine.NewEntity("food/1", api.NewPoint(x, y), api.NewSize(1, 1), &styleOne)
		food.GetCanvas().WriteStringInCanvas(".", &styleOne)
		food.SetSolid(true)
		mainScene.AddEntity(food)
		tools.Logger.WithField("module", "snake").
			WithField("struct", "AppHandler").
			WithField("function", "SetUpMainScene").
			Debugf("foodTimer generates new food")
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

	////textPoints := &TextPoints{
	////    Text: widgets.NewText(TextPointsWidgetName, api.NewPoint(0, 0), api.NewSize(20, 1), &styleThree, "Points: 0"),
	////}
	//textPoints := &TextPoints{
	//    Text: widgets.NewText(TextPointsWidgetName, nil, nil, &styleThree, ""),
	//}
	//// These should be the values the instance has to take every time it is
	//// initialized.
	//textPoints.SetCustomStart(func() {
	//    textPoints.SetPosition(api.NewPoint(0, 0))
	//    textPoints.SetSize(api.NewSize(20, 1))
	//    textPoints.SetText("Points: 0")
	//})

	//mainScene.AddEntity(textPoints)
	// TODO this is required in the initialized part
	engine.GetMailbox().CreateTopic(SnakePointsTopic)
	engine.GetMailbox().Subscribe(SnakePointsTopic, TextPointsWidgetName)

	//box := engine.NewEntity(BoxWidgetName, api.NewPoint(0, 1), api.NewSize(80, 20), &styleTwo)
	//box.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &styleTwo, engine.CanvasRectSingleLine)
	//mainScene.AddEntity(box)

	//foodTimer := widgets.NewTimer(FoodTimerWidgetName, 5*time.Second, widgets.ForeverTimer)
	//foodTimer.SetWidgetCallback(func(entity engine.IEntity, args ...any) bool {
	//    x := rand.Intn(78) + 1
	//    y := rand.Intn(18) + 2
	//    food := engine.NewEntity("food/1", api.NewPoint(x, y), api.NewSize(1, 1), &styleOne)
	//    food.GetCanvas().WriteStringInCanvas(".", &styleOne)
	//    food.SetSolid(true)
	//    mainScene.AddEntity(food)
	//    return true
	//})
	//mainScene.AddEntity(foodTimer)

	//snake := NewSnake(api.NewPoint(0, 1), &styleTwo)
	//mainScene.AddEntity(snake)

	engineScene, err := appEngine.CreateEngineScene()
	if err != nil {
		panic(err)
	}
	engineScene.AddEntity(theAppHandler)
	appEngine.GetSceneManager().SetSceneAsActive(engineScene)

	appEngine.GetSceneManager().AddScene(mainScene)
	appEngine.GetSceneManager().ActivateScene(mainScene)

	gameOverScene := engine.NewScene(GameOverSceneName, camera)
	//gameOverText := widgets.NewText(TextGameOverName, api.NewPoint(40, 10), api.NewSize(20, 1), &styleFour, "GAME OVER")
	gameOverText := &TextGameOver{Text: widgets.NewText(TextGameOverName, api.NewPoint(40, 10), api.NewSize(20, 1), &styleFour, "GAME OVER")}
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
