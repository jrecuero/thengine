package api

type IRhune interface {
	Execute(IAvatar) (any, error)
	GetCat() IComparable
	GetDescription() string
	GetName() string
	GetShort() string
	SetExecute(func(IAvatar) (any, error))
}

type Rhune struct {
	cat         ERhuneCat
	description string
	execute     func(IAvatar) (any, error)
	name        string
	short       string
}

func NewRhune(name string, short string, description string, cat ERhuneCat,
	execute func(IAvatar) (any, error)) *Rhune {
	return &Rhune{
		cat:         cat,
		description: description,
		execute:     execute,
		name:        name,
		short:       short,
	}
}

func (r *Rhune) Execute(avatar IAvatar) (any, error) {
	if r.execute != nil {
		return r.execute(avatar)
	}
	return nil, nil
}

func (r *Rhune) GetCat() IComparable {
	return r.cat
}

func (r *Rhune) GetDescription() string {
	return r.description
}

func (r *Rhune) GetName() string {
	return r.name
}

func (r *Rhune) GetShort() string {
	return r.short
}

func (r *Rhune) SetExecute(execute func(IAvatar) (any, error)) {
	r.execute = execute
}

var _ IRhune = (*Rhune)(nil)
