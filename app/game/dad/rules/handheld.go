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

// IHandheld interface defines al methods any Weapon or Shield should implement
type IHandheld interface {
	GetAC() int
	GetCost() int
	GetDamage() IDiceThrow
	GetDamageType() DamageType
	GetDescription() string
	GetHType() *HandheldType
	GetName() string
	GetWeight() int
	RollDamage() int
	SetAC(int)
	SetCost(int)
	SetDamage(IDiceThrow)
	SetDamageType(DamageType)
	SetDescription(string)
	SetHType(*HandheldType)
	SetName(string)
	SetWeight(int)
}

// -----------------------------------------------------------------------------
//
// Handheld
//
// -----------------------------------------------------------------------------

// Handheld structure represents any armament that can be used by a hand, like
// a weapon or a shield.
type Handheld struct {
	name        string
	description string
	cost        int
	weight      int
	htype       *HandheldType
	damage      IDiceThrow
	damageType  DamageType
	ac          int
}

func NewHandheld(name string, cost int, weight int, htype *HandheldType) *Handheld {
	return &Handheld{
		name:        name,
		description: name,
		cost:        cost,
		weight:      weight,
		htype:       htype,
		damage:      nil,
		damageType:  NullDamage,
		ac:          0,
	}
}

// -----------------------------------------------------------------------------
// Handheld public methods
// -----------------------------------------------------------------------------

func (h *Handheld) GetAC() int {
	return h.ac
}

func (h *Handheld) GetCost() int {
	return h.cost
}

func (h *Handheld) GetDamage() IDiceThrow {
	return h.damage
}

func (h *Handheld) GetDamageType() DamageType {
	return h.damageType
}

func (h *Handheld) GetHType() *HandheldType {
	return h.htype
}

func (h *Handheld) GetDescription() string {
	return h.description
}

func (h *Handheld) GetName() string {
	return h.name
}

func (h *Handheld) GetWeight() int {
	return h.weight
}

func (h *Handheld) RollDamage() int {
	if h.damage != nil {
		return h.damage.Roll()
	}
	return 0
}

func (h *Handheld) SetAC(ac int) {
	h.ac = ac
}

func (h *Handheld) SetCost(cost int) {
	h.cost = cost
}

func (h *Handheld) SetDamage(damage IDiceThrow) {
	h.damage = damage
}

func (h *Handheld) SetDamageType(damageType DamageType) {
	h.damageType = damageType
}

func (h *Handheld) SetDescription(description string) {
	h.description = description
}

func (h *Handheld) SetHType(htype *HandheldType) {
	h.htype = htype
}

func (h *Handheld) SetName(name string) {
	h.name = name
}

func (h *Handheld) SetWeight(weight int) {
	h.weight = weight
}

var _ IHandheld = (*Handheld)(nil)
