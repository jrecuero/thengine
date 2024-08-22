package api

type ERhuneCat int

const (
	BaseRhune ERhuneCat = iota
	ExtraRhune
)

func (r ERhuneCat) Equal(other IComparable) bool {
	if o, ok := other.(ERhuneCat); ok {
		return r == o
	}
	return false
}
