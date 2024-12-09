package main

import "github.com/jrecuero/thengine/pkg/engine"

// -----------------------------------------------------------------------------
//
// StoryHandler
//
// -----------------------------------------------------------------------------

type StoryHandler struct {
	*engine.Entity
}

// -----------------------------------------------------------------------------
// New StoryHandler functions
// -----------------------------------------------------------------------------

func NewStoryHandler() *StoryHandler {
	name := "handler/story/1"
	return &StoryHandler{
		Entity: engine.NewHandler(name),
	}
}
