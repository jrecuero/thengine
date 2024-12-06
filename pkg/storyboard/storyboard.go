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
	nodes []IBaseNode

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
// StoryBoard private functions
// -----------------------------------------------------------------------------

// getAllNodes method returns all children nodes for the given node. All
// children include INode and IQuestion instances.
func (s *StoryBoard) getAllNodes(basenode IBaseNode) []IBaseNode {
	nodes := []IBaseNode{basenode}
	for _, n := range basenode.GetNext() {
		// nodes = append(nodes, n.GetNode())
		nodes = append(nodes, s.getAllNodes(n.GetNode())...)
	}
	if node, ok := basenode.(INode); ok {
		for _, question := range node.GetQuestions() {
			// nodes = append(nodes, question)
			nodes = append(nodes, s.getAllNodes(question)...)
		}
	}
	return nodes
}

func (s *StoryBoard) populateNodes(node INode) {
	s.start = node
	s.nodes = s.getAllNodes(node)
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

func (s *StoryBoard) GetNodes() []IBaseNode {
	return s.nodes
}

func (s StoryBoard) GetStart() INode {
	return s.start
}

func (s *StoryBoard) GetNodeWithID(id string) INode {
	if result, found := tools.RetrieveGenericAny(s.nodes, id, func(a IBaseNode, b string) bool {
		return a.GetID() == b
	}); found {
		return result.(INode)
	}
	return nil
}

func (s *StoryBoard) SetCurrent(node INode) {
	if s.current == nil {
		s.populateNodes(node)
	}
	s.current = node
}

func (s *StoryBoard) SetName(name string) {
	s.name = name
}

func (s *StoryBoard) SetStart(node INode) {
	s.current = nil
	s.SetCurrent(node)
}
