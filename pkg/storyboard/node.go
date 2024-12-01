// node.go contains attributes and logic related with any storyboard node
// with narrator information for the storyboard.
package storyboard

// -----------------------------------------------------------------------------
//
// INode
//
// -----------------------------------------------------------------------------

type INode interface {
	IBaseNode
	AddQuestion(...IQuestion)
	GetQuestions() []IQuestion
}

// -----------------------------------------------------------------------------
//
// Node
//
// -----------------------------------------------------------------------------

type Node struct {
	// BaseNode contains base and common Node information with Question.
	*BaseNode

	// question is a list of questions that are used to decide what is the
	// next Node.
	questions []IQuestion
}

// -----------------------------------------------------------------------------
// New Node functions
// -----------------------------------------------------------------------------

func NewNode(id string) *Node {
	return &Node{
		BaseNode:  NewBaseNode(id),
		questions: nil,
	}
}

// -----------------------------------------------------------------------------
// Node public methods
// -----------------------------------------------------------------------------

func (n *Node) AddQuestion(question ...IQuestion) {
	n.questions = append(n.questions, question...)
}

func (n *Node) GetID() string {
	return n.id
}

func (n *Node) GetQuestions() []IQuestion {
	return n.questions
}

var _IBaseNode = (*Node)(nil)
var _ INode = (*Node)(nil)
