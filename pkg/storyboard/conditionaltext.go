// conditionaltext.go contains attributes and logic related with any text and
// related condition when text has to be applied.
package storyboard

// -----------------------------------------------------------------------------
//
// IConditionalText
//
// -----------------------------------------------------------------------------

type IConditionalText interface {
	GetCondition() ICondition
	GetText() string
	SetCondition(ICondition)
	SetText(string)
}

// -----------------------------------------------------------------------------
//
// ConditionalText
//
// -----------------------------------------------------------------------------

type ConditionalText struct {
	// condition is a string or a function to be used to decide if the text
	// has to displayed or not.
	condition ICondition

	// text is the string to be displayed.
	text string
}

// -----------------------------------------------------------------------------
// New ConditionalText function
// -----------------------------------------------------------------------------

func NewConditionalText(text string, condition ICondition) *ConditionalText {
	return &ConditionalText{
		condition: condition,
		text:      text,
	}
}

// -----------------------------------------------------------------------------
// ConditionalText public methods
// -----------------------------------------------------------------------------

func (c *ConditionalText) GetCondition() ICondition {
	return c.condition
}

func (c *ConditionalText) GetText() string {
	return c.text
}

func (c *ConditionalText) SetCondition(condition ICondition) {
	c.condition = condition
}

func (c *ConditionalText) SetText(text string) {
	c.text = text
}

var _ IConditionalText = (*ConditionalText)(nil)
