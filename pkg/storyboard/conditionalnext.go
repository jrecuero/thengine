// conditionalnext.go contains attributes and logic related with the next Node
// to be selected in the storyboard
package storyboard

// -----------------------------------------------------------------------------
//
// IConditionalNext
//
// -----------------------------------------------------------------------------

type IConditionalNext interface {
	GetCondition() ICondition
	GetNode() INode
	SetCondition(ICondition)
	SetNode(INode)
}

// -----------------------------------------------------------------------------
//
// ConditionalNext
//
// -----------------------------------------------------------------------------

type ConditionalNext struct {
	// condition is a string or a function to be used to decide if the this has
	// to be the next Node in the story board
	condition ICondition

	//  node is the INode instance to follow up.
	node INode
}

// -----------------------------------------------------------------------------
// New ConditionalNext functions
// -----------------------------------------------------------------------------

func NewConditionalNext(node INode, condition ICondition) *ConditionalNext {
	return &ConditionalNext{
		condition: condition,
		node:      node,
	}
}

// -----------------------------------------------------------------------------
// ConditionalNext public methods
// -----------------------------------------------------------------------------

func (c *ConditionalNext) GetCondition() ICondition {
	return c.condition
}

func (c *ConditionalNext) GetNode() INode {
	return c.node
}

func (c *ConditionalNext) SetCondition(condition ICondition) {
	c.condition = condition
}

func (c *ConditionalNext) SetNode(node INode) {
	c.node = node
}

var _ IConditionalNext = (*ConditionalNext)(nil)
