package api

type IKnowledge interface {
	Dec(int)
	Get() int
	Inc(int)
	Set(int)
}

type Knowledge struct {
	value int
}

func NewKnowledge(value int) *Knowledge {
	return &Knowledge{
		value: value,
	}
}

func (k *Knowledge) Dec(value int) {
	k.value -= value
	if k.value < 0 {
		k.value = 0
	}
}

func (k *Knowledge) Get() int {
	return k.value
}

func (k *Knowledge) Inc(value int) {
	k.value += value
}

func (k *Knowledge) Set(value int) {
	k.value = value
}

var _ IKnowledge = (*Knowledge)(nil)
