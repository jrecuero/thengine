// battlegear.go contains all common atributes and behavior related to any
// battle gear like weapons, shield or any piece of armor.
package rules

// -----------------------------------------------------------------------------
//
// IBattleGear
//
// -----------------------------------------------------------------------------

// IBattleGear interface defines all methods any piece of battle gear should be
// implementing
type IBattleGear interface {
	IDamage
	GetAC() int
	GetCost() int
	GetDescription() string
	GetMaterial() string
	GetModifiers() []any
	GetName() string
	GetProps() []any
	GetQuality() string
	GetUName() string
	GetWeight() int
	RollDamage() int
	SetAC(int)
	SetCost(int)
	SetDescription(string)
	SetMaterial(string)
	SetModifiers([]any)
	SetName(string)
	SetProps([]any)
	SetQuality(string)
	SetUName(string)
	SetWeight(int)
}

// -----------------------------------------------------------------------------
//
// BattleGear
//
// -----------------------------------------------------------------------------

// BattleGear structure represents all data and behavior for any piece of
// battle gear.
type BattleGear struct {
	*Damage
	ac          int
	cost        int
	description string
	material    string
	modifiers   []any
	name        string
	props       []any
	quality     string
	uname       string
	weight      int
}

func NewBattleGear(name string, uname string, cost int, weight int) *BattleGear {
	shield := &BattleGear{
		ac:          0,
		cost:        cost,
		description: name,
		material:    "",
		modifiers:   nil,
		name:        name,
		props:       nil,
		quality:     "",
		uname:       uname,
		weight:      weight,
	}
	shield.Damage = NewNoDamage()
	return shield
}

// -----------------------------------------------------------------------------
// BattleGear public methods
// -----------------------------------------------------------------------------

func (h *BattleGear) GetAC() int {
	return h.ac
}

func (h *BattleGear) GetCost() int {
	return h.cost
}

func (h *BattleGear) GetDescription() string {
	return h.description
}

func (h *BattleGear) GetMaterial() string {
	return h.material
}

func (h *BattleGear) GetModifiers() []any {
	return h.modifiers
}

func (h *BattleGear) GetName() string {
	return h.name
}

func (h *BattleGear) GetProps() []any {
	return h.props
}

func (h *BattleGear) GetQuality() string {
	return h.quality
}

func (h *BattleGear) GetUName() string {
	return h.uname
}

func (h *BattleGear) GetWeight() int {
	return h.weight
}

func (h *BattleGear) RollDamage() int {
	return h.RollDamageValue()
}

func (h *BattleGear) SetAC(ac int) {
	h.ac = ac
}

func (h *BattleGear) SetCost(cost int) {
	h.cost = cost
}

func (h *BattleGear) SetDescription(description string) {
	h.description = description
}

func (h *BattleGear) SetMaterial(material string) {
	h.material = material
}

func (h *BattleGear) SetModifiers(modifiers []any) {
	h.modifiers = modifiers
}

func (h *BattleGear) SetName(name string) {
	h.name = name
}

func (h *BattleGear) SetProps(props []any) {
	h.props = props
}

func (h *BattleGear) SetQuality(quality string) {
	h.quality = quality
}

func (h *BattleGear) SetUName(uname string) {
	h.uname = uname
}

func (h *BattleGear) SetWeight(weight int) {
	h.weight = weight
}

var _ IDamage = (*BattleGear)(nil)
var _ IBattleGear = (*BattleGear)(nil)
