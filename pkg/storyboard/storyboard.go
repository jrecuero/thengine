// storyboard.go contains all attributes and logic in order to handle a
// storyboard dialog sequence.
package storyboard

import "github.com/jrecuero/thengine/pkg/tools"

// -----------------------------------------------------------------------------
//
// StoryBoard
//
// -----------------------------------------------------------------------------

type StoryBoard struct {
	// current INode instance with the node being processed.
	current INode

	// name string with the name of the story board.
	name string

	// node slice of INode with all nodes in the story board.
	nodes []INode

	// start INode instance to the initial node in the story board.
	start INode
}

// -----------------------------------------------------------------------------
// New StoryBoard functions
// -----------------------------------------------------------------------------

func NewStoryBoard(name string) *StoryBoard {
	return &StoryBoard{
		current: nil,
		name:    name,
		nodes:   nil,
		start:   nil,
	}
}

// -----------------------------------------------------------------------------
// StoryBoard public functions
// -----------------------------------------------------------------------------

func (s *StoryBoard) GetCurrent() INode {
	return s.current
}

func (s *StoryBoard) GetName() string {
	return s.name
}

func (s *StoryBoard) GetNodes() []INode {
	return s.nodes
}

func (s StoryBoard) GetStart() INode {
	return s.start
}

func (s *StoryBoard) GetNodeWithID(id string) INode {
	if result, found := tools.RetrieveGenericAny(s.nodes, id, func(a INode, b string) bool {
		return a.GetID() == b
	}); found {
		return result
	}
	return nil
}

func (s *StoryBoard) SetCurrent(node INode) {
	s.current = node
}

func (s *StoryBoard) SetName(name string) {
	s.name = name
}

func (s *StoryBoard) SetStart(node INode) {
	s.start = node
}
