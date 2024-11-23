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
	GetNode() string
	SetCondition(ICondition)
	SetNode(string)
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

	//  node is the
	node string
}

// -----------------------------------------------------------------------------
// New ConditionalNext functions
// -----------------------------------------------------------------------------

func NewConditionalNext(node string, condition ICondition) *ConditionalNext {
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

func (c *ConditionalNext) GetNode() string {
	return c.node
}

func (c *ConditionalNext) SetCondition(condition ICondition) {
	c.condition = condition
}

func (c *ConditionalNext) SetNode(node string) {
	c.node = node
}

var _ IConditionalNext = (*ConditionalNext)(nil)
