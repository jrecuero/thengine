// handheld.go contains all data and methods related with weapons or shields
// that can be used by any hand.
package rules

// -----------------------------------------------------------------------------
//
// HandheldType
//
// -----------------------------------------------------------------------------

// HandheldType structure defines the type of handheld, like if it can be used
// by one or two hands, used in the offhand, dual hands, dual grip,...
type HandheldType struct {
	hands     int
	offhand   bool
	dualHands bool
	dualGrip  bool
	finesse   bool
	martial   bool
}

func NewHandheldType(hands int) *HandheldType {
	return &HandheldType{
		hands:     hands,
		offhand:   false,
		dualHands: false,
		dualGrip:  false,
		finesse:   false,
		martial:   false,
	}
}

func NewShieldHandheldType() *HandheldType {
	return &HandheldType{
		hands:     1,
		offhand:   true,
		dualHands: false,
		dualGrip:  false,
		finesse:   false,
		martial:   false,
	}
}

// -----------------------------------------------------------------------------
// HandheldType public methods
// -----------------------------------------------------------------------------

func (h *HandheldType) Hands() int {
	return h.hands
}

func (h *HandheldType) IsDualHands() bool {
	return h.dualHands
}

func (h *HandheldType) IsDualGrip() bool {
	return h.dualGrip
}

func (h *HandheldType) IsFinesse() bool {
	return h.finesse
}

func (h *HandheldType) IsOffhand() bool {
	return h.offhand
}

func (h *HandheldType) IsMartial() bool {
	return h.martial
}

func (h *HandheldType) SetDualHands(flag bool) *HandheldType {
	h.dualHands = flag
	return h
}

func (h *HandheldType) SetDualGrip(flag bool) *HandheldType {
	h.dualGrip = flag
	return h
}

func (h *HandheldType) SetFinesse(flag bool) *HandheldType {
	h.finesse = flag
	return h
}

func (h *HandheldType) SetHands(hands int) *HandheldType {
	h.hands = hands
	return h
}

func (h *HandheldType) SetOffhand(flag bool) *HandheldType {
	h.offhand = flag
	return h
}

func (h *HandheldType) SetMartial(flag bool) *HandheldType {
	h.martial = flag
	return h
}

// -----------------------------------------------------------------------------
//
// IHandheld
//
// -----------------------------------------------------------------------------

// IHandheld interface defines all methods any Weapon or Shield should implement
type IHandheld interface {
	IBattleGear
	GetHType() *HandheldType
	SetHType(*HandheldType)
}

// -----------------------------------------------------------------------------
//
// Handheld
//
// -----------------------------------------------------------------------------

// Handheld structure represents any armament that can be used by a hand, like
// a weapon or a shield.
type Handheld struct {
	*BattleGear
	htype *HandheldType
}

func NewHandheld(name string, uname string, cost int, weight int, htype *HandheldType) *Handheld {
	return &Handheld{
		BattleGear: NewBattleGear(name, uname, cost, weight),
		htype:      htype,
	}
}

// -----------------------------------------------------------------------------
// Handheld public methods
// -----------------------------------------------------------------------------

func (h *Handheld) GetHType() *HandheldType {
	return h.htype
}

func (h *Handheld) SetHType(htype *HandheldType) {
	h.htype = htype
}

var _ IDamage = (*Handheld)(nil)
var _ IBattleGear = (*Handheld)(nil)
var _ IHandheld = (*Handheld)(nil)
