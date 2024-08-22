package api

// -----------------------------------------------------------------------------
//
// ICatable
//
// -----------------------------------------------------------------------------

type ICatable interface {
	GetCat() IComparable
}

func FindByCat[T ICatable](items []T, cat IComparable) []T {
	result := []T{}
	for _, item := range items {
		if item.GetCat().Equal(cat) {
			result = append(result, item)
		}
	}
	return result
}

// -----------------------------------------------------------------------------
//
// IComparable
//
// -----------------------------------------------------------------------------

type IComparable interface {
	Equal(IComparable) bool
}

// -----------------------------------------------------------------------------
//
// INameable
//
// -----------------------------------------------------------------------------

type INameable interface {
	GetName() string
}

func FindByName[T INameable](items []T, name string) (T, bool) {
	item, _, found := FindByNameWithIndex(items, name)
	return item, found
}

func FindByNameWithIndex[T INameable](items []T, name string) (T, int, bool) {
	for index, item := range items {
		if item.GetName() == name {
			return item, index, true
		}
	}
	var zeroValue T
	return zeroValue, -1, false
}
