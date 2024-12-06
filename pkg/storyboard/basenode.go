// basenode.go contains a base node with all common and share functionality by
// Node and Question instances.
package storyboard

// -----------------------------------------------------------------------------
//
// IBaseNode
//
// -----------------------------------------------------------------------------

type IBaseNode interface {
	AddNext(...any)
	AddText(...any)
	GetID() string
	GetNext() []IConditionalNext
	GetNextNodes() []IBaseNode
	GetSpeaker() string
	GetText() []IConditionalText
	SetSpeaker(string)
}

// -----------------------------------------------------------------------------
//
// BaseNode
//
// -----------------------------------------------------------------------------

type BaseNode struct {
	// id is a unique identification for the node
	id string

	// next is a slice with possible next node where the storyboard will
	// continue. Next Node could be a fixed one or there could be some
	// conditions about what could be the next Node.
	next []IConditionalNext

	// speaker is the storyboard character for this dialog.
	speaker string

	// text are a slice of strings to be displayed for the Node.
	text []IConditionalText
}

// -----------------------------------------------------------------------------
// New Node functions
// -----------------------------------------------------------------------------

func NewBaseNode(id string) *BaseNode {
	return &BaseNode{
		id:      id,
		next:    nil,
		speaker: "unknown",
		text:    make([]IConditionalText, 0),
	}
}

// -----------------------------------------------------------------------------
// BaseNode public methods
// -----------------------------------------------------------------------------

func (n *BaseNode) AddNext(nexts ...any) {
	for _, entry := range nexts {
		var input IConditionalNext
		switch next := entry.(type) {
		case INode:
			input = NewConditionalNext(next, nil)
		case IConditionalNext:
			input = next
		default:
			input = nil
		}
		if input != nil {
			n.next = append(n.next, input)
		}
	}
}

func (n *BaseNode) AddText(texts ...any) {
	for _, entry := range texts {
		var input IConditionalText
		switch t := entry.(type) {
		case string:
			input = NewConditionalText(t, nil)
		case IConditionalText:
			input = t
		default:
			input = nil
		}
		if input != nil {
			n.text = append(n.text, input)
		}
	}
}

func (n *BaseNode) GetID() string {
	return n.id
}

func (n *BaseNode) GetNext() []IConditionalNext {
	return n.next
}

func (n *BaseNode) GetNextNodes() []IBaseNode {
	var nodes []IBaseNode
	for _, conditionalNext := range n.GetNext() {
		nodes = append(nodes, conditionalNext.GetNode())
	}
	return nodes
}

func (n *BaseNode) GetSpeaker() string {
	return n.speaker
}

func (n *BaseNode) GetText() []IConditionalText {
	return n.text
}

func (n *BaseNode) SetSpeaker(speaker string) {
	n.speaker = speaker
}

var _ IBaseNode = (*BaseNode)(nil)
