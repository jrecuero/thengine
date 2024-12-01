// question.go contains attributes and logic related with any storyboard
// question.
package storyboard

// -----------------------------------------------------------------------------
//
// IQuestion
//
// -----------------------------------------------------------------------------

type IQuestion interface {
	IBaseNode
	GetCondition() ICondition
	SetCondition(ICondition)
}

// -----------------------------------------------------------------------------
//
// Question
//
// -----------------------------------------------------------------------------

type Question struct {
	// BaseNode contains base and common Question information with Node.
	*BaseNode

	// condition is a string or a function to be used to decide is the
	// question has to be displayed or not.
	condition ICondition
}

// -----------------------------------------------------------------------------
// New Question functions
// -----------------------------------------------------------------------------

func NewQuestion() *Question {
	return &Question{
		BaseNode:  NewBaseNode(""),
		condition: nil,
	}
}

// -----------------------------------------------------------------------------
// Question public methods
// -----------------------------------------------------------------------------

func (q *Question) GetCondition() ICondition {
	return q.condition
}

func (q *Question) SetCondition(condition ICondition) {
	q.condition = condition
}

var _ IBaseNode = (*Question)(nil)
var _ IQuestion = (*Question)(nil)
