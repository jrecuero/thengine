package main

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/builder"
	"github.com/jrecuero/thengine/pkg/constants"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

// -----------------------------------------------------------------------------
// Module private methods
// -----------------------------------------------------------------------------

func buildBoxes(scene engine.IScene) {
	headerTextOffset := api.NewPoint(5, 0)
	boxStyle := &constants.WhiteOverBlack

	gameBox := engine.NewEntity(TheGameBoxName, TheGameBoxOrigin,
		TheGameBoxSize, boxStyle)
	gameBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, gameBox.GetStyle(),
		engine.CanvasRectSingleLine)
	scene.AddEntity(gameBox)

	stageBox := engine.NewEntity(TheStageBoxName, TheStageBoxOrigin,
		TheStageBoxSize, boxStyle)
	stageBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, stageBox.GetStyle(),
		engine.CanvasRectSingleLine)
	scene.AddEntity(stageBox)

	stageBoxHeaderTextOrigin := api.ClonePoint(TheStageBoxOrigin)
	stageBoxHeaderTextOrigin.Add(headerTextOffset)
	stageBoxHeaderText := widgets.NewText(TheStageBoxHeaderName,
		stageBoxHeaderTextOrigin, api.NewSize(10, 1), boxStyle, "[ STAGE ]")
	scene.AddEntity(stageBoxHeaderText)

	diceBox := engine.NewEntity(TheDiceBoxName, TheDiceBoxOrigin,
		TheDiceBoxSize, boxStyle)
	diceBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, diceBox.GetStyle(),
		engine.CanvasRectSingleLine)
	scene.AddEntity(diceBox)

	diceBoxHeaderTextOrigin := api.ClonePoint(TheDiceBoxOrigin)
	diceBoxHeaderTextOrigin.Add(headerTextOffset)
	diceBoxHeaderText := widgets.NewText(TheDiceBoxHeaderName,
		diceBoxHeaderTextOrigin, api.NewSize(8, 1), boxStyle, "[ DICE ]")
	scene.AddEntity(diceBoxHeaderText)

	playerBox := engine.NewEntity(ThePlayerBoxName, ThePlayerBoxOrigin, ThePlayerBoxSize, boxStyle)
	playerBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, playerBox.GetStyle(),
		engine.CanvasRectSingleLine)
	scene.AddEntity(playerBox)

	playerBoxHeaderTextOrigin := api.ClonePoint(ThePlayerBoxOrigin)
	playerBoxHeaderTextOrigin.Add(headerTextOffset)
	playerBoxHeaderText := widgets.NewText(ThePlayerBoxHeaderName, playerBoxHeaderTextOrigin,
		api.NewSize(10, 1), boxStyle, "[ PLAYER ]")
	scene.AddEntity(playerBoxHeaderText)

	enemyBox := engine.NewEntity(TheEnemyBoxName, TheEnemyBoxOrigin, TheEnemyBoxSize, boxStyle)
	enemyBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, enemyBox.GetStyle(),
		engine.CanvasRectSingleLine)
	scene.AddEntity(enemyBox)

	enemyBoxHeaderTextOrigin := api.ClonePoint(TheEnemyBoxOrigin)
	enemyBoxHeaderTextOrigin.Add(headerTextOffset)
	enemyBoxHeaderText := widgets.NewText(TheEnemyBoxHeaderName, enemyBoxHeaderTextOrigin,
		api.NewSize(10, 1), boxStyle, "[ ENEMY ]")
	scene.AddEntity(enemyBoxHeaderText)

	keysBox := engine.NewEntity(TheKeysBoxName, TheKeysBoxOrigin, TheKeysBoxSize, boxStyle)
	keysBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, keysBox.GetStyle(),
		engine.CanvasRectSingleLine)
	scene.AddEntity(keysBox)

	keysBoxHeaderTextOrigin := api.ClonePoint(TheKeysBoxOrigin)
	keysBoxHeaderTextOrigin.Add(headerTextOffset)
	keysBoxHeaderText := widgets.NewText(TheKeysBoxHeaderName, keysBoxHeaderTextOrigin,
		api.NewSize(8, 1), boxStyle, "[ KEYS ]")
	scene.AddEntity(keysBoxHeaderText)

	commandLineBox := engine.NewEntity(TheCommandLineBoxName, TheCommandLineBoxOrigin,
		TheCommandLineBoxSize, boxStyle)
	commandLineBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, commandLineBox.GetStyle(),
		engine.CanvasRectSingleLine)
	scene.AddEntity(commandLineBox)

	commandLineBoxHeaderTextOrigin := api.ClonePoint(TheCommandLineBoxOrigin)
	commandLineBoxHeaderTextOrigin.Add(headerTextOffset)
	commandLineBoxHeader := widgets.NewText(TheCommandLineBoxHeaderName,
		commandLineBoxHeaderTextOrigin, api.NewSize(16, 1), boxStyle, "[ COMMAND LINE ]")
	scene.AddEntity(commandLineBoxHeader)
}

func buildDungeon(scene engine.IScene) {

	cell := engine.NewCell(&constants.YellowOverBlack, '#')

	room1 := builder.BuildRoomWithDoors("room/1", api.NewPoint(3, 3),
		api.NewSize(10, 7), cell, []bool{false, false, false, true}, []int{0, 0, 0, 1})
	scene.AddEntity(room1.Sprite)

	room2 := builder.BuildRoomWithDoors("room/2", api.NewPoint(20, 3),
		api.NewSize(12, 7), cell, []bool{false, false, true, false}, []int{0, 0, 1, 0})
	scene.AddEntity(room2.Sprite)

	corridor1 := builder.ConnectRooms("corridor/1", room1.GetDoorAt(builder.RightDoor),
		room2.GetDoorAt(builder.LeftDoor), cell)
	scene.AddEntity(corridor1)

	//cell := engine.NewCell(&constants.YellowOverBlack, '#')
	//room1 := builder.BuildRoom("room/1", api.NewPoint(2, 2), api.NewSize(10, 7), cell,
	//    []bool{false, false, false, true})
	//scene.AddEntity(room1)

	//corridor1 := builder.BuildCorridor("corridor/1", api.NewPoint(12, 5), api.NewPoint(16, 5), cell)
	//scene.AddEntity(corridor1)

	//room2 := builder.BuildRoom("room/2", api.NewPoint(17, 3), api.NewSize(5, 5), cell,
	//    []bool{false, true, true, false})
	//scene.AddEntity(room2)

	//corridor2 := builder.BuildCorridor("corridor/2", api.NewPoint(19, 8), api.NewPoint(19, 11), cell)
	//scene.AddEntity(corridor2)

	//room3 := builder.BuildRoom("room/3", api.NewPoint(13, 12), api.NewSize(13, 7), cell,
	//    []bool{true, false, false, true})
	//scene.AddEntity(room3)

	//corridor3 := builder.BuildCorridor("corridor/3", api.NewPoint(26, 15), api.NewPoint(31, 15), cell)
	//scene.AddEntity(corridor3)

	//room4 := builder.BuildRoom("room/4", api.NewPoint(32, 14), api.NewSize(3, 3), cell,
	//    []bool{true, false, true, false})
	//scene.AddEntity(room4)

	//corridor4 := builder.BuildCorridor("corridor/4", api.NewPoint(33, 11), api.NewPoint(33, 13), cell)
	//scene.AddEntity(corridor4)

	//room5 := builder.BuildRoom("room/5", api.NewPoint(29, 5), api.NewSize(9, 7), cell,
	//    []bool{false, true, false, true})
	//scene.AddEntity(room5)

	//corridor5 := builder.BuildCorridor("corridor/5", api.NewPoint(38, 8), api.NewPoint(51, 8), cell)
	//scene.AddEntity(corridor5)

	//room6 := builder.BuildRoom("room/6", api.NewPoint(52, 3), api.NewSize(20, 11), cell,
	//    []bool{false, false, true, false})
	//scene.AddEntity(room6)
}
