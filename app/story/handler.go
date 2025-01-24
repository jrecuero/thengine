package main

import (
	"github.com/jrecuero/thengine/app/story/book"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/storyboard"
)

// -----------------------------------------------------------------------------
//
// StoryHandler
//
// -----------------------------------------------------------------------------

type StoryHandler struct {
	*engine.Entity
	book *storyboard.StoryBoard
}

// -----------------------------------------------------------------------------
// New StoryHandler functions
// -----------------------------------------------------------------------------

func NewStoryHandler() *StoryHandler {
	name := "handler/story/1"
	return &StoryHandler{
		Entity: engine.NewHandler(name),
		book:   book.TheBook,
	}
}

// -----------------------------------------------------------------------------
// StoryHandler public methods
// -----------------------------------------------------------------------------
