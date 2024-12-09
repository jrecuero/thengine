package main

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

// -----------------------------------------------------------------------------
// Module private methods
// -----------------------------------------------------------------------------

func buildBoxes(scene engine.IScene, handler *StoryHandler) {
	headerTextOffset := api.NewPoint(5, 0)

	storyBox := engine.NewEntity(TheStoryBoxName, TheStoryBoxOrigin,
		TheStoryBoxSize, theBoxStyle)
	storyBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, storyBox.GetStyle(),
		engine.CanvasRectSingleLine)
	scene.AddEntity(storyBox)

	storyNameOrigin := api.ClonePoint(TheStoryBoxOrigin)
	storyNameOrigin.Add(headerTextOffset)
	storyNameText := widgets.NewText(TheStoryTextName, storyNameOrigin,
		api.NewSize(9, 1), theBoxStyle, "[ STORY ]")
	scene.AddEntity(storyNameText)

}
