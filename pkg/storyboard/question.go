// question.go contains attributes and logic related with any storyboard
// question.
package storyboard

// -----------------------------------------------------------------------------
//
// IQuestion
//
// -----------------------------------------------------------------------------

type IQuestion interface {
	AddNext(...IConditionalNext)
	AddText(...IConditionalText)
	GetCondition() ICondition
	GetNext() []IConditionalNext
	GetText() []IConditionalText
	SetCondition(ICondition)
}

// -----------------------------------------------------------------------------
//
// Question
//
// -----------------------------------------------------------------------------

type Question struct {
	// condition is a string or a function to be used to decide is the
	// question has to be displayed or not.
	condition ICondition

	// next is a slice with possible next node where the storyboard will
	// continue. Next Node could be a fixed one or there could be some
	// conditions about what could be the next Node.
	next []IConditionalNext

	// texts are a slice of strings to be displyes for the Question.
	text []IConditionalText
}

// -----------------------------------------------------------------------------
// New Question functions
// -----------------------------------------------------------------------------

func NewQuestion() *Question {
	return &Question{
		condition: nil,
		next:      nil,
		text:      make([]IConditionalText, 0),
	}
}

// -----------------------------------------------------------------------------
// Question public methods
// -----------------------------------------------------------------------------

func (q *Question) AddNext(next ...IConditionalNext) {
	q.next = append(q.next, next...)
}

func (q *Question) AddText(text ...IConditionalText) {
	q.text = append(q.text, text...)
}

func (q *Question) GetCondition() ICondition {
	return q.condition
}

func (q *Question) GetNext() []IConditionalNext {
	return q.next
}

func (q *Question) GetText() []IConditionalText {
	return q.text
}

func (q *Question) SetCondition(condition ICondition) {
	q.condition = condition
}

var _ IQuestion = (*Question)(nil)
