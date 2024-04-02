package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type Player struct {
	*widgets.Widget
	origin  *api.Point
	TileMap *widgets.TileMap
}

func NewPlayer(name string, position *api.Point, size *api.Size, style *tcell.Style) *Player {
	player := &Player{
		Widget: widgets.NewWidget(name, position, size, style),
		origin: position,
	}
	player.SetFocusType(engine.SingleFocus)
	player.SetFocusEnable(true)
	return player

}

func (p *Player) Reset(args ...any) {
	p.SetPosition(p.origin)
}

func (p *Player) Move(args ...any) {
	steps := args[0].(int)
	direction := args[1].(string)
	tools.Logger.WithField("module", "main").
		WithField("struct", "Player").
		WithField("function", "Move").
		Debugf("steps %d direction %s", steps, direction)
	x, y := p.GetPosition().Get()
	if direction == "up" {
		p.SetPosition(api.NewPoint(x, y-1))
		if p.TileMap != nil {
			offsetX, offsetY := p.TileMap.GetCameraOffset().Get()
			p.TileMap.SetCameraOffset(api.NewPoint(offsetX, offsetY-1))
		}
	} else if direction == "down" {
		p.SetPosition(api.NewPoint(x, y+1))
		if p.TileMap != nil {
			offsetX, offsetY := p.TileMap.GetCameraOffset().Get()
			p.TileMap.SetCameraOffset(api.NewPoint(offsetX, offsetY+1))
		}
	} else if direction == "right" {
		p.SetPosition(api.NewPoint(x+1, y))
		if p.TileMap != nil {
			offsetX, offsetY := p.TileMap.GetCameraOffset().Get()
			p.TileMap.SetCameraOffset(api.NewPoint(offsetX+1, offsetY))
		}
	} else if direction == "left" {
		p.SetPosition(api.NewPoint(x-1, y))
		if p.TileMap != nil {
			offsetX, offsetY := p.TileMap.GetCameraOffset().Get()
			p.TileMap.SetCameraOffset(api.NewPoint(offsetX-1, offsetY))
		}
	}
}

func (p *Player) Update(event tcell.Event) {
	if !p.HasFocus() {
		return
	}
	actions := []*widgets.KeyboardAction{
		{
			Key:      tcell.KeyUp,
			Callback: p.Move,
			Args:     []any{1, "up"},
		},
		{
			Key:      tcell.KeyDown,
			Callback: p.Move,
			Args:     []any{1, "down"},
		},
		{
			Key:      tcell.KeyRight,
			Callback: p.Move,
			Args:     []any{1, "right"},
		},
		{
			Key:      tcell.KeyLeft,
			Callback: p.Move,
			Args:     []any{1, "left"},
		},
		{
			Rune:     'x',
			Callback: p.Reset,
			Args:     nil,
		},
	}
	p.HandleKeyboardForActions(event, actions)
}

const (
	DemoNineMoveTopic = "demo/nine/topic"
)

type ninePlayer struct {
	*widgets.Widget
	origin  *api.Point
	mailbox *engine.Mailbox
}

func newNinePlayer(name string, position *api.Point, size *api.Size, style *tcell.Style) *ninePlayer {
	player := &ninePlayer{
		Widget:  widgets.NewWidget(name, position, size, style),
		origin:  position,
		mailbox: engine.GetMailbox(),
	}
	player.mailbox.Subscribe(DemoNineMoveTopic, name)
	return player
}

func (p *ninePlayer) Reset(args ...any) {
	p.SetPosition(p.origin)
}

func (p *ninePlayer) Move(args ...any) {
	steps := args[0].(int)
	direction := args[1].(string)
	tools.Logger.WithField("module", "main").
		WithField("struct", "ninePlayer").
		WithField("function", "Move").
		Debugf("steps %d direction %s", steps, direction)
	x, y := p.GetPosition().Get()
	if direction == "up" {
		p.SetPosition(api.NewPoint(x, y-steps))
	} else if direction == "down" {
		p.SetPosition(api.NewPoint(x, y+steps))
	} else if direction == "right" {
		p.SetPosition(api.NewPoint(x+steps, y))
	} else if direction == "left" {
		p.SetPosition(api.NewPoint(x-steps, y))
	}
}

func (p *ninePlayer) Update(event tcell.Event) {
}

func (p *ninePlayer) Consume() {
	if message, _ := p.mailbox.Consume(DemoNineMoveTopic, p.GetName()); message != nil {
		tools.Logger.WithField("module", "main").
			WithField("struct", "ninePlayer").
			WithField("function", "Consume").
			Debugf("message %+v", message)
		content := message.Content.([]any)
		p.Move(content...)
	}
}

type nineHandler struct {
	*widgets.Widget
	mailbox *engine.Mailbox
}

func newNineHandler() *nineHandler {
	h := &nineHandler{
		Widget:  widgets.NewWidget("nine-handler", nil, nil, nil),
		mailbox: engine.GetMailbox(),
	}
	h.SetFocusType(engine.SingleFocus)
	h.SetFocusEnable(true)
	h.mailbox.CreateTopic(DemoNineMoveTopic)
	return h
}

func (h *nineHandler) Move(args ...any) {
	tools.Logger.WithField("module", "main").
		WithField("struct", "nineHandler").
		WithField("function", "Update").
		Debugf("args %+v", args)
	message := &engine.Message{
		Topic:   DemoNineMoveTopic,
		Src:     h.GetName(),
		Dst:     "broadcast",
		Content: args,
	}
	h.mailbox.Publish(DemoNineMoveTopic, message)
}

func (h *nineHandler) Update(event tcell.Event) {
	actions := []*widgets.KeyboardAction{
		{
			Key:      tcell.KeyUp,
			Callback: h.Move,
			Args:     []any{1, "up"},
		},
		{
			Key:      tcell.KeyDown,
			Callback: h.Move,
			Args:     []any{1, "down"},
		},
		{
			Key:      tcell.KeyRight,
			Callback: h.Move,
			Args:     []any{1, "right"},
		},
		{
			Key:      tcell.KeyLeft,
			Callback: h.Move,
			Args:     []any{1, "left"},
		},
	}
	h.HandleKeyboardForActions(event, actions)
}

func demoOne() {
	fmt.Println("ThEngine demo-one")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	defaultStyle := tcell.StyleDefault
	text := engine.NewCanvasFromString("Hello World", &defaultStyle)
	text.Render(camera)
	appEngine := engine.GetEngine()
	appEngine.Init()
	camera.Draw(true, appEngine.GetScreen())
	appEngine.Run(60.0)
}

func demoTwo() {
	fmt.Println("ThEngine demo-two")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	//defaultStyle := tcell.StyleDefault.Foreground(tcell.Color104).Background(tcell.ColorBlack).Attributes(tcell.AttrBlink)
	styleOne := tcell.StyleDefault.Foreground(tcell.Color100).Background(tcell.ColorBlack)
	styleTwo := tcell.StyleDefault.Foreground(tcell.Color101).Background(tcell.ColorBlack)
	scene := engine.NewScene("scene", camera)
	textOne := engine.NewEntity("text-one", api.NewPoint(0, 0), api.NewSize(1, 1), &styleOne)
	textOneCanvas := engine.NewCanvasFromString("Hello World!!!", &styleOne)
	textOne.SetCanvas(textOneCanvas)
	scene.AddEntity(textOne)
	textTwo := engine.NewEntity("text-two", api.NewPoint(0, 1), api.NewSize(1, 1), &styleTwo)
	textTwoCanvas := engine.NewCanvasFromString("Hello World******", &styleTwo)
	textTwo.SetCanvas(textTwoCanvas)
	scene.AddEntity(textTwo)
	appEngine := engine.GetEngine()
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	if !appEngine.GetSceneManager().SetSceneAsActive(scene) {
		panic(fmt.Sprintf("can not set scene %s as active", scene.GetName()))
	}
	if !appEngine.GetSceneManager().SetSceneAsVisible(scene) {
		panic(fmt.Sprintf("can not set scene %s as visible", scene.GetName()))
	}
	appEngine.Init()
	appEngine.Run(60.0)
}

func demoThree() {
	tools.Logger.WithField("module", "main").Infof("ThEngine demo-three")
	fmt.Println("ThEngine demo-three")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	scene := engine.NewScene("scene", camera)
	textOne := engine.NewEntity("text-one", api.NewPoint(0, 0), api.NewSize(3, 3), &styleOne)
	textOneCanvas := engine.NewCanvasFromString("\\ /\n O \n/ \\", &styleOne)
	textOne.SetCanvas(textOneCanvas)
	scene.AddEntity(textOne)
	appEngine := engine.GetEngine()
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	appEngine.Init()
	appEngine.Run(60.0)
}

func demoFour(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-four")
	fmt.Println("ThEngine demo-four")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	scene := engine.NewScene("scene", camera)
	textOne := widgets.NewText("text-one", api.NewPoint(0, 0), api.NewSize(3, 3), &styleOne, "Hello\nWorld!")
	scene.AddEntity(textOne)
	appEngine := engine.GetEngine()
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	// appEngine.SetDryRun(dryRun)
	appEngine.Init()
	appEngine.Run(60.0)
}

func demoFive(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-five")
	fmt.Println("ThEngine demo-five")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	styleTwo := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	styleThree := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorRed)
	scene := engine.NewScene("scene", camera)
	textOne := widgets.NewText("text-one", api.NewPoint(0, 0), api.NewSize(1, 1), &styleOne, "Name:")
	scene.AddEntity(textOne)
	inputOne := widgets.NewTextInput("input-one", api.NewPoint(6, 0), api.NewSize(30, 1), &styleTwo, "Jose Carlos Recuero")
	scene.AddEntity(inputOne)
	appEngine := engine.GetEngine()
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	// appEngine.SetDryRun(dryRun)
	appEngine.Init()
	screen := appEngine.GetScreen()
	//style := tcell.StyleDefault
	screen.SetStyle(styleThree)
	screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	//screen.SetStyle(style)
	screen.ShowCursor(0, 0)
	appEngine.Run(60.0)
}

func demoSix(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-six")
	fmt.Println("ThEngine demo-six")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	styleTwo := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorLightBlue)
	scene := engine.NewScene("scene", camera)
	textFirstName := widgets.NewText("text-first-name", api.NewPoint(0, 0), api.NewSize(1, 1), &styleOne, "First Name:")
	scene.AddEntity(textFirstName)
	textLastName := widgets.NewText("text-last-name", api.NewPoint(0, 1), api.NewSize(1, 1), &styleOne, "Last Name:")
	scene.AddEntity(textLastName)
	textAge := widgets.NewText("text-age", api.NewPoint(0, 2), api.NewSize(1, 1), &styleOne, "Age:")
	scene.AddEntity(textAge)
	inputFirstName := widgets.NewTextInput("input-first-name", api.NewPoint(12, 0), api.NewSize(30, 1), &styleTwo, "Jose Carlos")
	scene.AddEntity(inputFirstName)
	inputLastName := widgets.NewTextInput("input-last-name", api.NewPoint(12, 1), api.NewSize(30, 1), &styleTwo, "Recuero Arias")
	scene.AddEntity(inputLastName)
	inputAge := widgets.NewTextInput("input-age", api.NewPoint(12, 2), api.NewSize(30, 1), &styleTwo, "57")
	scene.AddEntity(inputAge)
	appEngine := engine.GetEngine()
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	appEngine.Init()
	screen := appEngine.GetScreen()
	//screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	screen.SetCursorStyle(tcell.CursorStyleSteadyBlock)
	appEngine.Run(60.0)
}

func demoSeven(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-seven")
	fmt.Println("ThEngine demo-seven")
	camera := engine.NewCamera(api.NewPoint(0, 0), api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	scene := engine.NewScene("scene", camera)
	player := NewPlayer("player", api.NewPoint(0, 0), api.NewSize(3, 3), &styleOne)
	playerCanvas := engine.NewCanvasFromString("\\ /\n O \n/ \\", &styleOne)
	player.SetCanvas(playerCanvas)
	scene.AddEntity(player)
	appEngine := engine.GetEngine()
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	appEngine.Init()
	appEngine.Run(60.0)
}

func demoEight(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-eight")
	fmt.Println("ThEngine demo-eight")
	camera := engine.NewCamera(api.NewPoint(0, 0), api.NewSize(10, 5))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	scene := engine.NewScene("scene", camera)
	tileMap := widgets.NewTileMap("tile-map", api.NewPoint(5, 5), api.NewSize(20, 5), &styleOne, api.NewPoint(0, 0), api.NewSize(10, 5))
	cell := engine.NewCell(&styleOne, '|')
	tileMap.GetCanvas().WriteStringInCanvasAt("01234567890123456789", &styleOne, api.NewPoint(0, 0))
	tileMap.GetCanvas().SetCellAt(api.NewPoint(0, 1), cell)
	tileMap.GetCanvas().SetCellAt(api.NewPoint(0, 2), cell)
	tileMap.GetCanvas().SetCellAt(api.NewPoint(0, 3), cell)
	tileMap.GetCanvas().SetCellAt(api.NewPoint(19, 1), cell)
	tileMap.GetCanvas().SetCellAt(api.NewPoint(19, 2), cell)
	tileMap.GetCanvas().SetCellAt(api.NewPoint(19, 3), cell)
	tileMap.GetCanvas().WriteStringInCanvasAt("01234567890123456789", &styleOne, api.NewPoint(0, 4))
	scene.AddEntity(tileMap)
	player := NewPlayer("player", api.NewPoint(7, 7), api.NewSize(1, 1), &styleOne)
	playerCanvas := engine.NewCanvasFromString("x", &styleOne)
	player.SetCanvas(playerCanvas)
	player.TileMap = tileMap
	scene.AddEntity(player)
	appEngine := engine.GetEngine()
	appEngine.GetSceneManager().AddScene(scene)
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	appEngine.Init()
	appEngine.Run(60.0)
}

func demoNine(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-nine")
	fmt.Println("ThEngine demo-nine")
	camera := engine.NewCamera(api.NewPoint(0, 0), api.NewSize(20, 10))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	scene := engine.NewScene("scene", camera)
	handler := newNineHandler()
	scene.AddEntity(handler)
	player := newNinePlayer("player", api.NewPoint(7, 7), api.NewSize(1, 1), &styleOne)
	playerCanvas := engine.NewCanvasFromString("x", &styleOne)
	player.SetCanvas(playerCanvas)
	scene.AddEntity(player)
	appEngine := engine.GetEngine()
	appEngine.GetSceneManager().AddScene(scene)
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	appEngine.Init()
	appEngine.Run(60.0)
}

func demoTen(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-ten")
	fmt.Println("ThEngine demo-ten")
	camera := engine.NewCamera(api.NewPoint(0, 0), api.NewSize(20, 10))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	styleTwo := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorWhite)
	scene := engine.NewScene("scene", camera)
	buttonOne := widgets.NewButton("button/1", api.NewPoint(1, 1), api.NewSize(20, 1), &styleOne, "button-one")
	buttonOne.SetWidgetCallbackArgs("pushed", "one")
	scene.AddEntity(buttonOne)
	buttonTwo := widgets.NewButton("button/2", api.NewPoint(1, 2), api.NewSize(20, 1), &styleTwo, "button-two")
	buttonTwo.SetWidgetCallbackArgs("PUSHED", 2, "active")
	scene.AddEntity(buttonTwo)
	appEngine := engine.GetEngine()
	appEngine.GetSceneManager().AddScene(scene)
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	appEngine.Init()
	appEngine.Run(60.0)
}

func main() {
	demoTen(false)
}
