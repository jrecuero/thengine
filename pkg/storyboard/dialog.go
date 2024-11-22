// dialog.go package contains all data and logic related with any dialog to be
// included in any event in the storyboard.
package storyboard

// -----------------------------------------------------------------------------
//
// IDialog
//
// -----------------------------------------------------------------------------

type IDialog interface {
	GetSpeaker() ICharacter
	GetText() string
	SetSpeaker(ICharacter)
	SetText(string)
}

// -----------------------------------------------------------------------------
//
// Dialog
//
// -----------------------------------------------------------------------------

type Dialog struct {
	speaker ICharacter
	text    string
}

// -----------------------------------------------------------------------------
// Dialog public methods
// -----------------------------------------------------------------------------

func (d *Dialog) GetSpeaker() ICharacter {
	return d.speaker
}

func (d *Dialog) GetText() string {
	return d.text
}

func (d *Dialog) SetSpeaker(speaker ICharacter) {
	d.speaker = speaker
}

func (d *Dialog) SetText(text string) {
	d.text = text
}

var _ IDialog = (*Dialog)(nil)
