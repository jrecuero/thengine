// node.go contains attributes and logic related with any storyboard node
// with narrator information for the storyboard.
package storyboard

// -----------------------------------------------------------------------------
//
// INode
//
// -----------------------------------------------------------------------------

type INode interface {
	AddNext(...IConditionalNext)
	AddQuestion(...IQuestion)
	AddText(...IConditionalText)
	GetID() string
	GetNext() []IConditionalNext
	GetQuestions() []IQuestion
	GetSpeaker() string
	GetText() []IConditionalText
	SetSpeaker(string)
}

// -----------------------------------------------------------------------------
//
// Node
//
// -----------------------------------------------------------------------------

type Node struct {
	// id is a unique identification for the node
	id string

	// next is a slice with possible next node where the storyboard will
	// continue. Next Node could be a fixed one or there could be some
	// conditions about what could be the next Node.
	next []IConditionalNext

	// question is a list of questions that are used to decide what is the
	// next Node.
	questions []IQuestion

	// speaker is the storyboard character for this dialog.
	speaker string

	// text are a slice of strings to be displayed for the Node.
	text []IConditionalText
}

// -----------------------------------------------------------------------------
// New Node functions
// -----------------------------------------------------------------------------

func NewNode(id string) *Node {
	return &Node{
		id:        id,
		next:      nil,
		questions: nil,
		speaker:   "unknown",
		text:      make([]IConditionalText, 0),
	}
}

// -----------------------------------------------------------------------------
// Node public methods
// -----------------------------------------------------------------------------

func (n *Node) AddNext(next ...IConditionalNext) {
	n.next = append(n.next, next...)
}

func (n *Node) AddQuestion(question ...IQuestion) {
	n.questions = append(n.questions, question...)
}

func (n *Node) AddText(text ...IConditionalText) {
	n.text = append(n.text, text...)
}

func (n *Node) GetID() string {
	return n.id
}

func (n *Node) GetNext() []IConditionalNext {
	return n.next
}

func (n *Node) GetQuestions() []IQuestion {
	return n.questions
}

func (n *Node) GetSpeaker() string {
	return n.speaker
}

func (n *Node) GetText() []IConditionalText {
	return n.text
}

func (n *Node) SetSpeaker(speaker string) {
	n.speaker = speaker
}

var _ INode = (*Node)(nil)
